package classfile

/*
	EnclosingMethod_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	    u2 class_index;
	    u2 method_index;
	}
*/
type EnclosingMethodAttribute struct {
	cp          ConstantPool
	classIndex  uint16
	methodIndex uint16
}

func (this *EnclosingMethodAttribute) readInfo(reader *ClassReader) {
	this.classIndex = reader.readUint16()
	this.methodIndex = reader.readUint16()
}

func (this *EnclosingMethodAttribute) ClassName() string {
	return this.cp.getClassName(this.classIndex)
}

func (this *EnclosingMethodAttribute) MethodNameAndDescriptor() (string, string) {
	if this.methodIndex > 0 {
		return this.cp.getNameAndType(this.methodIndex)
	} else {
		return "", ""
	}
}
