package utils

import "golang.org/x/crypto/bcrypt"

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