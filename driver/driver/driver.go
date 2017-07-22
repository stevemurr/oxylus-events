package driver

// Driver interface implements a runnable action
type Driver interface {
	// Description --
	Description() string
	// Run represents the function that will be executed for this action
	Run(action string) error
	// Commands returns a list of objects for each runnable command
	Commands() []map[string]string
	// Get returns a value from a sensor
	Get(action string) (interface{}, error)
	// Raw command allows you to send arbitrary args to the driver - for debugging
	Raw(args ...string) error
}
