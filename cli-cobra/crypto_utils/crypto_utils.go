package crypto_utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
)

type KeyDerivationConfig struct {
	Passphrase []byte
	Salt       []byte
}

func MakeNonce(crypter cipher.AEAD) ([]byte, error) {
	nonce := make([]byte, crypter.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	return nonce, nil
}

func MakeCrypter(key []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM cipher: %w", err)
	}

	return gcm, nil
}

func DeriveKey(config KeyDerivationConfig) ([]byte, error) {
	if len(config.Passphrase) == 0 || len(config.Salt) == 0 {
		return nil, errors.New("passphrase and salt are required")
	}

	return argon2.IDKey(config.Passphrase, config.Salt, 1, 64*1024, 4, 32), nil
}
