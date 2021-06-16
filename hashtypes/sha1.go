package hashtypes

import (
	"encoding/hex"

	"crypto/sha1"
)

func SHA1Hash(text string) string {
	hash := sha1.New()
	hash.Write([]byte(text))
	sha1_hash := hex.EncodeToString(hash.Sum(nil))

	return sha1_hash
}
