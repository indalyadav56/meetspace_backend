package handlers

import (
	"errors"
	"meetspace_backend/config"
	"meetspace_backend/user/types"
	"meetspace_backend/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateUserHandler godoc
//	@Summary		User create
//	@Tags			User
//	@Produce		json
// @Param user body types.CreateUserData true "User create details"
//	@Router			/v1/users [post]
func CreateUserHandler(c *fiber.Ctx) error {
	var reqBody types.CreateUserData

	c.JSON(reqBody)
	return nil
}

// GetUserByID godoc
//	@Summary		User create
//	@Tags			User
//	@Produce		json
// @Param user body types.CreateUserData true "User create details"
//	@Router			/v1/users/{id} [get]
func GetUserByID(c *fiber.Ctx) error{
	userId := c.Query("userId")
	
	resp := config.UserService.GetUserByID(userId)
	
	c.JSON(resp)
	return nil
}

// GetUserByID godoc
//	@Summary		User create
//	@Tags			User
//	@Produce		json
// @Param user body types.CreateUserData true "User create details"
//	@Router			/v1/users [get]
func GetAllUsers(c *fiber.Ctx) error{
	email := c.Query("email")
	users, err := config.UserService.GetAllUsers(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound).JSON(utils.NotFoundErrorResponse("not found", "user not found"))
			return nil
		}
		resp := utils.ErrorResponse("error", err.Error())
		c.JSON(resp)
		return nil
	}
	
	resp := utils.SuccessResponse("success", users)
	c.JSON(resp)
	return nil
}

// GetUserByID godoc
//	@Summary		User create
//	@Tags			User
//	@Produce		json
// @Param user body types.CreateUserData true "User create details"
//	@Router			/v1/users [put]
func UpdateUser(c *fiber.Ctx) error {
	

	var reqBody types.UpdateUserData
	

	file, _ := c.FormFile("profile_pic")

	reqBody.ProfilePic = file
	
	response := config.UserService.UpdateUser("currentUser.ID.String()", reqBody)
	
	c.JSON(response)
	return nil
}

// CheckUserEmail godoc
//	@Summary		User create
//	@Tags			User
//	@Produce		json
// @Param user body types.CreateUserData true "User create details"
//	@Router			/v1/user/check-email [get]
func CheckUserEmail(c *fiber.Ctx) error {
	email := c.Query("email")
	
	user, err := config.UserService.GetUserByEmail(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(utils.NotFoundErrorResponse("not found", "user not found"))
			return nil
		}
		resp := utils.ErrorResponse("error", err.Error())
		c.JSON(resp)
		return nil
	}
	
	resp := utils.SuccessResponse("success", user)
	
	c.JSON(resp)
	return nil
}
