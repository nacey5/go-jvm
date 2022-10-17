package heap

// 这是真正的逻辑结构判断
func (this *Class) isAssignableFrom(other *Class) bool {
	s, t := other, this
	if s == t {
		return true
	}
	//如果s不是数组
	if !s.IsArray() {
		//s类不是接口
		if !s.IsInterface() {
			//t不是接口
			if !t.IsInterface() {
				return s.IsSubClassOf(t)
			} else {
				return s.IsImplements(t)
			}
		} else {
			if t.IsInterface() {
				return t.isJlObject()
			} else {
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		if !t.IsArray() {
			if !t.IsInterface() {
				return t.isJlObject()
			} else {
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}

	return false

}

func (this *Class) IsSubClassOf(other *Class) bool {
	for c := this.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// 判断T是否是S的超类
// this implements iface
func (this *Class) IsImplements(iface *Class) bool {
	for c := this; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// 判断S是否实现了T接口，S或S的超类是否实现了某个接口Tx
func (this *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range this.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// c extends this
func (this *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(this)
}
