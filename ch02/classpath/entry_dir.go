package classpath

import (
	"io/ioutil"
	"path/filepath"
)

/*
 *@description:
 *@version:
  0.0.1 初始化
*/
type DirEntry struct {
	absDir string // 绝对路径
}

/*
 *@description: 根据路径创建条目结构体
 *@params: path: 路径
 *@return: 一个新的条目指针
 *@version:
  0.0.1 初始化
*/
func newDirEntry(path string) *DirEntry {
	// 获取绝对路径
	absDir, err := filepath.Abs(path)
	// 出错
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

/*
 *@description:	通过路径加载class
 *@params: className: class文件的名称
 *@return: data: class文件的内容, self: 当前的DirEntry, err: 读取class过程中的错误信息
 *@version:
  0.0.1 初始化
*/
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className) // 获取class文件绝对路径
	data, err := ioutil.ReadFile(fileName)            // 从class文件中读取内容
	return data, self, err
}

/*
 *@description: 返回一个描述当前条目的字符串
 *@params:
 *@return: 当前条目的绝对路径字符串
 *@version:
  0.0.1 初始化
*/
func (self *DirEntry) String() string {
	return self.absDir
}
