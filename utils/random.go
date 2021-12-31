package utils

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
)

//the default way of relying on unix nanosecond time to get a new random seed didnt work for me and produced the same result over and over again as the function ran too quickly.
//this is why we have to rely on the crypto/rand package instead to seed
func GetNewRandomSeed() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

//gets a random number between 1-i
func GetRandomNumber(i int) int {
	GetNewRandomSeed()
	return math_rand.Intn(i)
}

//gets a random number between min-max
func GetRandomNumberInRange(min int, max int) int {
	GetNewRandomSeed()
	return math_rand.Intn(max-min) + min
}
