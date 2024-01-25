package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// GetValidator ensures a single validator instance is created and used throughout the application
func GetValidator() *validator.Validate {
    if validate == nil {
        validate = validator.New()
    }
    return validate
}

// ValidateStruct validates a struct using the singleton validator
func ValidateStruct(s interface{}) error {
    if err := GetValidator().Struct(s); err != nil {
        return err
    }
    return nil
}

func HandleValidationError(err error, structToValidate interface{}) []map[string]interface{} {
    var errorList []map[string]interface{}
    
    if validationErrors, ok := err.(validator.ValidationErrors); ok {
        t := reflect.TypeOf(structToValidate)
        
        for _, e := range validationErrors {
            field, _ := t.FieldByName(e.Field())
            jsonTag := field.Tag.Get("json")

            errorList = append(errorList, map[string]interface{}{
                jsonTag: e.Tag() + " validation failed for " + jsonTag,
            })
        }
        
    }
    
    return errorList
}