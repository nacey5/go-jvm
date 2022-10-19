package classfile

/*
	CONSTANT_String_info {
	    u1 tag;
	    u2 string_index;
	}
*/
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (this *ConstantStringInfo) readInfo(reader *ClassReader) {
	this.stringIndex = reader.readUint16()
}
func (this *ConstantStringInfo) String() string {
	return this.cp.getUtf8(this.stringIndex)
}
