package util

// 这是一个简易的转换成UTF8编码的工具类,简易版，如果不包含null或者补充字符是可以工作的
func DecodeMUTF8(bytes []byte) string {
	return string(bytes)
}
