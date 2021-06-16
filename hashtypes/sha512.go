package hashtypes

import (
	"crypto/sha512"

	"encoding/hex"
)

func SHA512Hash(text string) string {
	hash := sha512.New()
	hash.Write([]byte(text))
	sha512_hash := hex.EncodeToString(hash.Sum(nil))

	return sha512_hash
}
