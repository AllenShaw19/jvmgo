package classpath

import "os"

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) 
}