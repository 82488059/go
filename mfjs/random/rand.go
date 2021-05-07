package random

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))
var str = "0123456789abcdefghijklmnopqrstuvwxyz"
var az09 = []byte(str)
var laz09 = len(az09)

// GetRandomStringaz09 random string a-z 0-9
func GetRandomStringaz09(length int) string {
	result := []byte{}
	for i := 0; i < length; i++ {
		result = append(result, az09[r.Intn(laz09)])
	}
	return string(result)
}

// Int random int
func Int(min int, max int) int {
	return min + r.Intn(max-min)
}
