package cfg

// Cfg is the root configuration file.
// This global variable is not concurrency safe
var Cfg *Root

// Parse is the global parser for configuration
// This function is not concurrency safe
func Parse(filename string) error {
	root, err := Decode(filename)
	if err != nil {
		return err
	}
	Cfg = root
	return nil
}
