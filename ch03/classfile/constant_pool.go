package classfile

// ConstantPool
// 表头给出的常量池大小比实际大1
// 常量池的实际大小为n-1
// 0是无效索引
// Long_info 和Double_info占两个位置
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]constantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongIngo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}
func (this ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := this[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}
func (this ConstantPool) getNameAndType(index uint16) (string, string) {
	nameAndType := this.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := this.getUtf8(nameAndType.nameIndex)
	_type := this.getUtf8(nameAndType.descriptorIndex)
	return name, _type
}
func (this ConstantPool) getClassName(index uint16) string {
	classInfo := this.getConstantInfo(index).(*ConstantClassInfo)
	return this.getUtf8(classInfo.nameIndex)
}
func (this ConstantPool) getUtf8(index uint16) string {
	utf8Info := this.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
