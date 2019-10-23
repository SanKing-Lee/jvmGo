package classfile

type MarkerAttribute struct {
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// 假装读取了数据
}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}
