package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"os"
)

func Encrypt(plaintext string) (string, error) {
	text := []byte(plaintext)
	key := []byte(os.Getenv("Secret"))
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Println(err)
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println(err)
		return "", err
	}
	gcm.Seal(nonce, nonce, text, nil)
	return base64.RawURLEncoding.EncodeToString(gcm.Seal(nonce, nonce, text, nil)), nil
}

func Decrypt(ciphertext string) (string, error) {
	key := []byte(os.Getenv("Secret"))
	temp, _ := base64.RawURLEncoding.DecodeString(ciphertext)
	ciphertext = string(temp)
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Println(err)
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Println(err)
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		log.Println(err)
		return "", err
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		log.Println(err)
	}
	return string(plaintext), nil
}

func GetHash(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(hash), nil
}

func CompareHashAndPassword(hash string, password string) bool {
	byteHash := []byte(hash)
	bytePassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
