package db

import (
	"context"
	"reflect"

	"github.com/donbarrigon/utils/herror"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoModel interface {
	GetID() bson.ObjectID
	Coll() string
}

type OdmModel interface {
	CollectionName() string
	GetID() bson.ObjectID
	SetID(id bson.ObjectID)
	BeforeCreate() herror.Error
	BeforeUpdate() herror.Error
	BeforeDelete() herror.Error
	AfterCreate() herror.Error
	AfterUpdate() herror.Error
	AfterDelete() herror.Error

	// funciones nesesarias para el move to trash
	Create() herror.Error
	Delete() herror.Error

	// funciones nesesarias para el history record
	GetOriginal() map[string]any
	GetDirty() map[string]any
	SetOriginal(original map[string]any)
	SetDirty(dirty map[string]any)
}

type Collection []OdmModel

type Odm struct {
	Model    OdmModel       `bson:"-" json:"-"`
	dirty    map[string]any `bson:"-" json:"-"`
	original map[string]any `bson:"-" json:"-"`
}

// fuciones para parchar cosas estas se tienen que sobreescribir
func (o *Odm) CollectionName() string { return "" }
func (o *Odm) GetID() bson.ObjectID   { return bson.NewObjectID() }
func (o *Odm) SetID(id bson.ObjectID) {}

// funciones observers
func (o *Odm) BeforeCreate() herror.Error { return nil }
func (o *Odm) BeforeUpdate() herror.Error { return nil }
func (o *Odm) BeforeDelete() herror.Error { return nil }
func (o *Odm) AfterCreate() herror.Error  { return nil }
func (o *Odm) AfterUpdate() herror.Error  { return nil }
func (o *Odm) AfterDelete() herror.Error  { return nil }

// funciones utiles de para las funciones fill y cositas por ahi
func (o *Odm) GetOriginal() map[string]any         { return o.original }
func (o *Odm) GetDirty() map[string]any            { return o.dirty }
func (o *Odm) SetOriginal(original map[string]any) { o.original = original }
func (o *Odm) SetDirty(dirty map[string]any)       { o.dirty = dirty }

// funciones para azucar sintactico
func (o *Odm) Fill(validator any) herror.Error      { return Fill(o, validator) }
func (o *Odm) FillDirty(validator any) herror.Error { return FillDirty(o, validator) }

// ================================================================
// funciones CRUD
// ================================================================

func (o *Odm) FindByHexID(id string) herror.Error {

	objectId, e := bson.ObjectIDFromHex(id)
	if e != nil {
		return herror.HexID(e)
	}
	filter := bson.D{bson.E{Key: "_id", Value: objectId}}
	if e := Mongo.Collection(o.Model.CollectionName()).FindOne(context.TODO(), filter).Decode(o.Model); e != nil {
		return herror.Mongo(e)
	}
	return nil
}

func (o *Odm) FindByID(id bson.ObjectID) herror.Error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	if e := Mongo.Collection(o.Model.CollectionName()).FindOne(context.TODO(), filter).Decode(o.Model); e != nil {
		return herror.Mongo(e)
	}
	return nil
}

func (o *Odm) First(field string, value any) herror.Error {
	filter := bson.D{bson.E{Key: field, Value: value}}
	if e := Mongo.Collection(o.Model.CollectionName()).FindOne(context.TODO(), filter).Decode(o.Model); e != nil {
		return herror.Mongo(e)
	}
	return nil
}

func (o *Odm) FindOne(filter bson.D, opts ...options.Lister[options.FindOneOptions]) herror.Error {
	if e := Mongo.Collection(o.Model.CollectionName()).FindOne(context.TODO(), filter, opts...).Decode(o.Model); e != nil {
		return herror.Mongo(e)
	}
	return nil
}

func (o *Odm) Find(result any, filter bson.D, opts ...options.Lister[options.FindOptions]) herror.Error {
	ctx := context.TODO()
	cursor, e := Mongo.Collection(o.Model.CollectionName()).Find(ctx, filter, opts...)
	if e != nil {
		return herror.Mongo(e)
	}
	if e = cursor.All(ctx, result); e != nil {
		return herror.Mongo(e)
	}
	return nil
}

// busqueda eq
func (o *Odm) FindByField(result any, field string, value any, opts ...options.Lister[options.FindOptions]) herror.Error {
	filter := bson.D{bson.E{Key: field, Value: value}}
	ctx := context.TODO()
	cursor, e := Mongo.Collection(o.Model.CollectionName()).Find(ctx, filter, opts...)
	if e != nil {
		return herror.Mongo(e)
	}
	if e = cursor.All(ctx, result); e != nil {
		return herror.Mongo(e)
	}
	return nil
}

