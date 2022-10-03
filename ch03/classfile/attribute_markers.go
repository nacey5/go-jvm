package classfile

// DeprecatedAttribute 标记属性
type DeprecatedAttribute struct {
	MarkerAttribute
}
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

// 不读取任何属性
func (this *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
