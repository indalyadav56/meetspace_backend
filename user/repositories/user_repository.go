package repositories

import (
	"errors"
	"fmt"
	authModel "meetspace_backend/auth/models"
	"meetspace_backend/config"
	"meetspace_backend/user/models"
	"meetspace_backend/user/types"
	"meetspace_backend/utils"

	"gorm.io/gorm"
)


type UserRepository struct {
	db *gorm.DB
}


func NewUserRepository() *UserRepository {
	return &UserRepository{
		db:      config.DB,
	}
}


func (userRepo *UserRepository) CreateRecord(user models.User) (*models.User, error) {
    if err := userRepo.db.Create(&user).Error; err != nil {
        fmt.Println("User create repos error", err)
        return nil, err
    }

    // update user_id in verification
    var verification authModel.Verification
	userRepo.db.Where("email = ? OR phone_number = ? AND is_verified = ?", user.Email, user.PhoneNumber, true).First(&verification)
	
	if verification.Email != "" {
		verification.UserID = user.ID
		userRepo.db.Save(&verification) 
	}
  

    return &user, nil
}


func (userRepo *UserRepository) GetUserByID(userID string) (models.User, error) {
    var user models.User
    userRepo.db.Where("id = ?", userID).First(&user)
    return user, nil
}


func (userRepo *UserRepository) GetUserByEmail(email string) (models.User, error) {
    var user models.User
    err := userRepo.db.Where("email=?", email).First(&user).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return user, err
    }
    return user, nil
}


func (userRepo *UserRepository) GetAllUserRecord(email string) ([]models.User, error) {
    var users []models.User
   
    if email != "" {
        err := userRepo.db.Where("email=?", email).First(&users).Error
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, err
        }
        return users, nil
    }
    
    userRepo.db.Find(&users).Order("created_at DESC")
    return users, nil
}


func (userRepo *UserRepository) GetUsersByClientId(clientId string)([]models.User, error) {
    var users []models.User
    userRepo.db.Where("client_id = ? AND role <> ?", clientId, "admin").Find(&users).Order("updated_at DESC")
    return users, nil
}


func (userRepo *UserRepository) UpdateUser(userId string, updateUser map[string]interface{}) (types.UserResponse, error) {
    var resp types.UserResponse
    
    data, _ := utils.RemoveKeysNotInStruct(models.User{}, updateUser)
    
    if err := userRepo.db.Model(&resp).Where("id=?", userId).Updates(data).First(&resp).Error; err != nil {
        return resp, err
    }
    
    return resp, nil
}
