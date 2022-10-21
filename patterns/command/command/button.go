package command

// Button ...
type Button struct {
	Command Command
}

// Press ...
func (b *Button) Press() {
	b.Command.Execute()
}

