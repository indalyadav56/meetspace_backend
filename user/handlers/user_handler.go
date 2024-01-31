package handlers

import (
	"errors"
	"meetspace_backend/config"
	"meetspace_backend/user/services"
	"meetspace_backend/user/types"
	"meetspace_backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler{
	return &UserHandler{
		UserService:  userService,
	}
}

// CreateUserHandler godoc
//	@Summary		User create
//	@Tags			User
//	@Produce		json
// @Param user body types.CreateUserData true "User create details"
//	@Router			/v1/users [post]
func (handler *UserHandler) CreateUserHandler(c *gin.Context) {
	var reqBody types.CreateUserData
	utils.BindJsonData(c, &reqBody)

	resp, _ := config.UserService.CreateUser(reqBody)
	
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
	
	resp := config.UserService.GetUserByID(userId)
	
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

	users, err := config.UserService.GetAllUsers(email)

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
	
	response := config.UserService.UpdateUser(currentUser.ID.String(), reqBody)
	
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
	
	user, err := config.UserService.GetUserByEmail(email)

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