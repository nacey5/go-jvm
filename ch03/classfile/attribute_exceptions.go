package classfile

// ExceptionsAttribute 异常属性
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (this *ExceptionsAttribute) readInfo(reader *ClassReader) {
	this.exceptionIndexTable = reader.readUint16s()
}

func (this *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return this.exceptionIndexTable
}
