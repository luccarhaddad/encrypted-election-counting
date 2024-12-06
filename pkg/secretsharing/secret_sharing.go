package secretsharing

import "github.com/hashicorp/vault/shamir"

func SplitKey(secret []byte, n, k int) ([][]byte, error) {
	return shamir.Split(secret, n, k)
}

func CombineKey(parts [][]byte) ([]byte, error) {
	return shamir.Combine(parts)
}
