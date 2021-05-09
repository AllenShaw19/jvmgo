package classpath

type Classpath struct {
	bootstrapClasspath Entry
	extensionsClasspath Entry
	appClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parse
}