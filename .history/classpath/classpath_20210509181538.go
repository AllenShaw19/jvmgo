package classpath

type Classpath struct {
	bootstrapClasspath Entry
	extensionsClasspath Entry
	appClasspath Entry
}

func Parse(jreOption, cp)