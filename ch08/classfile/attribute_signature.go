package classfile

// SignatureAttribute 类、方法表和属性表 用于支持泛型情况下的方法签名

type SignatureAttribute struct {
	cp             ConstantPool
	signatureIndex uint16
}

func (this *SignatureAttribute) readInfo(reader *ClassReader) {
	this.signatureIndex = reader.readUint16()
}

func (this *SignatureAttribute) Signature() string {
	return this.cp.getUtf8(this.signatureIndex)
}
