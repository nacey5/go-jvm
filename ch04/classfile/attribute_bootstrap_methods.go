package classfile

// 引导类加载器
type BootstrapMethodsAttribute struct {
	bootstrapMethods []*BootstrapMethod
}

func (this *BootstrapMethodsAttribute) readInfo(reader *ClassReader) {
	numBootstrapMethods := reader.readUint16()
	this.bootstrapMethods = make([]*BootstrapMethod, numBootstrapMethods)
	for i := range this.bootstrapMethods {
		this.bootstrapMethods[i] = &BootstrapMethod{
			bootstrapMethodRef: reader.readUint16(),
			bootstrapArguments: reader.readUint16s(),
		}
	}
}

// 引导加载方法
type BootstrapMethod struct {
	bootstrapMethodRef uint16
	bootstrapArguments []uint16
}
