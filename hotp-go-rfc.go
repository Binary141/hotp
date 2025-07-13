package hotp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math"
)

// converts the string value into the index that is used for parsing the 4 bytes
func stToNum(char byte) int32 {
	bits := int32(char) & 0xf

	return bits
}

func dt(secret string, counter uint64) int32 {
	// HS := sha1(K,C)
	hasher := hmac.New(sha1.New, []byte(secret))

	// a uint64 is 8 bytes
	bigEndCount := make([]byte, 8)
	binary.BigEndian.PutUint64(bigEndCount, counter)

	_, err := hasher.Write(bigEndCount)
	if err != nil {
		panic(err)
	}

	hash := hasher.Sum(nil)

	offsetBits := hash[0 : 19+1]

	offset := stToNum(offsetBits[19])
	if offset < 0 || offset > 15 {
		panic(fmt.Sprintf("offset has to be >= 0 and <= 15. Got: %d", offset))
	}

	P := hash[offset : offset+3+1]

	a := int32(P[0] & 0x7f)
	b := int32(P[1] & 0xff)
	c := int32(P[2] & 0xff)
	d := int32(P[3] & 0xff)

	return int32(a<<24 | b<<16 | c<<8 | d)
}

func Hotp(secret string, counter uint64, digits int) int {
	Sbits := dt(secret, counter)

	return int(Sbits % int32(math.Pow10(digits)))
}
