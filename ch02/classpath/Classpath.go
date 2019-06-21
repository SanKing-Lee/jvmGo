package classpath

import (
	"os"
	"path/filepath"
)

/*
 *@description: 类路径，主要由启动类、扩展类和用户自定义类组成
 *@version:
  0.0.1 初始化
*/
type Classpath struct {
	bootClasspath Entry // 启动类
	extClasspath  Entry // 扩展类
	userClasspath Entry // 用户自定义类
}

/*
 *@description: 解析类路径
 *@params: jreOption: 传入的jre路径, cpOption: 传入的用户类路径
 *@return: 指向一个新的Classpath的指针
 *@version:
  0.0.1 初始化
*/
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

/*
 *@description: 根据传入的jre路径寻找启动类和扩展类
 *@params: jre路径
 *@return:
 *@version:
  0.0.1 初始化
*/
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// 在jre/lib中去寻找启动类
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// 在jre/lib/ext中寻找扩展类
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

/*
 *@description:根据传入的jre路径寻找jre目录
 *@params:jreOption:传入的jre路径
 *@return:寻找到的jre目录路径
 *@version:
  0.0.1 初始化
*/
func getJreDir(jreOption string) string {
	// 如果指定jre路径不为空且该路径对应文件是存在的，则直接返回该路径
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// 如果当前目录下存在jre，则返回./jre
	if exists("./jre") {
		return "./jre"
	}
	// 在JAVA_HOME这个环境变量的路径下去寻找jre
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	// 在环境变量下也没有找到，报错
	panic("Can not find jre folder!")
}

/*
 *@description: 判断一个路径对应的文件是否存在
 *@params: path:需要判断的路径
 *@return: 该文件存在则返回true，否则返回false
 *@version:
  0.0.1 初始化
*/
func exists(path string) bool {
	// 读取路径对应文件的状态，返回一个FileInfo和error，如果该文件存在问题，则error为*PathErr
	if _, err := os.Stat(path); err != nil {
		// 文件不存在
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/*
 *@description: 解析用户类路径
 *@params: cpOption: 用户指定的用户类路径
 *@return:
 *@version:
  0.0.1 初始化
*/
func (self *Classpath) parseUserClasspath(cpOption string) {
	// 如果忽略了cpOption，则默认为当前的目录
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

/*
 *@description: 加载指定类
 *@params: className: 指定类的名称
 *@return: 从指定类读取到的内容，读取到该类的入口，错误信息
 *@version:
  0.0.1 初始化
*/
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	// 传入的className不包含.class后缀，加上
	className = className + ".class"
	// 从启动类中找到了该类，返回数据
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	// 从扩展类中找到了该类
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	// 在用户类中寻找该类
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
