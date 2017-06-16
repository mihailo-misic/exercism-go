package erratum

import "io"

type ResourceOpener func() (Resource, error)

type TransientError struct {
	err error
}

type FrobError struct {
	defrobTag string
	inner     error
}

type Resource interface {
	io.Closer
	Frob(string)
	Defrob(string)
}

func (e TransientError) Error() string {
	return e.err.Error()
}

func (e FrobError) Error() string {
	return e.inner.Error()
}
