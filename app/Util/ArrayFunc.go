package Util

//InArray 判断指定字符是否在切片中
func InArray(target string, strArray []string) bool {
	for _, element := range strArray {
		if target == element {
			return true
		}
	}
	return false
}
