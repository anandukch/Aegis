package crypto

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testKey = []byte("12345678901234567890123456789012") // 32 bytes

func TestEncryptDecryptRoundtrip(t *testing.T) {
	plaintext := "john@example.com"
	enc, nonce, err := Encrypt(plaintext, testKey)
	require.NoError(t, err)
	assert.NotEmpty(t, enc)
	assert.NotEmpty(t, nonce)

	dec, err := Decrypt(enc, nonce, testKey)
	require.NoError(t, err)
	assert.Equal(t, plaintext, dec)
}

func TestDecryptTamperDetection(t *testing.T) {
	enc, nonce, err := Encrypt("secret", testKey)
	require.NoError(t, err)

	tampered := enc[:len(enc)-2] + "XX"
	_, err = Decrypt(tampered, nonce, testKey)
	assert.Error(t, err)
}

func TestEncryptProducesUniqueNonce(t *testing.T) {
	enc1, nonce1, _ := Encrypt("same", testKey)
	enc2, nonce2, _ := Encrypt("same", testKey)
	assert.NotEqual(t, nonce1, nonce2)
	assert.NotEqual(t, enc1, enc2)
}

func TestMaskEmail(t *testing.T) {
	result := MaskValue("john@example.com", "email")
	assert.True(t, strings.HasPrefix(result, "j***@"))
	assert.Contains(t, result, "example.com")
}

func TestMaskPhone(t *testing.T) {
	result := MaskValue("9876543210", "phone")
	assert.Equal(t, "******3210", result)
}

func TestMaskCard(t *testing.T) {
	result := MaskValue("4111111111114242", "card_number")
	assert.Equal(t, "****-****-****-4242", result)
}

func TestMaskAadhaar(t *testing.T) {
	result := MaskValue("123456781234", "aadhaar")
	assert.Equal(t, "XXXX-XXXX-1234", result)
}

func TestMaskName(t *testing.T) {
	result := MaskValue("John Doe", "name")
	assert.Equal(t, "J*** D***", result)
}

func TestMaskDOB(t *testing.T) {
	result := MaskValue("1990-07-15", "dob")
	assert.Equal(t, "****-**-15", result)
}
