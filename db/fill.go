package db

import (
	"errors"
	"reflect"
	"strings"

	"github.com/donbarrigon/utils/herror"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Fill rellena los campos de model con los valores de request.
// Usa el tag bson como clave pero no compara valores previos.
// Actualiza el model con los valores nuevos de request.
// @return error
func Fill(model any, request any) herror.Error {

	modelValue := reflect.ValueOf(model)
	requestValue := reflect.ValueOf(request)

	if modelValue.Kind() != reflect.Ptr || requestValue.Kind() != reflect.Ptr {
		return herror.InternalServerError(errors.New("El modelo y el request deben ser punteros."))
	}

	modelValue = modelValue.Elem()
	requestValue = requestValue.Elem()

	if modelValue.Kind() != reflect.Struct || requestValue.Kind() != reflect.Struct {
		return herror.InternalServerError(errors.New("El modelo y el request deben ser estructuras."))
	}

	requestType := requestValue.Type()

	for i := 0; i < requestType.NumField(); i++ {
		requestField := requestType.Field(i)
		fieldName := requestField.Name
		requestFieldValue := requestValue.Field(i)

		if !requestField.IsExported() {
			continue
		}

		modelField := modelValue.FieldByName(fieldName)
		if !modelField.IsValid() || !modelField.CanSet() {
			continue
		}

		if requestFieldValue.IsZero() {
			continue
		}

		var newValue reflect.Value
		if modelField.Type() == reflect.TypeOf(bson.ObjectID{}) &&
			requestFieldValue.Kind() == reflect.String {

			oid, e := bson.ObjectIDFromHex(requestFieldValue.String())
			if e != nil {
				continue
			}
			newValue = reflect.ValueOf(oid)

		} else if requestFieldValue.Type().AssignableTo(modelField.Type()) {
			newValue = requestFieldValue
		} else if requestFieldValue.Type().ConvertibleTo(modelField.Type()) {
			newValue = requestFieldValue.Convert(modelField.Type())
		} else {
			continue
		}

		// Actualizamos el modelo directamente sin comparar
		modelField.Set(newValue)
	}

	return nil
}

// Fill compara model con request y retorna los valores antiguos y los nuevos.
// Usa el tag bson como clave.
// Además actualiza el model con los valores nuevos.
// @return error
func Filld(model OdmModel, request any) herror.Error {

	original := map[string]any{}
	dirty := map[string]any{}

	modelValue := reflect.ValueOf(model)
	requestValue := reflect.ValueOf(request)

	if modelValue.Kind() != reflect.Ptr || requestValue.Kind() != reflect.Ptr {
		model.SetOriginal(original)
		model.SetDirty(dirty)
		return herror.InternalServerError(errors.New("model and request must be pointers"))
	}

	modelValue = modelValue.Elem()
	requestValue = requestValue.Elem()

	if modelValue.Kind() != reflect.Struct || requestValue.Kind() != reflect.Struct {
		model.SetOriginal(original)
		model.SetDirty(dirty)
		return herror.InternalServerError(errors.New("model and request must be structs"))
	}

	modelType := modelValue.Type()
	requestType := requestValue.Type()

	for i := 0; i < requestType.NumField(); i++ {
		requestField := requestType.Field(i)
		fieldName := requestField.Name
		requestFieldValue := requestValue.Field(i)

		if !requestField.IsExported() {
			continue
		}

		modelField := modelValue.FieldByName(fieldName)
		if !modelField.IsValid() || !modelField.CanSet() {
			continue
		}

		if requestFieldValue.IsZero() {
			continue
		}

		var newValue reflect.Value
		if modelField.Type() == reflect.TypeOf(bson.ObjectID{}) &&
			requestFieldValue.Kind() == reflect.String {

			oid, err := bson.ObjectIDFromHex(requestFieldValue.String())
			if err != nil {
				continue
			}
			newValue = reflect.ValueOf(oid)

		} else if requestFieldValue.Type().AssignableTo(modelField.Type()) {
			newValue = requestFieldValue
		} else if requestFieldValue.Type().ConvertibleTo(modelField.Type()) {
			newValue = requestFieldValue.Convert(modelField.Type())
		} else {
			continue
		}

		currentValue := modelField
		if !reflect.DeepEqual(currentValue.Interface(), newValue.Interface()) {
			// sacar el tag bson
			modelFieldStruct, found := modelType.FieldByName(fieldName)
			if !found {
				continue
			}

			bsonTag := modelFieldStruct.Tag.Get("bson")
			if bsonTag == "" {
				bsonTag = strings.ToLower(fieldName)
			} else if commaIndex := strings.Index(bsonTag, ","); commaIndex != -1 {
				bsonTag = bsonTag[:commaIndex]
			}

			// Guardamos viejo y nuevo
			original[bsonTag] = currentValue.Interface()
			dirty[bsonTag] = newValue.Interface()

			// Actualizamos el modelo
			modelField.Set(newValue)
		}
	}
	model.SetOriginal(original)
	model.SetDirty(dirty)
	return nil
}
