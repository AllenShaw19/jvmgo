package classpath

import "os"

const pathListSeparator = string(os.PathListSeparator)

const 


type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

