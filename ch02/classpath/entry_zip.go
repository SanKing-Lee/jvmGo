package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

/*
 *@description: 用于表示zip文件或jar文件层次的class文件入口
 *@version:
  0.0.1 初始化
*/
type ZipEntry struct {
	absPath string
}

/*
 *@description: 构造一个新的zip入口
 *@params: path: zip文件或jar文件的路径
 *@return: 	根据该zip文件或jar文件路径构造的zip入口
 *@version:
  0.0.1 初始化
*/
func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

/*
 *@description:	读取zip文件并找到对应的class文件，如果没有找到则返回错误
 *@params: className: 需要寻找的class文件名
 *@return: []byte: 读取到的class文件的内容，Entry: 当前的ZipEntry，error: 错误信息
 *@version:
  0.0.1 初始化
*/
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 从zip文件中读取所有的class文件
	r, err := zip.OpenReader(self.absPath)
	// 从zip文件读取出错
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	// 遍历所有文件
	for _, f := range r.File {
		// 找到了需要的class文件
		if f.Name == className {
			// 打开这个对应的class文件
			rc, err := f.Open()
			// 打开文件出错
			if err != nil {
				return nil, nil, err
			}
			// 程序执行完毕后关闭读取流
			defer rc.Close()
			// 读取class文件的内容
			data, err := ioutil.ReadAll(rc)
			// 从class文件中读取出错
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	// 如果zip文件中没有想要找到的class文件
	return nil, nil, errors.New("class not found: " + className)
}

/*
 *@description: 描述该zipEntry的字符串
 *@params:
 *@return: string: 该zipEntry的绝对路径
 *@version:
  0.0.1 初始化
*/
func (self *ZipEntry) String() string {
	return self.absPath
}
