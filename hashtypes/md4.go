package hashtypes

import (
	"encoding/hex"

	"golang.org/x/crypto/md4"
)

func MD4Hash(text string) string {
	hash := md4.New()
	hash.Write([]byte(text))
	md4_hash := hex.EncodeToString(hash.Sum(nil))

	return md4_hash
}
