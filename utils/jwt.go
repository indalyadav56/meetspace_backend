package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// Constants
var (
  AccessExpiration = 1200 * time.Hour
  RefreshExpiration = 2004 * time.Hour
  SigningAlgorithm = "HS256" 
  SecretKey = os.Getenv("JWT_SECRET_KEY")
)

// Custom claims structure
type TokenClaims struct {
  UserID string `json:"user_id"`
  jwt.StandardClaims
}

// Generate access and refresh token pair
func GenerateTokenPair(userId string) (accessToken, refreshToken string, err error) {

  // Access token claims  
  accessClaims := &TokenClaims{ 
    userId, 
    jwt.StandardClaims{
      ExpiresAt: time.Now().Add(AccessExpiration).Unix(),
    },
  }

  // Refresh token claims
  refreshClaims := &TokenClaims{
    userId,
    jwt.StandardClaims{
     ExpiresAt: time.Now().Add(RefreshExpiration).Unix(),
    },
  }

  // Create tokens with claims
  accessToken, err = jwt.NewWithClaims(
    jwt.SigningMethodHS256, 
    accessClaims,
  ).SignedString([]byte(SecretKey))

  if err != nil {
    return  
  }

  refreshToken, err = jwt.NewWithClaims(
    jwt.SigningMethodHS256,
    refreshClaims,  
  ).SignedString([]byte(SecretKey))

  return
}

func VerifyToken(tokenString string) (string, error) {
  
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      // Validate method
      if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected signing method") 
      }
      
      // Return key for validation
      return []byte(SecretKey), nil
    })

  if err != nil {
      fmt.Println("Error:", err)
      return "", err
  }

  if !token.Valid {
      return "", errors.New("Invalid token")
  }

  // Extract the user ID from the token claims
  claims := token.Claims.(jwt.MapClaims)
  userID := claims["user_id"].(string)

  return userID, nil
}


// generate tokens for user
func GenerateUserToken(userId string) (map[string]string, error) {
	accessToken, refreshToken, _ := GenerateTokenPair(userId)

	tokenData := map[string]string{
		"access": accessToken,
		"refresh": refreshToken,
	}

  return tokenData, nil

}

// import (
//   "context"
//   "github.com/go-redis/redis/v8"
// )

// var (
//   // Redis client 
//   rdb = redis.NewClient(&redis.Options{
//       Addr: "localhost:6379", 
//   })

//   // Refresh token ctx key 
//   refreshKey = "refresh_token_" 
// )

// // User struct
// type User struct {
//   ID int
//   Username string
// }

// // SET token in redis 
// func SetRefreshToken(user *User, token string) error {

//   err := rdb.Set(context.Background(), refreshKey+user.ID, token, 0).Err()
//   if err != nil {
//       return err
//   }

//   return nil
// }

// // GET token from redis
// func GetRefreshToken(user *User) (string, error) {

//   token, err := rdb.Get(context.Background(), refreshKey+user.ID).Result()
//   if err != nil {
//       return "", err
//   }

//   return token, nil
// }  

// // DELETE on logout  
// func DeleteRefreshToken(userId int) error {
  
//   err := rdb.Del(context.Background(), refreshKey+userId).Err()
//   if err != nil {
//       return err 
//   }

// import (
//   "context"
//   "github.com/go-redis/redis/v8"
// )

// var (
//   // Redis client 
//   rdb = redis.NewClient(&redis.Options{
//       Addr: "localhost:6379", 
//   })

//   // Refresh token ctx key 
//   refreshKey = "refresh_token_" 
// )

// // User struct
// type User struct {
//   ID int
//   Username string
// }

// // SET token in redis 
// func SetRefreshToken(user *User, token string) error {

//   err := rdb.Set(context.Background(), refreshKey+user.ID, token, 0).Err()
//   if err != nil {
//       return err
//   }

//   return nil
// }

// // GET token from redis
// func GetRefreshToken(user *User) (string, error) {

//   token, err := rdb.Get(context.Background(), refreshKey+user.ID).Result()
//   if err != nil {
//       return "", err
//   }

//   return token, nil
// }  

// // DELETE on logout  
// func DeleteRefreshToken(userId int) error {
  
//   err := rdb.Del(context.Background(), refreshKey+userId).Err()
//   if err != nil {
//       return err 
//   }

//   return nil
// }

//   return nil
// }