package commands

var (
	// CommandMap maps strings to associated commands
	CommandMap = make(map[string]Command)
)

func init() {
	CommandMap["ping"] = NewCommand(pingPreRun, pingRun, pingPostRun, pingHelp)
}

// NewCommand returns new Command-objects
func NewCommand(preRun CommandFunc, run CommandFunc, postRun CommandFunc, help CommandFunc) Command {
	return Command{PreRun: preRun, Run: run, PostRun: postRun, Help: help}
}
