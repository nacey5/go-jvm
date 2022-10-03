package classfile

type MemberInfo struct {
	//保存常量池的指针
	constantPool    ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

// 读取字段表或者方法表
func readMembers(reader *ClassReader, constantPool ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, constantPool)
	}
	return members
}

// 读取字段或方法数据
func readMember(reader *ClassReader, constantPool ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    constantPool,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, constantPool),
	}
}

func (this *MemberInfo) AccessFlags() uint16 {
	return this.accessFlags
}

// Name 从常量池查找字段或者方法名
func (this *MemberInfo) Name() string {
	return this.constantPool.getUtf8(this.nameIndex)
}

// Descriptor 从常量池查找字段或者方法描述
func (this *MemberInfo) Descriptor() string {
	return this.constantPool.getUtf8(this.descriptorIndex)
}
