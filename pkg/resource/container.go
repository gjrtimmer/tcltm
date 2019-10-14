package resource

// Storage container for resources
type container struct {
	storage map[string][]byte
}

func newResourceContainer() *container {
	return &container{
		storage: make(map[string][]byte, 0),
	}
}

// Exists will check if the requested resource exists
func (c *container) Exists(file string) bool {
	if _, ok := c.storage[file]; ok {
		return true
	}

	return false
}

// Get will return the requested resource
func (c *container) Get(file string) ([]byte, bool) {
	if f, ok := c.storage[file]; ok {
		return f, ok
	}

	return nil, false
}

// Add will add the resource to the container
func (c *container) Add(file string, b []byte) {
	c.storage[file] = b
}

// EOF
