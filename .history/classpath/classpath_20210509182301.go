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
	jreExtPath := filepath.Join(jreDir, "")

	self.bootstrapClasspath = newWildcardEntry(jreLibPath)

	

}

func (c *Classpath) parseAppClasspath(cpOption string) {}
