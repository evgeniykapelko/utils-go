package cipherer

import (
	"encoding/base64"
	"errors"
)

func Cipher(rowString, secret string) (string, error) {
	if len(secret) == 0 {
		return "", errors.New("row string is empty")
	}

	encrypedBytes, err := process([]byte(rowString), []byte(secret))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encrypedBytes), nil
}

func Decipher(cipheredText, secret string) (string, error) {
	if len(secret) == 0 {
		return "", errors.New("Secret key cannot be empty")
	}

	cipheredBytes, err := base64.StdEncoding.DecodeString(cipheredText)
	if err != nil {
		return "", errors.New("Error decoding ciphered text")
	}

	decryptedBytes, err := process(cipheredBytes, []byte(secret))
	if err != nil {
		return "", err
	}

	return string(decryptedBytes), nil
}

func process(input, secret []byte) ([]byte, error) {
	for i, b := range input {
		input[i] = b ^ secret[i%len(secret)] // 0..len(secret)
	}

	return input, nil
}
