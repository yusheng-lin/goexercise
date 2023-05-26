package utility

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

func GetHash(source string) string {
	var hashfunc hash.Hash = sha256.New()
	hashfunc.Write([]byte(source))
	hash := hex.EncodeToString(hashfunc.Sum(nil))
	return hash
}
