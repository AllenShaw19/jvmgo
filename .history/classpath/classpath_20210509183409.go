package classpath

import "path/filepath"

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
}
