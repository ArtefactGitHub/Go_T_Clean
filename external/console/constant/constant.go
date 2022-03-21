package constant

type CommandType int

const (
	ReadAll CommandType = iota
	Read
	Create
	Update
	Delete
	Help
	Exit
	None
)
