package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	// 注册时
	userPassword := "Dazhou520"
	hashedPassword, err := hashPassword(userPassword)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}
	fmt.Println("Hashed Password:", hashedPassword)

	// 登录时
	providedPassword := "Dazhou520"
	if checkPasswordHash(providedPassword, hashedPassword) {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Incorrect password.")
	}
}
