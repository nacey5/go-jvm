package classfile

/*
	Code_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	    u2 max_stack;
	    u2 max_locals;
	    u4 code_length;
	    u1 code[code_length];
	    u2 exception_table_length;
	    {   u2 start_pc;
	        u2 end_pc;
	        u2 handler_pc;
	        u2 catch_type;
	    } exception_table[exception_table_length];
	    u2 attributes_count;
	    attribute_info attributes[attributes_count];
	}
*/
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

func (this *CodeAttribute) readInfo(reader *ClassReader) {
	this.maxStack = reader.readUint16()
	this.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	this.code = reader.readBytes(codeLength)
	this.exceptionTable = readExceptionTable(reader)
	this.attributes = readAttributes(reader, this.cp)
}

func (this *CodeAttribute) MaxStack() uint {
	return uint(this.maxStack)
}
func (this *CodeAttribute) MaxLocals() uint {
	return uint(this.maxLocals)
}
func (this *CodeAttribute) Code() []byte {
	return this.code
}
func (this *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return this.exceptionTable
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
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

func (this *CodeAttribute) LineNumberTableAttribute() *LineNumberTableAttribute {
	for _, attrInfo := range this.attributes {
		switch attrInfo.(type) {
		case *LineNumberTableAttribute:
			return attrInfo.(*LineNumberTableAttribute)
		}
	}
	return nil
}
