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

func getJreDir(jreOption string) string {
	if jreOption != "" && exist
}
