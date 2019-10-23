package classfile

type ConstantPool []ConstantInfo

//description: 从class文件中读取常量池内容
//params: reader；阅读器
//return: ConstantPool：从Class文件中读取的常量池
//version:
//0.0.1 初始化
func readConstantPool(reader *ClassReader) ConstantPool {
	// 读取常量计数器
	cpCount := int(reader.readUint16())
	// 初始化常量池
	cp := make([]ConstantInfo, cpCount)
	// 从1开始遍历填充常量池
	for i := 1; i < cpCount; i++ {
		// 调用函数读取一个ConstantInfo
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		// 如果这个ConstantInfo是long和double，则需要占两个位
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	// 成功从常量池读取到索引对应的常量信息
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	// 从常量表获取名称类型的信息结构体
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	// 获得名称
	name := self.getUtf8(ntInfo.nameIndex)
	// 获得类型
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	name := self.getUtf8(classInfo.nameIndex)
	return name
}

func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
