package classfile

// 可选定长属性
type SourceFileAttribute struct {
	constantPool    ConstantPool
	sourceFileIndex uint16
}

func (this *SourceFileAttribute) readInfo(reader *ClassReader) {
	this.sourceFileIndex = reader.readUint16()
}

func (this *SourceFileAttribute) FileName() string {
	return this.constantPool.getUtf8(this.sourceFileIndex)
}
