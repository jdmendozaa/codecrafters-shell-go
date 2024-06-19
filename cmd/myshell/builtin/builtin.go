package builtin

type Commands struct {
	CommandsMap map[string]BuiltinCommand
	Path        []string
}

func NewBuiltinCommands(path []string) *Commands {
	return &Commands{
		CommandsMap: map[string]BuiltinCommand{
			"exit": &ExitCommand{},
			"echo": &EchoCommand{},
			"type": nil,
		},
		Path: path,
	}
}

type BuiltinCommand interface {
	Execute(args ...string) error
}

func (builtinCommands *Commands) ExecuteBuiltinCommand(command string, args ...string) error {

	c := builtinCommands.CommandsMap[command]

	if c != nil {
		err := c.Execute(args...)
		if err != nil {
			return err
		}
	}
	return nil
}
