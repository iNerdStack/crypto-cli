package hashtypes

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Hash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	md5_hash := hex.EncodeToString(hash.Sum(nil))

	return md5_hash
}
