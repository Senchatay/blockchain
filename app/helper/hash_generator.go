package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type HashGenerator struct {
}

func (*HashGenerator) Get(index int, prevHash string, timestamp time.Time, data string) string {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%d%s%s%s", index, prevHash, timestamp.String(), data)))
	return hex.EncodeToString(hash.Sum(nil))
}
