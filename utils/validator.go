package utils

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
    Field string `json:"field"`
    Message   string `json:"message"`
}

var validate *validator.Validate

func init() {
    validate  =  validator.New()
    validate.RegisterValidation("not_blank", NotBlank)
}

// GetValidator ensures a single validator instance is created and used throughout the application
func GetValidator() *validator.Validate {
    if validate == nil {
        validate = validator.New()
    }
    return validate
}


func NotBlank(fl validator.FieldLevel) bool {
	field := fl.Field()

	switch field.Kind() {
	case reflect.String:
		return len(strings.TrimSpace(field.String())) > 0
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}

func getErrorMsg(fe validator.FieldError) string {
    switch fe.Tag() {
        case "required":
            return "This field is required"
        case "lte":
            return "Should be less than " + fe.Param()
        case "gte":
            return "Should be greater than " + fe.Param()
        case "min":
            return "Should be greater than " + fe.Param()
        case "email":
            return "Email is not valid."
        case "not_blank":
            return "This field should not be blank" + fe.Param()
    }
    return "Unknown error"
}

func ParseError(err error, structToValidate interface{}) interface{} {
    var ve validator.ValidationErrors
    t := reflect.TypeOf(structToValidate)
    
    if errors.As(err, &ve) {
        errData := make([]ErrorMsg, len(ve))
        
        for i, fe := range ve {
            field, _ := t.FieldByName(fe.Field())
            jsonTag := field.Tag.Get("json")
            errData[i] = ErrorMsg{
                Field: jsonTag, 
                Message: getErrorMsg(fe),
            }
        }
    	return errData
    }
    return nil
}
