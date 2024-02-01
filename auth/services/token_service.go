package services

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


  type TokenClaims struct {
    UserID string `json:"user_id"`
    jwt.StandardClaims
  }

type TokenService struct {
    signingKey []byte
    expiry     time.Duration
}

// NewTokenService creates a new token service
func NewTokenService() *TokenService {
    return &TokenService{
        signingKey: []byte("signingKey"),
        expiry:     time.Minute * time.Duration(10),
    }
}

// GenerateToken generates new JWT tokens
func (ts *TokenService) GenerateToken(userID string) (map[string]string, error) {
    tokenData := make(map[string]string)

    // Access token claims  
    accessClaims := &TokenClaims{ 
        userID, 
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(AccessExpiration).Unix(),
        },
    }
    // Create tokens with claims
    accessToken, _ := jwt.NewWithClaims(
        jwt.SigningMethodHS256, 
        accessClaims,
    ).SignedString([]byte(SecretKey))
        
    
    // Refresh token claims
    refreshClaims := &TokenClaims{
        userID,
        jwt.StandardClaims{
        ExpiresAt: time.Now().Add(RefreshExpiration).Unix(),
        },
    }
    refreshToken, _ := jwt.NewWithClaims(
        jwt.SigningMethodHS256,
        refreshClaims,  
      ).SignedString([]byte(SecretKey))

    tokenData = map[string]string{
        "access": accessToken,
        "refresh": refreshToken,
    }
    
    return tokenData, nil
}

func (ts *TokenService) VerifyToken(tokenString string) (string, error) {
  
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
  
