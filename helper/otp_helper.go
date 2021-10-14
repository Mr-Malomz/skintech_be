package helper

import (
	"crypto/rand"
	"math/big"
)

func GenerateOTP() *big.Int{
	otp, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	return otp
}
