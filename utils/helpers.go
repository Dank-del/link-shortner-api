package utils

import (
	"encoding/json"
	"github.com/Dank-del/link-shortner-api/types"
	"io/ioutil"
	"log"
	"unsafe"
)

func RandStringBytesMaskImprSrcUnsafe(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func GetConfig() *types.Config {
	var d *types.Config
	r, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln("Can't read config file")
	}
	err = json.Unmarshal(r, &d)
	if err != nil {
		log.Fatalln(err)
	}
	return d
}
