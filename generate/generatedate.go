package generate

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func numbergenerate() int64 {
	safeNum, err := rand.Int(rand.Reader, big.NewInt(800000))
	if err != nil {
		fmt.Println(err)
	}
	return safeNum.Int64()

}
