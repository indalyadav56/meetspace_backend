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
    secretKey = os.Getenv("JWT_SECRET_KEY")
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
    ).SignedString([]byte(secretKey))
        
    
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
      ).SignedString([]byte(secretKey))

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
        return []byte(secretKey), nil
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

// RefreshToken checks if token is expired, generates a new one
func (ts *TokenService) RefreshToken(oldToken string) (string, error) {
    // Parse the token
    token, err := jwt.ParseWithClaims(oldToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    fmt.Println("token:->", token)

    // Invalid token
    if err != nil { 
        return "", err
    }

    // // Token is valid, get claims
    // claims, ok := token.Claims.(*Claims)
    // if !ok {
    //     return "", err
    // }
  
    // // Issued at time is older than expiry time
    // if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
    //     return ts.GenerateToken(claims.u) 
    // }

    // Token still valid, return old token
    return oldToken, nil
}

func RotateToken(tokenString string) (string, error) {
    // token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
    //     return secretKey, nil
    // })
    // if err != nil {
    //     return "", err
    // }

    // claims, ok := token.Claims.(*Claims) 
    // if ok && token.Valid { 
    //     claims.ExpiresAt = time.Now().Add(time.Minute * 15).Unix()
    //     return generateToken(claims.Username)
    // } else {
    //     return "", err
    // }
    return "new token", nil
}