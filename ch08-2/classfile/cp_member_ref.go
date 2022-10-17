package classfile

/*
	CONSTANT_Fieldref_info {
	    u1 tag;
	    u2 class_index;
	    u2 name_and_type_index;
	}

	CONSTANT_Methodref_info {
	    u1 tag;
	    u2 class_index;
	    u2 name_and_type_index;
	}

	CONSTANT_InterfaceMethodref_info {
	    u1 tag;
	    u2 class_index;
	    u2 name_and_type_index;
	}
*/
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	this.classIndex = reader.readUint16()
	this.nameAndTypeIndex = reader.readUint16()
}

func (this *ConstantMemberrefInfo) ClassName() string {
	return this.cp.getClassName(this.classIndex)
}
func (this *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return this.cp.getNameAndType(this.nameAndTypeIndex)
}
