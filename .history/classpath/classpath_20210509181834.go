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

}

func (c *Classpath) parseAppClasspath(cp)