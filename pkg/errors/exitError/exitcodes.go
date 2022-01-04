package exitError

type ExitCode int

const (
	Success ExitCode = iota
	Failure
	NotImplemented
	NotFound
)
