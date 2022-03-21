package command

type none struct {
}

func newNoneCommand() Command {
	cmd := none{}
	return &cmd
}

func (cmd *none) Do() (bool, error) {
	return true, nil
}
