package classfile

// 常量池的所有tag项
const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

// 用于表示常量池表中所有项的通用格式，可通过实现该接口实现常量池表的所有项的类型
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

//description: 从文件中读取标签，并创建新的常量项
//params: reader: class文件阅读器，cp: 当前的常量池
//return: 常量项
//version:
//0.0.1 初始化
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

//description: 根据传入的标签生成不同的常量项目
//params: tag: 常量项标签, cp: 当前的常量池
//return: 常量项
//version:
//0.0.1 初始化
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}

	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}

	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}

	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	// case CONSTANT_MethodType:
	// 	return &ConstantMethodTypeInfo{}
	// case CONSTANT_MethodHandle:
	// 	return &ConstantMethodHandleInfo{}
	// case CONSTANT_InvokeDynamic:
	// 	return &ConstantInvokeDynamicINfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
