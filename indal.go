package main

import (
	"fmt"
)

func StartService(){
	fmt.Println("Hello world!")
}




// type ErrorMsg struct {
//     Field string `json:"field"`
//     Message   string `json:"message"`
// }

// func getErrorMsg(fe validator.FieldError) string {
//     switch fe.Tag() {
//         case "required":
//             return "This field is required"
//         case "lte":
//             return "Should be less than " + fe.Param()
//         case "gte":
//             return "Should be greater than " + fe.Param()
//     }
//     return "Unknown error"
// }

// 	UserRegister godoc
//	@Summary		Register User account
//	@Description	Register User account
//	@Tags			Auth
//	@Produce		json
// 	@Param user body types.RegisterRequest true "User registration details"
//	@Router			/v1/auth/register [post]
// func UserRegister(c *gin.Context){
// 	var req types.RegisterRequest
// 	if err := utils.BindJsonData(c, &req); err != nil {
// 		resp:= utils.ErrorResponse("Invalid JSON", err.Error())
// 		c.JSON(resp.StatusCode, resp)
//         return 
//     }

// 	validate := validator.New()
//     if err := validate.Struct(req); err != nil {
// 		var ve validator.ValidationErrors
//         if errors.As(err, &ve) {
//             out := make([]ErrorMsg, len(ve))
//             for i, fe := range ve {
//                 out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
//             }
// 			resp := utils.ErrorResponse("Invalid Data", out)
// 			c.JSON(http.StatusBadRequest, resp)
// 			return
//         }
		
//     }

// 	user, err := config.AuthService.Register(req)
// 	if err != nil {
// 		return 
// 	}

// 	accessToken, refreshToken, _ := utils.GenerateTokenPair(user.ID.String())

// 	tokenData := map[string]interface{}{
// 		"access": accessToken,
// 		"refresh": refreshToken,
// 	}
// 	resData := types.AuthResponse{
// 		User: user,
// 		Token: tokenData,
// 	}

// 	resp := utils.SuccessResponse(constants.USER_REGISTER_MSG, resData)
// 	c.JSON(resp.StatusCode, resp)
// 	return
// }

// go test -coverprofile=coverage.out ./...

// Then, use the go tool cover command to convert the coverage profile into HTML:
// go tool cover -html=coverage.out -o coverage.html