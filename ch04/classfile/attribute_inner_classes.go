package classfile

// InnerClassesAttribute 内部类
type InnerClassesAttribute struct {
	classes []*InnerClassInfo
}

type InnerClassInfo struct {
	//内部类索引
	innerClassInfoIndex uint16
	//外部引用索引
	outerClassInfoIndex uint16
	//内部类名字索引
	innerNameIndex uint16
	//内部类访问标志
	innerClassAccessFlags uint16
}

func (self *InnerClassesAttribute) readInfo(reader *ClassReader) {
	numberOfClasses := reader.readUint16()
	self.classes = make([]*InnerClassInfo, numberOfClasses)
	for i := range self.classes {
		self.classes[i] = &InnerClassInfo{
			innerClassInfoIndex:   reader.readUint16(),
			outerClassInfoIndex:   reader.readUint16(),
			innerNameIndex:        reader.readUint16(),
			innerClassAccessFlags: reader.readUint16(),
		}
	}
}
