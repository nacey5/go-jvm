package classfile

// CodeAttribute 边长属性，存放字节码中方法的相关信息
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

func (this *CodeAttribute) ConstantPool() ConstantPool {
	return this.constantPool
}

func (this *CodeAttribute) SetConstantPool(constantPool ConstantPool) {
	this.constantPool = constantPool
}

func (this *CodeAttribute) MaxStack() uint16 {
	return this.maxStack
}

func (this *CodeAttribute) SetMaxStack(maxStack uint16) {
	this.maxStack = maxStack
}

func (this *CodeAttribute) MaxLocals() uint16 {
	return this.maxLocals
}

func (this *CodeAttribute) SetMaxLocals(maxLocals uint16) {
	this.maxLocals = maxLocals
}

func (this *CodeAttribute) Code() []byte {
	return this.code
}

func (this *CodeAttribute) SetCode(code []byte) {
	this.code = code
}

func (this *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return this.exceptionTable
}

func (this *CodeAttribute) SetExceptionTable(exceptionTable []*ExceptionTableEntry) {
	this.exceptionTable = exceptionTable
}

func (this *CodeAttribute) Attributes() []AttributeInfo {
	return this.attributes
}

func (this *CodeAttribute) SetAttributes(attributes []AttributeInfo) {
	this.attributes = attributes
}

// ExceptionTableEntry 异常表实体类
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

func (this *ExceptionTableEntry) StartPc() uint16 {
	return this.startPc
}
func (this *ExceptionTableEntry) EndPc() uint16 {
	return this.endPc
}
func (this *ExceptionTableEntry) HandlerPc() uint16 {
	return this.handlerPc
}
func (this *ExceptionTableEntry) CatchType() uint16 {
	return this.catchType
}
