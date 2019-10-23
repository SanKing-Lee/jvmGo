package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

/*
 *@description: 从字节流中读取一个字节，即java中u1类型数据
 *@params:
 *@return: 字节流首字节
 *@version:
  0.0.1 初始化
*/
func (self *ClassReader) readUint8() uint8 {
	// 读取一个字节
	val := self.data[0]
	// 字节流后移
	self.data = self.data[1:]
	return val
}

/*
 *@description: 从字节流中读取一个字，即java中u2类型数据
 *@params:
 *@return: 字节流首字
 *@version:
  0.0.1 初始化
*/
func (self *ClassReader) readUint16() uint16 {
	// 读取一个字
	val := binary.BigEndian.Uint16(self.data)
	// 字节流后移
	self.data = self.data[2:]
	return val
}

/*
 *@description: 从字节流中读取一个双字，即java类型中的u4类型数据
 *@params:
 *@return: 字节流首双字
 *@version:
  0.0.1 初始化
*/
func (self *ClassReader) readUint32() uint32 {
	// 读取一个双字
	val := binary.BigEndian.Uint32(self.data)
	// 字节流后移
	self.data = self.data[4:]
	return val
}

/*
 *@description: 从字节流中读取一个四字
 *@params:
 *@return: 字节流的首个四字
 *@version:
  0.0.1 初始化
*/
func (self *ClassReader) readUint64() uint64 {
	// 读取一个四字
	val := binary.BigEndian.Uint64(self.data)
	//　字节流后移
	self.data = self.data[8:]
	return val
}

/*
 *@description:	读取uint16表，表的大小由首个uint16数据指出
 *@params:
 *@return:读取到的uint16表
 *@version:
  0.0.1 初始化
*/
func (self *ClassReader) readUint16s() []uint16 {
	// 获取表的长度
	n := self.readUint16()
	// 初始化表
	s := make([]uint16, n)
	// 遍历填充表
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

/*
 *@description: 从字节流中读取指定长度的字节
 *@params:
 *@return: 指定长度的字节切片
 *@version:
  0.0.1 初始化
*/
func (self *ClassReader) readBytes(n uint32) []byte {
	// 读取指定数量的字节
	bytes := self.data[:n]
	// 字节流后移
	self.data = self.data[n:]
	return bytes
}
