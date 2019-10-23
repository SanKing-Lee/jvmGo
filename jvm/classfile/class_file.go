package classfile

import "fmt"

type ClassFile struct {
	magic        uint32          // 魔数
	minorVersion uint16          // 副版本号
	majorVersion uint16          // 主版本号
	constantPool ConstantPool    // 常量池
	accessFlags  uint16          // 访问标志
	thisClass    uint16          // 类索引
	superClass   uint16          // 父类索引
	interfaces   []uint16        // 接口表
	fields       []*MemberInfo   // 字段表
	methods      []*MemberInfo   // 方法表
	attributes   []AttributeInfo // 属性表
}

/*
 *@description: 解析class文件内容，转换为一个ClassFile结构
 *@params: classData: class文件内容
 *@return: cf: 根据ClassData解析得到的ClassFile结构，err: 错误信息
 *@version:
  0.0.1 初始化
*/
func Parse(classData []byte) (cf *ClassFile, err error) {
	// 崩溃恢复机制
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	// 将需要处理的class文件内容放入一个新的ClassReader
	cr := &ClassReader{classData}
	// 创建一个新的ClassFile并调用read函数将class文件解析为一个ClassFile结构
	cf = &ClassFile{}
	cf.read(cr)
	return
}

/*
 *@description: 通过reader从class文件中读取ClassFile所需的各项属性数据
 *@params: reader: classData阅读器
 *@return:
 *@version:
  0.0.1 初始化
*/
func (self *ClassFile) read(reader *ClassReader) {
	// 解析魔数
	self.readAndCheckMagic(reader)
	// 解析版本号
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

/*
 *@description: 读取魔数，并判断该魔是否符合class文件的格式要求即0xCAFEBABE
 *@params: reader: 阅读器
 *@return:
 *@version:
  0.0.1 初始化
*/
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	// 读取首个双字，获取魔数
	self.magic = reader.readUint32()
	// 检查魔数是否符合标准格式
	if self.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/*
 *@description: 检查版本号，支持45.0~52.0的版本
 *@params: reader: 阅读器
 *@return:
 *@version:
  0.0.1 初始化
*/
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	// 读取版本号
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	// 判断版本号，支持45.x 和46.0, 47.0, 48.0, 49.0, 50.0, 51.0, 52.0
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

//*************************************************************getter*********************************************************//
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
