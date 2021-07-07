package utils

import (
	"math/rand"
	"github.com/satori/go.uuid"
	"time"
)

// GetRandomString from len return random string
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//生成唯一uuid
func GenUuid() string {
	u2 := uuid.NewV4()
	return u2.String()

}
