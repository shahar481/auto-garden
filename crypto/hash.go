package crypto

import "crypto/sha256"

func HashPassword(password string, salt string) [32]byte {
	return sha256.Sum256([]byte(password + salt))
}
