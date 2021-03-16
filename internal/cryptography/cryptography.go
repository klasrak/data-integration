package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"

	"github.com/klasrak/data-integration/internal/converter"
)

// Encrypter ...
type Encrypter interface {
	Encrypt(value string) []byte
	Decrypt([]byte) string
}

type encrypter struct {
	secret string
}

// New ...
func New(secret string) Encrypter {
	return &encrypter{
		secret: secret,
	}
}

func createHash(secret string) []byte {
	hash := sha256.Sum256([]byte(secret))
	return hash[:]
}

func (e *encrypter) Encrypt(value string) []byte {
	block, _ := aes.NewCipher(createHash(e.secret))
	gcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	encryptedText := gcm.Seal(nonce, nonce, converter.StringToBytes(value), nil)

	return encryptedText
}

func (e *encrypter) Decrypt(data []byte) string {
	block, err := aes.NewCipher(createHash(e.secret))

	if err != nil {
		panic(err.Error())
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	nonceSize := gcm.NonceSize()

	nonce, cipherText := data[:nonceSize], data[nonceSize:]

	plainText, err := gcm.Open(nil, nonce, cipherText, nil)

	if err != nil {
		panic(err.Error())
	}

	return converter.BytesToString(plainText)
}
