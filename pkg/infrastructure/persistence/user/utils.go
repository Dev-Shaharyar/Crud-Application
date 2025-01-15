package user

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
)

// MongoDBModelToBson converts a MongoDB model to a bson.D document.
// It takes a model, which can be a struct or a pointer to a struct, and returns a bson.D.
// If the provided model is a pointer, it dereferences it to access the underlying struct.
// The function panics if the model is not a struct or a pointer to a struct.
func MongoDBModelToBson(model interface{}) bson.D {
	val := reflect.ValueOf(model)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Ensure the provided model is a struct
	if val.Kind() != reflect.Struct {
		panic("CreateMockResponseFromModel: model must be a struct or a pointer to a struct")
	}

	var response bson.D
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldName := field.Tag.Get("bson")
		if fieldName == "" {
			fieldName = field.Name
		}
		fieldValue := val.Field(i).Interface()
		response = append(response, bson.E{Key: fieldName, Value: fieldValue})
	}

	return response
}
