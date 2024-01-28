package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
    Field string `json:"field"`
    Message   string `json:"message"`
}


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
func getErrorMsg(fe validator.FieldError) string {
    switch fe.Tag() {
        case "required":
            return "This field is required"
        case "lte":
            return "Should be less than " + fe.Param()
        case "gte":
            return "Should be greater than " + fe.Param()
    }
    return "Unknown error"
}

func HandleValidationError(err error, structToValidate interface{}) []map[string]interface{} {
    var errorList []map[string]interface{}
    
    if validationErrors, ok := err.(validator.ValidationErrors); ok {
        t := reflect.TypeOf(structToValidate)
        
        for _, e := range validationErrors {
            field, _ := t.FieldByName(e.Field())
            jsonTag := field.Tag.Get("json")
            errorList = append(errorList, map[string]interface{}{
                "field": jsonTag,
                "message": "this field is " + e.Tag(),
            })
        }
    }
    
    return errorList
}


// if err := validate.Struct(req); err != nil {
//     data := utils.HandleValidationError(err, req)
//     // var ve validator.ValidationErrors
//     // if errors.As(err, &ve) {
//     //     out := make([]ErrorMsg, len(ve))
//     //     for i, fe := range ve {
//     //         out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
//     //     }
//     // 	// resp := utils.ErrorResponse("Invalid Data", out)
//     // 	c.JSON(http.StatusBadRequest, out)
//     // 	return
//     // }
//     c.JSON(400, data)
//     return
// }