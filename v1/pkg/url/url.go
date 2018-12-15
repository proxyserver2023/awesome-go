package url

type Error struct {
	Op, URL string
	Err     error
}

func (e *Error) Error() string { return e.Op + " " + e.URL + ": " + e.Err.Error() }

type timeout interface {
	Timeout() bool
}

func (e *Error) Timeout() bool {
	t, ok := e.Err.(timeout)
	return ok && t.Timeout()
}
