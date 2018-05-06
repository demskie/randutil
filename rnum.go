package randutil

import (
	"bytes"
	"encoding/binary"
	"math/rand"
	"time"

	cryptorand "crypto/rand"
)

func CreateUniqueMathRnum() *rand.Rand {
	var randomness int64
	b := make([]byte, 8)
	cryptorand.Read(b)
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.LittleEndian, &randomness)
	return rand.New(rand.NewSource(randomness))
}

func CreateMultipleUniqueMathRnum(num int) []*rand.Rand {
	result := make([]*rand.Rand, num)
	for i := 0; i < num; i++ {
		result[i] = CreateUniqueMathRnum()
	}
	return result
}

func CreateBasicMathRnum() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetPosOrNegOne(rnum *rand.Rand) int {
	if rnum.Intn(2) == 0 {
		return 1
	}
	return -1
}
