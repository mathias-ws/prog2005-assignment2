package hashing

import (
	"assignment-2/internal/custom_errors"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

// HashString takes a string input and hashes it.
func HashString(stringToHash string) (string, error) {
	message := []byte(stringToHash)

	// Secret only for testing purposes, will change for production use.
	secret := []byte("detteerenhemlighetsomfåklarerågjette")

	hash := hmac.New(sha256.New, secret)

	_, err := hash.Write(message)

	if err != nil {
		log.Printf("Error hashing message: %v", err)
		return "", custom_errors.GetHashingError()
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
