package classfile

// UnparsedAttribute 属性反解析器
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (this *UnparsedAttribute) readInfo(reader *ClassReader) {
	this.info = reader.readBytes(this.length)
}

func (this *UnparsedAttribute) Info() []byte {
	return this.info
}
