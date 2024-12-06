package encryption

import (
	"github.com/tuneinsight/lattigo/v4/bgv"
	"github.com/tuneinsight/lattigo/v4/rlwe"
)

// GenerateKeys creates public and secret keys for encryption
func GenerateKeys() (*rlwe.SecretKey, *rlwe.PublicKey, error) {

	params, err := bgv.NewParametersFromLiteral(bgv.PN12QP109)
	if err != nil {
		return nil, nil, err
	}

	kgen := bgv.NewKeyGenerator(params)
	sk := kgen.GenSecretKey()
	pk := kgen.GenPublicKey(sk)
	return sk, pk, nil
}

func EncryptVote(publicKey *rlwe.PublicKey, vote int64) (*rlwe.Ciphertext, error) {
	params, err := bgv.NewParametersFromLiteral(bgv.PN12QP109)
	if err != nil {
		return nil, err
	}

	encoder := bgv.NewEncoder(params)
	encryptor := bgv.NewEncryptor(params, publicKey)

	plaintext := bgv.NewPlaintext(params, params.MaxLevel())
	encoder.Encode([]uint64{uint64(vote)}, plaintext)

	ciphertext := encryptor.EncryptNew(plaintext)

	return ciphertext, nil
}

func DecryptVote(secretKey *rlwe.SecretKey, ciphertext *rlwe.Ciphertext) (int64, error) {
	params, err := bgv.NewParametersFromLiteral(bgv.PN12QP109)
	if err != nil {
		return 0, err
	}

	if secretKey == nil {
		return 0, err
	}

	encoder := bgv.NewEncoder(params)
	decryptor := bgv.NewDecryptor(params, secretKey)

	plaintext := decryptor.DecryptNew(ciphertext)

	decoded := encoder.DecodeUintNew(plaintext)
	return int64(decoded[0]), nil
}
