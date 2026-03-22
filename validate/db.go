package validate

import (
	"context"
	"reflect"
	"time"

	"github.com/donbarrigon/utils/db"
	"github.com/donbarrigon/utils/str"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Unique validates that a value is unique in the specified collection and field
// Params: [collection, field, excludeId]
func Unique(value reflect.Value, params ...string) (string, str.Placeholder, bool) {

	ph := str.Placeholder{}
	if len(params) < 3 {
		return "The collection, field, and exclude id parameters are required.", ph, true
	}

	col := params[0]
	field := params[1]
	excludeId := params[2]

	// Build the filter
	filter := bson.M{field: value.Interface()}

	// Try to convert to ObjectID if it looks like one
	if oid, err := bson.ObjectIDFromHex(excludeId); err == nil {
		filter["_id"] = bson.M{"$ne": oid}
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if document exists
	var result map[string]any
	err := db.Col(col).FindOne(ctx, filter).Decode(&result)

	if err != nil {
		// If no document found, value is unique
		if err == mongo.ErrNoDocuments {
			return "", nil, false
		}

		// Other errors (connection, timeout, etc.)
		return "Unable to verify uniqueness for :field: " + err.Error(), ph, true
	}

	// Document found, value is not unique
	// ph["value"] = value.Interface()
	return "The :field has already been taken.", ph, true
}

// Exists validates that a value exists in the specified collection and field
// Params: [collection, field]
func Exists(value reflect.Value, params ...string) (string, str.Placeholder, bool) {
	if len(params) < 2 {
		return "The collection and field parameters are required.", str.Placeholder{}, true
	}

	col := params[0]
	field := params[1]
	ph := str.Placeholder{{Key: "collection", Value: col}}

	// Skip validation if value is nil or empty
	if !value.IsValid() || value.IsZero() {
		return "The :field field is required.", ph, true
	}

	// Build the filter
	var filter bson.M
	if field == "_id" || field == "id" {
		switch v := value.Interface().(type) {
		case string:
			// Try to convert string to ObjectID
			if oid, e := bson.ObjectIDFromHex(v); e == nil {
				filter = bson.M{"_id": oid}
			} else {
				filter = bson.M{"_id": v}
			}
		case bson.ObjectID:
			filter = bson.M{"_id": v}
		default:
			filter = bson.M{"_id": v}
		}
	} else {
		filter = bson.M{field: value.Interface()}
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if document exists
	var result map[string]any
	e := db.Col(col).FindOne(ctx, filter).Decode(&result)

	if e != nil {
		// If no document found, value doesn't exist
		if e == mongo.ErrNoDocuments {
			// ph["value"] = fmt.Sprintf("%v", value.Interface())
			return "The selected :field is invalid.", ph, true
		}

		// Other errors (connection, timeout, etc.)
		return "Unable to verify existence for :field: " + e.Error(), ph, true
	}

	// Document found, value exists
	return "", nil, false
}

// NotExists validates that a value does NOT exist in the specified collection and field
// Params: [collection, field]
func NotExists(value reflect.Value, params ...string) (string, str.Placeholder, bool) {
	if len(params) < 2 {
		return "The collection and field parameters are required.", str.Placeholder{}, true
	}

	col := params[0]
	field := params[1]
	ph := str.Placeholder{{Key: "collection", Value: col}}

	// Skip validation if value is nil or empty
	if !value.IsValid() || value.IsZero() {
		return "", nil, false // No error if empty (optional field)
	}

	// Build the filter
	var filter bson.M
	if field == "_id" || field == "id" {
		switch v := value.Interface().(type) {
		case string:
			// Try to convert string to ObjectID
			if oid, err := bson.ObjectIDFromHex(v); err == nil {
				filter = bson.M{"_id": oid}
			} else {
				filter = bson.M{"_id": v}
			}
		case bson.ObjectID:
			filter = bson.M{"_id": v}
		default:
			filter = bson.M{"_id": v}
		}
	} else {
		filter = bson.M{field: value.Interface()}
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if document exists
	var result map[string]any
	err := db.Col(col).FindOne(ctx, filter).Decode(&result)

	if err != nil {
		// If no document found, value doesn't exist (which is what we want)
		if err == mongo.ErrNoDocuments {
			return "", nil, false // No error, validation passed
		}

		// Other errors (connection, timeout, etc.)
		return "Unable to verify availability for :field: " + err.Error(), ph, true
	}

	// Document found, value already exists (validation failed)
	return "The :field has already been taken.", ph, true
}
