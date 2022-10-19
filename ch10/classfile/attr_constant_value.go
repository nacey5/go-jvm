package classfile

/*
	ConstantValue_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	    u2 constantvalue_index;
	}
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (this *ConstantValueAttribute) readInfo(reader *ClassReader) {
	this.constantValueIndex = reader.readUint16()
}

func (this *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return this.constantValueIndex
}
