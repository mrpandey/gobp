package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	secretLen := 20

	secretBytes := make([]byte, secretLen)
	_, err := rand.Read(secretBytes)
	if err != nil {
		fmt.Println("error reading random token:", err)
		return
	}

	secretBase64 := base64.RawURLEncoding.EncodeToString(secretBytes)
	fmt.Println("secret:", secretBase64)

	hashBytes, err := bcrypt.GenerateFromPassword(secretBytes, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("error hashing the secret:", err)
		return
	}

	hashBase64 := base64.RawURLEncoding.EncodeToString(hashBytes)
	fmt.Println("hash:", hashBase64)
}
