//go:generate go run embed.go -dir ../../resources

package resource

var (
	resources = newResourceContainer()
)

// Public Exported package functions

// Exists will check if the requested resource is present
func Exists(file string) bool {
	return resources.Exists(file)
}

// Get will return the requested resource
func Get(file string) ([]byte, bool) {
	return resources.Get(file)
}

// Add will add the resource to the container
func Add(file string, b []byte) {
	resources.Add(file, b)
}

// EOF
