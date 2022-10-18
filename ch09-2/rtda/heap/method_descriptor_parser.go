package heap

import "strings"

type MethodDescriptorParser struct {
	raw    string
	offset int
	parsed *MethodDescriptor
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{}
	return parser.parse(descriptor)
}

func (this *MethodDescriptorParser) parse(descriptor string) *MethodDescriptor {
	this.raw = descriptor
	this.parsed = &MethodDescriptor{}
	this.startParams()
	this.parseParamTypes()
	this.endParams()
	this.parseReturnType()
	this.finish()
	return this.parsed
}

func (this *MethodDescriptorParser) startParams() {
	if this.readUint8() != '(' {
		this.causePanic()
	}
}
func (this *MethodDescriptorParser) endParams() {
	if this.readUint8() != ')' {
		this.causePanic()
	}
}
func (this *MethodDescriptorParser) finish() {
	if this.offset != len(this.raw) {
		this.causePanic()
	}
}

func (this *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + this.raw)
}

func (this *MethodDescriptorParser) readUint8() uint8 {
	b := this.raw[this.offset]
	this.offset++
	return b
}
func (this *MethodDescriptorParser) unreadUint8() {
	this.offset--
}

func (this *MethodDescriptorParser) parseParamTypes() {
	for {
		t := this.parseFieldType()
		if t != "" {
			this.parsed.addParameterType(t)
		} else {
			break
		}
	}
}

func (this *MethodDescriptorParser) parseReturnType() {
	if this.readUint8() == 'V' {
		this.parsed.returnType = "V"
		return
	}

	this.unreadUint8()
	t := this.parseFieldType()
	if t != "" {
		this.parsed.returnType = t
		return
	}

	this.causePanic()
}

func (this *MethodDescriptorParser) parseFieldType() string {
	switch this.readUint8() {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return this.parseObjectType()
	case '[':
		return this.parseArrayType()
	default:
		this.unreadUint8()
		return ""
	}
}

func (this *MethodDescriptorParser) parseObjectType() string {
	unread := this.raw[this.offset:]
	semicolonIndex := strings.IndexRune(unread, ';')
	if semicolonIndex == -1 {
		this.causePanic()
		return ""
	} else {
		objStart := this.offset - 1
		objEnd := this.offset + semicolonIndex + 1
		this.offset = objEnd
		descriptor := this.raw[objStart:objEnd]
		return descriptor
	}
}

func (this *MethodDescriptorParser) parseArrayType() string {
	arrStart := this.offset - 1
	this.parseFieldType()
	arrEnd := this.offset
	descriptor := this.raw[arrStart:arrEnd]
	return descriptor
}
