package classpath

import (
	"os"
	"path/filepath"

	"github.com/jvmgo/logger"
)

type Classpath struct {
	bootstrapClasspath  Entry
	extensionsClasspath Entry
	appClasspath        Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseAppClasspath(cpOption)
	return cp
}

func (c *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")

	c.bootstrapClasspath = newWildcardEntry(jreLibPath)
	c.extensionsClasspath = newWildcardEntry(jreExtPath)
}

func (c *Classpath) parseAppClasspath(cpOption string) {
	// 如果用户没有提供-classpath或者-cp选型，则使用当前目录作为用户类的classpath
	if cpOption == "" {
		cpOption = "."
	}
	c.appClasspath = newEntry(cpOption)
}

// 获取jre目录
// 优先使用用户输入的-Xjre选项作为jre目录
// 如果用户没有输入-Xjre，则在当前目录下寻找jre目录
// 如果当前目录下没有jre目录，则尝试使用JAVA_HOME环境变量
func getJreDir(jreOption string) string {
	if jreOption != "" && dirExist(jreOption) {
		return jreOption
	}

	if dirExist("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	logger.Panic("Can not find jre folder!")
	return ""
}

func dirExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		logger.Fatal("find path %s, err %v", path, err)
	}

	return true
}

func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"

	if data,entry,err := c.bootstrapClasspath.readClass(class)	
}
