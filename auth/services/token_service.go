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
    accessExpiration = 1200 * time.Hour
    RefreshExpiration = 2004 * time.Hour
    SigningAlgorithm = "HS256" 
    secretKey = os.Getenv("JWT_SECRET_KEY")
)
const (
    ACCESS_TOKEN_TYPE = "access"
    REFRESH_TOKEN_TYPE = "refresh"
)


type TokenClaims struct {
    TokenType string `json:"token_type"`
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
        ACCESS_TOKEN_TYPE,
        userID, 
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(accessExpiration).Unix(),
        },
    }
    // Create tokens with claims
    accessToken, _ := jwt.NewWithClaims(
        jwt.SigningMethodHS256, 
        accessClaims,
    ).SignedString([]byte(secretKey))
    
    // Refresh token claims
    refreshClaims := &TokenClaims{
        REFRESH_TOKEN_TYPE,
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
        ACCESS_TOKEN_TYPE: accessToken,
        REFRESH_TOKEN_TYPE: refreshToken,
    }
    
    return tokenData, nil
}

// verify token
func (ts *TokenService) VerifyToken(tokenString string, expectedTokenType string) (map[string]interface{}, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Validate method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
          return nil, fmt.Errorf("Unexpected signing method") 
        }
        // Return key for validation
        return []byte(secretKey), nil
    })
    if err != nil {
        return nil, err
    }

    // Extract the token claims
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return nil, errors.New("Invalid token")
    }

    // Verify token type
    tokenType, ok := claims["token_type"].(string)
    if !ok || tokenType != expectedTokenType {
        return nil, fmt.Errorf("Invalid token type: expected %s", expectedTokenType)
    }
    return claims, nil
}

// RefreshToken checks if token is expired, generates a new one
func (ts *TokenService) RefreshToken(refreshToken string) (string, error) {
    token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
        // Validate method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
          return nil, fmt.Errorf("Unexpected signing method") 
        }
        // Return key for validation
        return []byte(secretKey), nil
    })
    // Extract the token claims
    claims, ok := token.Claims.(jwt.MapClaims)
    fmt.Println("claims", claims)
    if !ok || !token.Valid {
        return "", errors.New("Invalid token")
    }
    userId := claims["user_id"].(string)
    // Access token claims  
    accessTokenClaims := &TokenClaims{ 
        ACCESS_TOKEN_TYPE,
        userId, 
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(accessExpiration).Unix(),
        },
    }

    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
    accessTokenString, err := accessToken.SignedString([]byte(secretKey))
    if err != nil {
        return "", fmt.Errorf("Failed to generate access token")
    }
    return accessTokenString, nil
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