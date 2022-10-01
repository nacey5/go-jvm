package classfile

// ConstantPool
// 表头给出的常量池大小比实际大1
// 常量池的实际大小为n-1
// 0是无效索引
// Long_info 和Double_info占两个位置
type ConstantPool []ConstantInfo

// 定义ConstantInfo接口来表示常量信息
type ConstantInfo interface {
	readInfo(read *ClassReader)
}

// 读出tag值，再调用newConstantInfo创建具体常量，最后readInfo()方法读取常量信息
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	cInfo := newConstantInfo(tag, cp)
	cInfo.readInfo(reader)
	return cInfo
}

// 根据tag值创建具体的常量
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_CLASS:
		return &ConstantClassInfo{}
	case CONSTANT_DOUBLE:
		return &ConstantDoubleInfo{}
	case CONSTANT_FIELDREF:
		return &ConstantFieldInfo{ConstantMemberRefInfo{constantPool: cp}}
	case CONSTANT_FLOAT:
		return &ConstantFloatInfo{}
	case CONSTANT_INTEGER:
		return &ConstantIntegerInfo{}
	case CONSTANT_INTERFACEMETHODREF:
		return &ConstantInterfaceMethodRefInfo{}
	case CONSTANT_INVOKEDYNAMIC:
		return &ConstantInvokeDynamicInfo{}
	case CONSTANT_LONG:
		return &ConstantLongInfo{}
	case CONSTANT_METHODHANDlE:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_METHODREF:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{constantPool: cp}}
	case CONSTANT_METHODTYPE:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_NAMEANDTYPE:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_STRING:
		return &ConstantStringInfo{}
	case CONSTANT_UTF8:
		return &ConstantUtf8Info{}
	default:
		panic("java.lang.ClassFormatError:constant pool tag")
	}
}

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
