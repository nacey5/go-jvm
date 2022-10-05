package classfile

// ConstantMethodHandleInfo 方法控制器
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

// ConstantMethodTypeInfo 方法类型信息
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

// ConstantInvokeDynamicInfo 支持动态预压的引导方法属性类
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	self.referenceKind = reader.readUint8()
	self.referenceIndex = reader.readUint16()
}

func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.descriptorIndex = reader.readUint16()
}

func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootstrapMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}