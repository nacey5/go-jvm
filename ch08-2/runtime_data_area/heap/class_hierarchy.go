package heap

// jvms8 6.5.instanceof
// jvms8 6.5.checkcast
func (this *Class) isAssignableFrom(other *Class) bool {
	s, t := other, this

	if s == t {
		return true
	}

	if !s.IsArray() {
		if !s.IsInterface() {
			// s is class
			if !t.IsInterface() {
				// t is not interface
				return s.IsSubClassOf(t)
			} else {
				// t is interface
				return s.IsImplements(t)
			}
		} else {
			// s is interface
			if !t.IsInterface() {
				// t is not interface
				return t.isJlObject()
			} else {
				// t is interface
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		// s is array
		if !t.IsArray() {
			if !t.IsInterface() {
				// t is class
				return t.isJlObject()
			} else {
				// t is interface
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			// t is array
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}

	return false
}

// this extends c
func (this *Class) IsSubClassOf(other *Class) bool {
	for c := this.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

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

// this extends iface
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

// iface extends this
func (this *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(this)
}
