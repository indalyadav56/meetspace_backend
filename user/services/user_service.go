package services

import (
	"encoding/json"
	"meetspace_backend/user/models"
	"meetspace_backend/user/repositories"
	"meetspace_backend/user/types"
	"meetspace_backend/utils"
	"mime/multipart"
	"strings"

	"github.com/google/uuid"
)

type UserService struct {
    UserRepository *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}


func (us *UserService) CreateUser(userData types.CreateUserData) (*models.User, error) {
	userObj := models.User{
		FirstName: userData.FirstName,
		LastName: userData.LastName,
		Email: userData.Email,
		Password: userData.Password,
	}

	if userData.Role != ""{
		userObj.Role = userData.Role
	}

    user, err :=  us.UserRepository.CreateRecord(userObj)
	
	if err != nil {
		return nil, err
	}

	return user, nil
}


func (us *UserService) GetUserByID(userID string) *utils.Response {
	if strings.TrimSpace(userID) == "" {
		return utils.ErrorResponse("user id cannot be blank!", nil)
	}

	_, err := uuid.Parse(userID)
	if err != nil {
		return utils.ErrorResponse("user id is invalid", nil)
	}

    user, err :=  us.UserRepository.GetUserByID(userID)
	if err != nil {
		return utils.ErrorResponse(err.Error(), nil)
	}

	return utils.SuccessResponse("successfully get data!", user)
}

func (us *UserService) GetUserByEmail(email string) (models.User, error) {
    user, err :=  us.UserRepository.GetUserByEmail(email)
	return user, err
}

func (us *UserService) GetAllUsers(email string) ([]models.User, error) {
    users, err :=  us.UserRepository.GetAllUserRecord(email)
	
	if err != nil {
		return nil, err
	}
	
	return users, err
}

func (us *UserService) UpdateUser(userId string, updateData types.UpdateUserData) utils.Response{
	// err := validators.ValidateUpdateUserData(&updateData)

	// if err != nil {
	// 	return utils.ErrorResponse("error", []interface{}{})
	// }

	mapData := map[string]interface{}{
		"first_name": updateData.FirstName,
		"last_name": updateData.LastName,
	}

	// if err != nil {
	// 	return utils.ErrorResponse("error while updating user.", nil)
	// }
	
	if updateData.ProfilePic != nil{
		profilePicData := us.UploadUserProfilePic(updateData.ProfilePic, userId)
		mapData["profile_pic"] = profilePicData
	}
	
	userData, _ := us.UserRepository.UpdateUserByID(userId, mapData)

	userResponse := types.UserResponse{
		ID: userData.ID,
		FirstName: userData.FirstName,
		LastName: userData.LastName,
		Email: userData.Email,
		IsActive: userData.IsActive,
		ProfilePic: userData.ProfilePic,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	return *utils.SuccessResponse(
		"User updated successfully", 
		userResponse,
	)
	
}

func (us *UserService) UploadUserProfilePic(file  *multipart.FileHeader, userId string) string{
	tempFileName  := userId[:8]+".jpg"
	tempFilePath := "uploads/" + userId+"/" + "profile/" + tempFileName
	
	if err := utils.SaveFile(file, tempFilePath); err != nil {
		return err.Error()
	}

	type ProfilePicData struct {
		OriginalName string `json:"original_name"`
		TempName string `json:"temp_name"`
		Metadata map[string]interface{} `json:"metadata"`
	}

	var profileData ProfilePicData

	profileData.OriginalName = file.Filename
	profileData.TempName = tempFileName

	jsonData, _ := json.Marshal(profileData)

	return string(jsonData)
}