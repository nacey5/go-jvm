package classfile

import "fmt"

type AttributeInfo struct {
}

type ClassFile struct {
	//魔数
	magic uint32
	//次要版本
	minorVersion uint16
	//主要版本
	majorVersion uint16
	//常量池
	constantPool ConstantPool
	//类访问标志，设定类属于什么类型，共有还是私有
	accessFlags uint16
	//类索引
	thisClass uint16
	//超类索引
	superClass uint16
	//接口索引表
	interfaces []uint16
	fields     []*MemberInfo
	methods    []*MemberInfo
	attributes []AttributeInfo
}

// Parse 此函数将[]byte解析成ClassFile结构体
func Parse(classData []byte) (classFile *ClassFile, err error) {
	defer func() {
		//此函数可以恢复因为执行所导致的go宕机状态,正常情况下返回的是nil
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				fmt.Errorf("%v", r)
			}
		}
	}()

	classReader := &ClassReader{classData}
	classFile = &ClassFile{}
	classFile.read(classReader)
	return
}

// go语言没有处理机制，只有panic-recover来对class进行解析
func (this *ClassFile) read(reader *ClassReader) {
	this.readAndCheckMagic(reader)
	this.readAndCheckVersion(reader)
	this.constantPool = reader.readConstantPool(reader)
	this.accessFlags = reader.readUint16()
	this.thisClass = reader.readUint16()
	this.interfaces = reader.readUint16s()
	this.fields = readMembers(reader, this.constantPool)
	this.methods = readMembers(reader, this.constantPool)
	this.attributes = readAtttibutes(reader, this.constantPool)
}

// 检查魔数
func (this *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFABABE {
		panic("java.lang.ClassFormatError:magic!")
	}
}

func (this *ClassFile) readAndCheckVersion(reader *ClassReader) {
	this.minorVersion = reader.readUint16()
	this.majorVersion = reader.readUint16()
	//检查主版本
	switch this.majorVersion {
	case 45:
		return
		//检查次版本
	case 46, 47, 48, 49, 50, 51, 52:
		if this.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")

}

func (this *ClassFile) MinorVersion() uint16 {
	return this.minorVersion
}

func (this *ClassFile) MajorVersion() uint16 {
	return this.majorVersion
}

func (this *ClassFile) ConstantPool() ConstantPool {
	return this.constantPool
}

func (this *ClassFile) AccessFlags() uint16 {
	return this.accessFlags
}

func (this *ClassFile) Fields() []*MemberInfo {
	return this.fields
}

func (this *ClassFile) Methods() []*MemberInfo {
	return this.methods
}

func (this *ClassFile) ClassName() string {
	return this.constantPool.getClassName(this.thisClass)
}
func (this *ClassFile) SuperClassName() string {
	if this.superClass > 0 {
		return this.constantPool.getClassName(this.superClass)
	}
	//代表只拥有Object,而没有超类
	return ""
}
func (this *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(this.interfaces))
	for i, cpIndex := range this.interfaces {
		interfaceNames[i] = this.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
