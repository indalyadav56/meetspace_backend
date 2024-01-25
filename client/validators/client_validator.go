package validators

import (
	"errors"
	"meetspace_backend/user/types"
	"strings"
)

func ValidateUpdateUserData(data *types.UpdateUserData) error {
    // validator := utils.GetValidator()
    trimmedFirstName := strings.TrimSpace(data.FirstName)
    trimmedLastName := strings.TrimSpace(data.LastName)

    if trimmedFirstName == "" && trimmedLastName == "" && data.ProfilePic == nil{
        return errors.New("to update user data field requires")
    }
    
    if trimmedFirstName == "" {
        return errors.New("first name cannot be empty")
    }
    if trimmedFirstName != "" && data.ProfilePic != nil && trimmedLastName == "" {
        return errors.New("Last name cannot be empty")
    }

    return nil
}

    // for i := 0; i < val.NumField(); i++ {
    //     field := val.Type().Field(i)
    //     tag := field.Tag.Get("validate")  
    //     value := val.Field(i).Interface()

    //     fmt.Println("invalid", value)
    //     if field.Type.Field(i) == ""{
    //         fmt.Println("invalid2", field.Name)
    //         if err := validator.Var(value, tag); err != nil {
    //             return err
    //         } 
    //     }

        // if tag != "" {
        //     if err := v.Var(value, tag); err != nil {
        //         return err
        //     }
        // }
    // }