func (o *Odm) Aggregate(result any, pipeline mongo.Pipeline) herror.Error {
	ctx := context.TODO()
	cursor, e := Mongo.Collection(o.Model.CollectionName()).Aggregate(ctx, pipeline)
	if e != nil {
		return herror.Mongo(e)
	}
	if e = cursor.All(ctx, result); e != nil {
		return herror.Mongo(e)
	}
	return nil
}

func (o *Odm) AggregateOne(pipeline mongo.Pipeline) herror.Error {
	ctx := context.TODO()
	cursor, e := Mongo.Collection(o.Model.CollectionName()).Aggregate(ctx, pipeline)
	if e != nil {
		return herror.Mongo(e)
	}
	defer cursor.Close(ctx)
	if cursor.Next(ctx) {
		if e := cursor.Decode(o.Model); e != nil {
			return herror.Mongo(e)
		}
	} else {
		return herror.NotFoundMsg("AggregateOne: pipeline returned no document (!cursor.Next)", "No matching record was found for this query.")
	}
	return nil
}

func (o *Odm) Create() herror.Error {
	if e := o.Model.BeforeCreate(); e != nil {
		return e
	}
	result, e := Mongo.Collection(o.Model.CollectionName()).InsertOne(context.TODO(), o.Model)
	if e != nil {
		return herror.Mongo(e)
	}
	o.Model.SetID(result.InsertedID.(bson.ObjectID))
	return o.Model.AfterCreate()
}

func (o *Odm) CreateBy(validator any) herror.Error {
	if e := Fill(o.Model, validator); e != nil {
		return e
	}
	return o.Create()
}

// usela solo si tienes pereza.
func (o *Odm) CreateMany(data any) herror.Error {

	v := reflect.ValueOf(data)

	if v.Kind() != reflect.Slice {
		return herror.InternalServerErrorMsg("CreateMany: value must be a slice of models", "Bulk insert is only available for a list of records. Please check the data you are sending.")
	}
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i).Interface()
		if e := elem.(OdmModel).BeforeCreate(); e != nil {
			return e
		}
	}
	collection := Mongo.Collection(o.Model.CollectionName())
	result, e := collection.InsertMany(context.TODO(), data)
	if e != nil {
		return herror.Mongo(e)
	}
	he := []error{}
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i).Interface()
		elem.(OdmModel).SetID(result.InsertedIDs[i].(bson.ObjectID))
		if e := elem.(OdmModel).AfterCreate(); e != nil {
			he = append(he, e)
		}
	}
	if len(he) > 0 {
		return herror.InternalServerErrorMsg(he, "The records were inserted, but one or more post-save steps failed. Please try again or contact support if this continues.")
	}
	return nil
}

func (o *Odm) Update() herror.Error {
	if e := o.Model.BeforeUpdate(); e != nil {
		return e
	}
	filter := bson.D{bson.E{Key: "_id", Value: o.Model.GetID()}}
	update := bson.D{bson.E{Key: "$set", Value: o.Model}}

	result, e := Mongo.Collection(o.Model.CollectionName()).UpdateOne(context.TODO(), filter, update)
	if e != nil {
		return herror.Mongo(e)
	}
	if result.MatchedCount == 0 {
		return herror.NotFoundMsg("Update: no document matched MatchedCount==0", "The record you are trying to update does not exist or was already removed.")
	}

	if result.ModifiedCount == 0 {
		return herror.ConflictMsg("Update: ModifiedCount==0 (no fields changed)", "No changes were saved. The stored data may already match what you sent.")
	}
	return o.Model.AfterUpdate()
}

func (o *Odm) UpdateBy(validator any) herror.Error {

	if e := FillDirty(o.Model, validator); e != nil {
		return e
	}
	return o.Update()
}

// OjO no usa el hook BeforeUpdate

func (o *Odm) Delete() herror.Error {
	if e := o.Model.BeforeDelete(); e != nil {
		return e
	}

	filter := bson.D{bson.E{Key: "_id", Value: o.Model.GetID()}}

	result, e := Mongo.Collection(o.Model.CollectionName()).DeleteOne(context.TODO(), filter)
	if e != nil {
		return herror.Mongo(e)
	}
	if result.DeletedCount == 0 {
		return herror.ConflictMsg("Delete: DeletedCount==0", "The record could not be deleted. It may have already been removed.")
	}
	return o.Model.AfterDelete()
}
