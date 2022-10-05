package classfile

// AttributeInfo 属性信息
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, constantPool ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, constantPool)
	}
	return attributes
}

func readAttribute(reader *ClassReader, constantPool ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := constantPool.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attributeInfo := newAttributeInfo(attrName, attrLen, constantPool)
	attributeInfo.readInfo(reader)
	return attributeInfo
}

func newAttributeInfo(attrName string, attrLen uint32, constantPool ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{constantPool: constantPool}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{constantPool: constantPool}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
