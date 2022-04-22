package exitError

type ExitError struct {
	s string
	e ExitCode
}

func (e *ExitError) Error() string {
	return e.s
}

func New(text string, exitCode ExitCode) error {
	return &ExitError{
		s: text,
		e: exitCode,
	}
}
