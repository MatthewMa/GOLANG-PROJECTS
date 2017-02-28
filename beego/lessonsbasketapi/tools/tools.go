package tools

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"time"
)

//MD5 algorithm to encode the password
func EncodePassword(pwd string) (result string) {
	h := md5.New()
	h.Write([]byte(pwd))
	return hex.EncodeToString(h.Sum(nil))
}

//check string is not null or empty
func CheckString(s string) (result bool) {
	result = false
	if s != "" {
		result = true
	}
	return result
}

//generate token according to time
func GenerateToken() (token string) {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token = fmt.Sprintf("%x", h.Sum(nil))
	return token
}
