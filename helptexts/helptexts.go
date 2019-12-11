package helptexts

// DB is a map containing all helptexts
var DB map[string]string

func init() {
	// Initialize empty map, otherwise 'assignment to entry in nil map' errors!
	DB = make(map[string]string)
}
