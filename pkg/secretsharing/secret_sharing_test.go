package secretsharing

import (
	"bytes"
	"testing"
)

func TestSplitKey(t *testing.T) {
	secret := []byte("my-secret-key")
	n := 5
	k := 3

	shares, err := SplitKey(secret, n, k)
	if err != nil {
		t.Fatalf("Error splitting key: %v", err)
	}

	if len(shares) != n {
		t.Fatalf("Expected %d shares, got %d", n, len(shares))
	}

	for i := 0; i < len(shares); i++ {
		for j := i + 1; j < len(shares); j++ {
			if bytes.Equal(shares[i], shares[j]) {
				t.Fatalf("Shares %d and %d are the same; they should be unique", i, j)
			}
		}
	}
}

func TestCombineKey(t *testing.T) {
	secret := []byte("my-secret-key")
	n := 5
	k := 3

	shares, err := SplitKey(secret, n, k)

	reconstructedSecret, err := CombineKey(shares[:k])
	if err != nil {
		t.Fatalf("Error combining key: %v", err)
	}

	if !bytes.Equal(secret, reconstructedSecret) {
		t.Errorf("Reconstructed secret does not match original: got %s, want %s", reconstructedSecret, secret)
	}
}

func TestInvalidCombineKey(t *testing.T) {
	secret := []byte("my-secret-key")
	n := 5
	k := 3

	shares, err := SplitKey(secret, n, k)
	if err != nil {
		t.Fatalf("Error splitting key: %v", err)
	}

	invalidShare := append(shares[0], 0xFF)
	selectedShares := append([][]byte{invalidShare}, shares[1:3]...)

	_, err = CombineKey(selectedShares)
	if err == nil {
		t.Error("Expected error when combining invalid shares")
	}
}
