package command

type exit struct {
}

func newExitCommand() Command {
	cmd := exit{}
	return &cmd
}

func (cmd *exit) Do() (bool, error) {
	return false, nil
}
