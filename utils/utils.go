package utils

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"meetspace_backend/user/models"
	"mime/multipart"
	"os"
	"path/filepath"
	"reflect"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/mitchellh/mapstructure"
)

// EncryptPassword hashes a password using bcrypt.
func EncryptPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// ComparePassword compares a plain-text password with a hashed password.
func ComparePassword(hashedPassword, rawPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
    return err == nil
}

// get user from gin context
func GetUserFromContext(c *gin.Context) (*models.User, bool) {
	currentUser, exists := c.Get("user")
	if !exists {
		return nil, false
	}

	user, ok := currentUser.(models.User)
	if !ok {
		return nil, false
	}

	return &user, true
}

// StructToString converts a struct to a string
func StructToString(data interface{}) (string, error) {
	if reflect.TypeOf(data).Kind() != reflect.Struct {
		return "", fmt.Errorf("input data is not a struct")
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

// StringToStruct converts a string to a struct
func StringToStruct(str string, result interface{}) error {
	err := json.Unmarshal([]byte(str), &result)
	if err != nil {
		return err
	}
	return nil
}

func BindJsonData(c *gin.Context, target interface{}) error {
    if err := c.ShouldBind(&target); err != nil {
        return err
    }
    return nil
}

// SaveUploadedFile uploads the form file to specific dst.
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(dst), 0750); err != nil {
		return err
	}

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// check the key is in struct 
func HasKey(myStruct interface{}, key string) bool {
    structVal := reflect.ValueOf(myStruct)
    structType := structVal.Type()

    // Check for key in struct fields:
    _, found := structType.FieldByName(key)
    // fmt.Println("field", field)
    if found {
        return true
    }

    // Check for key in JSON tags:
    for i := 0; i < structVal.NumField(); i++ {
        field := structType.Field(i)
        jsonTag := field.Tag.Get("json")
        if jsonTag == key {
            return true
        }
    }

    return false
}

// remove the given map data key if not exists in given struct
func RemoveKeysNotInStruct(strcutData interface{}, mapData map[string]interface{})(map[string]interface{}, error){
	for key := range mapData {
        isFound := HasKey(strcutData, key)
       
		if !isFound{
            delete(mapData, key)
        }
    }
	return mapData, nil
}

func StructToMap(s interface{}) (map[string]interface{}, error) {
    result := make(map[string]interface{})
    decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
        Result:  &result,
        TagName: "json",
    })
    if err != nil {
        return nil, err
    }
    err = decoder.Decode(s)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func GenerateOTP() string {
    num, _ := rand.Int(rand.Reader, big.NewInt(1000000))
    return fmt.Sprintf("%06d", num)
}