package generate

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func numbergenerate() {
	r, err := rand.Int(rand.Reader, big.NewInt(80))
	fmt.Println(r, err)

}
