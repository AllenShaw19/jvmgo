package classpath

// WildcardEntry实际上是CompositeEntry
type WildcardEntry CompositeEntry

func newWildcardEntry(path string) WildcardEntry {
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}

	
}
