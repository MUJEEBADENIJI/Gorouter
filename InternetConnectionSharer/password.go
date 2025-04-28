package main

import (
    "golang.org/x/crypto/bcrypt"
)

// HashPassword returns the hashed password
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

// CheckPassword checks if the provided password matches the hashed password
func CheckPassword(password string, hashedPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}
