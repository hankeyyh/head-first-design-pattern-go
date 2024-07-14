package command

type MacroCommand struct {
	commands []Command
}

func NewMacroCommand(commands []Command) *MacroCommand {
	return &MacroCommand{commands: commands}
}

func (m *MacroCommand) Execute() {
	for _, command := range m.commands {
		command.Execute()
	}
}

func (m *MacroCommand) Undo() {
	// do backwards
	for i := len(m.commands) - 1; i >= 0; i-- {
		m.commands[i].Undo()
	}
}