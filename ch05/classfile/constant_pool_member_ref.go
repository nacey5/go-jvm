package classfile

type ConstantMemberRefInfo struct {
	constantPool     ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	this.classIndex = reader.readUint16()
	this.nameAndTypeIndex = reader.readUint16()
}

func (this *ConstantMemberRefInfo) ClassName() string {
	return this.constantPool.getClassName(this.classIndex)
}
func (this *ConstantMemberRefInfo) NameAndDescriptor() (string, string) {
	return this.constantPool.getNameAndType(this.nameAndTypeIndex)
}

type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}

//type ConstantFieldInfo struct {
//	ConstantMemberRefInfo
//}
