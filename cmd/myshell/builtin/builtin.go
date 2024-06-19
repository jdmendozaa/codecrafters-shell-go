package builtin

type BuiltinCommands struct {
	CommandsMap map[string]BuiltinCommand
	Path        []string
}

func NewBuiltinCommands(path []string) *BuiltinCommands {
	return &BuiltinCommands{
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

func (builtinCommands *BuiltinCommands) ExecuteBuiltinCommand(command string, args ...string) error {

	c := builtinCommands.CommandsMap[command]

	if c != nil {
		err := c.Execute(args...)
		if err != nil {
			return err
		}
	}
	return nil
}
