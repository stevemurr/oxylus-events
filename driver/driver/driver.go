package driver

// Driver interface implements a runnable action
type Driver interface {
	// Run represents the function that will be executed for this action
	Run(action string) error
	// Name describes the name of the driver
	Name() string
	// Description describes the steps the action will take
	Description() string
	// Commands returns a list of objects for each runnable command
	Commands() []map[string]string
	// Raw command allows you to send arbitrary args to the driver
	Raw(args ...string) error
	// Get returns a value from a sensor
	Get(action string, val interface{}) error
}
