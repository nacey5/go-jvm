package classfile

/*
	Signature_attribute {
	    u2 attribute_name_index;
	    u4 attribute_length;
	    u2 signature_index;
	}
*/
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
