package repositories

import (
	"errors"
	"fmt"
	"meetspace_backend/user/constants"
	"meetspace_backend/user/models"
	"meetspace_backend/user/types"
	"meetspace_backend/utils"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

// isDuplicateKeyError checks if the error is due to a duplicate key violation
func isDuplicateKeyError(err error) bool {
    // Check the error code or error message to identify duplicate key violation
    // This check may vary depending on the database driver used
    // For PostgreSQL, check if the error code is 23505
    // Example: PostgreSQL error code for duplicate key violation is 23505
    pgError, ok := err.(*pgconn.PgError)
    if ok && pgError.Code == "23505" {
        return true
    }
    return false
}


type UserRepository struct {
	db *gorm.DB
}


func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db:      db,
	}
}


func (userRepo *UserRepository) CreateRecord(user models.User) (*models.User, error) {
    result := userRepo.db.Create(&user)
    if result.Error != nil {
        if isDuplicateKeyError(result.Error) {
            return nil, errors.New(constants.EMAIL_ALREADY_EXISTS_MSG)
        }
    }

    // update user_id in verification
    // var verification authModel.Verification
	// userRepo.db.Where("email = ? OR phone_number = ? AND is_verified = ?", user.Email, user.PhoneNumber, true).First(&verification)
	
	// if verification.Email != "" {
	// 	verification.UserID = user.ID
	// 	userRepo.db.Save(&verification) 
	// }

    return &user, nil
}


func (userRepo *UserRepository) GetUserByID(userID string) (*models.User, error) {
    var user models.User
    err := userRepo.db.Where("id = ?", userID).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}


func (userRepo *UserRepository) GetUserByEmail(email string) (models.User, error) {
    var user models.User
    err := userRepo.db.Where("email=?", email).First(&user).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return user, errors.New(constants.EMAIL_NOT_FOUND)
    }
    return user, nil
}


func (userRepo *UserRepository) GetAllUserRecord(currentUserEmail, email string) ([]models.User, error) {
    var users []models.User
   
    if email != "" {
        err := userRepo.db.Where("email=?", email).First(&users).Error
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, err
        }
        return users, nil
    }
    
    userRepo.db.Not("email = ?", currentUserEmail).Find(&users).Order("created_at DESC")
    return users, nil
}


func (userRepo *UserRepository) GetUsersByClientId(clientId string)([]models.User, error) {
    var users []models.User
    userRepo.db.Where("client_id = ? AND role <> ?", clientId, "admin").Find(&users).Order("updated_at DESC")
    return users, nil
}


func (userRepo *UserRepository) UpdateUserByID(userId string, updateUser map[string]interface{}) (types.UserResponse, error) {
    var resp types.UserResponse
    
    data, _ := utils.RemoveKeysNotInStruct(models.User{}, updateUser)
    
    if err := userRepo.db.Model(&models.User{}).Where("id = ?", userId).Updates(data).First(&resp).Error; err != nil {
        return resp, err
    }
    
    return resp, nil
}

func (userRepo *UserRepository) UpdateUserByEmail(email string, updateUser map[string]interface{}) (types.UserResponse, error) {
    var resp types.UserResponse
    
    fmt.Println("data to be update before :- ", updateUser)
    data, _ := utils.RemoveKeysNotInStruct(models.User{}, updateUser)

    fmt.Println("data to be update", data)
    
    if err := userRepo.db.Model(&models.User{}).Where("email = ?", email).Updates(data).First(&resp).Error; err != nil {
        return resp, err
    }
    
    return resp, nil
}
