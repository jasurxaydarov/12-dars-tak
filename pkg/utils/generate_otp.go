package utils

import "math/rand"

func GenerateOtp()int{
return rand.Int()%1000
}