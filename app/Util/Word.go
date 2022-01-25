package Util

import (
	"strings"
	"unicode"
)

//ToUpper 转大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

//ToLower 转小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

//UnderscoreToUpperCamelCase 下划线单词转大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

//UnderscoreToLowerCamelCase 下划线单词转小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
