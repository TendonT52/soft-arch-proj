package domain

import "fmt"

type ErrKind int

const (
	KindUnknown ErrKind = iota
)

var (
	ErrKindUnknown = BackEndError{Kind: KindUnknown}
)

type BackEndError struct {
	Kind    ErrKind
	Message string
	Detail  map[string]string
	Err     error
}

func (e BackEndError) Error() string {
	return e.Message
}

func (e BackEndError) Is(err error) bool {
	switch errs := err.(type) {
	case BackEndError:
		return e.Kind == errs.Kind
	default:
		return false
	}
}

func (e BackEndError) With(message string, a ...any) *BackEndError {
	ne := e
	ne.Message = fmt.Sprintf(message, a...)
	return &ne
}

func (e BackEndError) WithDetail(message string, detail map[string]string) *BackEndError {
	ne := e
	ne.Message = message
	ne.Detail = detail
	return &ne
}

func (e BackEndError) From(message string, err error) *BackEndError {
	ne := e
	ne.Message = message
	ne.Err = err
	return &ne
}

func (e BackEndError) FromDetail(message string, detail map[string]string, err error) *BackEndError {
	ne := e
	ne.Message = message
	ne.Detail = detail
	ne.Err = err
	return &ne
}

func (e *BackEndError) Unwrap() error {
	return e.Err
}
