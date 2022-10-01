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

}
