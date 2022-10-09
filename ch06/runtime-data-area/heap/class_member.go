package heap

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}
