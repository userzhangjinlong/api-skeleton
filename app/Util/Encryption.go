package Util

import (
	"crypto/md5"
	"encoding/hex"
)

//Md5Encryption md5加密
func Md5Encryption(key string) (md5Val string) {
	h := md5.New()
	h.Write([]byte(key))
	md5Val = hex.EncodeToString(h.Sum(nil))
	return
}
