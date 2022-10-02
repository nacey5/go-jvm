package classfile

// 边长属性，存放字节码中方法的相关信息
type CodeAttribute struct {
	constantPool ConstantPool
	//操作栈的最大深度
	maxStack uint16
	//局部变量表的大小
	maxLocals uint16
	//字节码
	code []byte
	//异常表
	exceptionTable []*ExceptionTableEntry
	//属性表
	attributes []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (this *CodeAttribute) readInfo(reader *ClassReader) {
	this.maxStack = reader.readUint16()
	this.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	//读取代码长度
	this.code = reader.readBytes(codeLength)
	this.exceptionTable = readExceptionTable(reader)
	this.attributes = readAttributes(reader, this.constantPool)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	//读取异常表长度
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			//pc位置开始
			startPc: reader.readUint16(),
			//pc位置结束
			endPc: reader.readUint16(),
			//handler_pc从当前异常处理器用于处理异常(catch块)的第一条指令相对于字节码开始处的偏移量
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
