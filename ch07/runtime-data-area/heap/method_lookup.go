package heap

// 在类中查找方法，如果没有此方法的时候，继续从其父类继续查找
func LookupMethodInClass(class *Class, name, descriptor string) *Method {
	for c := class; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

// 通过接口寻找接口方法，当此接口方法寻找不到的时候，要继续向上层寻找接口继续遍历接口方法
func lookupMethodInInterfaces(interfaces []*Class, name string, descriptor string) *Method {
	for _, iface := range interfaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := lookupMethodInInterfaces(iface.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}

	return nil
}

// 查找类接口方法
func lookupInterfaceMethod(iface *Class, name string, descriptor string) *Method {
	for _, method := range iface.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return lookupMethodInInterfaces(iface.interfaces, name, descriptor)
}
