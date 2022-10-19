package classfile

/*
	attribute_info {
	    u2 attribute_name_index;
	    u4 attribute_length;
	    u1 info[attribute_length];
	}
*/
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
