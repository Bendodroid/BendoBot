package commands

var (
	CommandMap = make(map[string]Command)
)

func init() {
	CommandMap["ping"] = NewCommand(PingPreRun, PingRun, PingPostRun, PingHelp)
}

func NewCommand(preRun CommandFunc, run CommandFunc, postRun CommandFunc, help CommandFunc) Command {
	return Command{PreRun: preRun, Run: run, PostRun: postRun, Help: help}
}
