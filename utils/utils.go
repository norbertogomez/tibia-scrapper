package utils

import (
	"encoding/json"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func RandomString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func InArray(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func SleepRandom(max int, min int)  {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(max-min) + min
	time.Sleep(time.Duration(r) * time.Second)
}

func ToJson(slice interface{}) ([]byte, error) {
	j, err := json.MarshalIndent(slice, "", "  ")

	return j, err
}

func HandleErrorDefault(err error)  {
	if err != nil {
		log.Errorln(err)
	}
}

func GetFileNameWithTime(prefix string, format string) string {
	return prefix + time.Now().Format("2006-01-02-15-04") + format
}