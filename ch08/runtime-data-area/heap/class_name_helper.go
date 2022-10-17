package heap

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

func toDescriptor(className string) string {
	if className[0] == '[' {
		return className
	}
	if d, ok := primitiveTypes[className]; ok {
		return d
	}
	return "L" + className + ";"
}

func toClassName(descriptor string) string {
	if descriptor[0] == '[' { //array
		return descriptor
	}
	if descriptor[0] == 'L' {
		return descriptor[1 : len(descriptor)-1]
	}
	for className, d := range primitiveTypes {
		if d == descriptor { //primitive
			return className
		}
	}
	panic("Invalid descriptor:" + descriptor)
}
