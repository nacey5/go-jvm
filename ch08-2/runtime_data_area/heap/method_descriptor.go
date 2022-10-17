package heap

type MethodDescriptor struct {
	parameterTypes []string
	returnType     string
}

func (this *MethodDescriptor) addParameterType(t string) {
	pLen := len(this.parameterTypes)
	if pLen == cap(this.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, this.parameterTypes)
		this.parameterTypes = s
	}

	this.parameterTypes = append(this.parameterTypes, t)
}
