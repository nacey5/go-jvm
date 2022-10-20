package classfile

/*
	CONSTANT_Class_info {
	    u1 tag;
	    u2 name_index;
	}
*/
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (this *ConstantClassInfo) readInfo(reader *ClassReader) {
	this.nameIndex = reader.readUint16()
}
func (this *ConstantClassInfo) Name() string {
	return this.cp.getUtf8(this.nameIndex)
}
