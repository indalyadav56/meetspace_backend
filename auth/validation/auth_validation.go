package validation

import (
	"fmt"
)

func ValidateRegisterRequest() error {
	fmt.Println("Custom registration request validation.")
	return  nil
}


type CustomValidator struct {}

func (cv *CustomValidator) Validate() error {
    // custom validation logic
    return nil 
}