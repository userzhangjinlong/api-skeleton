package Util

import "math/rand"

var seeks = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//RandSeeks 返回指定长度的随机种子
func RandSeeks(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = seeks[rand.Intn(len(seeks))]
	}
	return string(b)
}
