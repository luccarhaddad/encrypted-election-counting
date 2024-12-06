package encryption

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncryptionDecryption(t *testing.T) {
	secretKey, publicKey, err := GenerateKeys()
	assert.NoError(t, err)
	assert.NotNil(t, secretKey)
	assert.NotNil(t, publicKey)

	testCases := []int64{1, 2, 3, 4, 5}

	for _, vote := range testCases {
		t.Run(fmt.Sprintf("Vote %d", vote), func(t *testing.T) {
			ciphertext, err := EncryptVote(publicKey, vote)
			assert.NoError(t, err)
			assert.NotNil(t, ciphertext)

			decryptedVote, err := DecryptVote(secretKey, ciphertext)
			assert.NoError(t, err)
			assert.Equal(t, vote, decryptedVote)
		})
	}
}

func TestInvalidInputs(t *testing.T) {
	sk, _, _ := GenerateKeys()

	// Use mismatched key for encryption/decryption
	otherSk, otherPk, _ := GenerateKeys()

	t.Run("MismatchedKeys", func(t *testing.T) {
		ciphertext, err := EncryptVote(otherPk, 42)
		assert.NoError(t, err)

		vote, err := DecryptVote(sk, ciphertext)
		expectedVote, err := DecryptVote(otherSk, ciphertext)
		assert.NotEqual(t, expectedVote, vote)
	})
}
