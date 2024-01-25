package handlers

import (
	"errors"
	"meetspace_backend/user/service_factory"
	"meetspace_backend/user/types"
	"meetspace_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateUserHandler godoc
//	@Summary		User create
//	@Tags			User
//	@Produce		json
// @Param user body types.CreateUserData true "User create details"
//	@Router			/v1/users [post]
func CreateUserHandler(c *gin.Context) {
	var reqBody types.CreateUserData
	utils.BindJsonData(c, &reqBody)

	userService := service_factory.GetUserService()
	resp, _ := userService.CreateUser(reqBody)
	
	c.JSON(200, resp)
	return
}

// GetUserByID godoc
//	@Summary		User create
//	@Tags			User
//	@Produce		json
// @Param user body types.CreateUserData true "User create details"
//	@Router			/v1/users/{id} [get]
func GetUserByID(c *gin.Context) {
	userId := c.Param("userId")
	
	userService := service_factory.GetUserService()
	resp := userService.GetUserByID(userId)
	
	c.JSON(resp.StatusCode, resp)
	return
}

// GetUserByID godoc
//	@Summary		User create
//	@Tags			User
//	@Produce		json
// @Param user body types.CreateUserData true "User create details"
//	@Router			/v1/users [get]
func GetAllUsers(c *gin.Context) {
	email := c.Query("email")
	userService := service_factory.GetUserService()
	users, err := userService.GetAllUsers(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, utils.NotFoundErrorResponse("not found", "user not found"))
			return
		}
		resp := utils.ErrorResponse("error", err.Error())
		c.JSON(resp.StatusCode, resp)
		return
	}
	
	resp := utils.SuccessResponse("success", users)
	c.JSON(resp.StatusCode, resp)
	return
}

// GetUserByID godoc
//	@Summary		User create
//	@Tags			User
//	@Produce		json
// @Param user body types.CreateUserData true "User create details"
//	@Router			/v1/users [put]
func UpdateUser(c *gin.Context) {
	currentUser, exists := utils.GetUserFromContext(c)
	
	if !exists{
		resp := utils.ErrorResponse("invalid user!", nil)
		c.JSON(resp.StatusCode, resp)
		return
	}

	var reqBody types.UpdateUserData
	
	if err := utils.BindJsonData(c, &reqBody); err != nil {
		resp:= utils.ErrorResponse("Invalid JSON", nil)
		c.JSON(resp.StatusCode, resp)
		return
	}

	file, _ := c.FormFile("profile_pic")

	reqBody.ProfilePic = file
	
	userService := service_factory.GetUserService()
	response := userService.UpdateUser(currentUser.ID.String(), reqBody)
	
	c.JSON(response.StatusCode, response)
	return
}

// CheckUserEmail godoc
//	@Summary		User create
//	@Tags			User
//	@Produce		json
// @Param user body types.CreateUserData true "User create details"
//	@Router			/v1/user/check-email [get]
func CheckUserEmail(c *gin.Context) {
	email := c.Query("email")
	
	userService := service_factory.GetUserService()
	user, err := userService.GetUserByEmail(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, utils.NotFoundErrorResponse("not found", "user not found"))
			return
		}
		resp := utils.ErrorResponse("error", err.Error())
		c.JSON(resp.StatusCode, resp)
		return
	}
	
	resp := utils.SuccessResponse("success", user)
	
	c.JSON(resp.StatusCode, resp)
	return
}
