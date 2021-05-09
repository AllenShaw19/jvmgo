package classpath

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
	jreLib

}

func (c *Classpath) parseAppClasspath(cpOption string) {}
