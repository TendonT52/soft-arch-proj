package domain

import "fmt"

type ErrKind int

const (
	_ ErrKind = iota
	PasswordNotMatch
	UserEmailNotFound
	UserIDNotFound
	DuplicateEmail
	NotChulaStudentEmail
	AlreadyVerified
	NotVerified
	Forbidden
	Unauthorized
	RedisNotFound
	MailNotSent
	InvalidStatus
	YearMustBeGreaterThanZero

	Internal
)

var (
	ErrPasswordNotMatch          = BackEndError{Kind: PasswordNotMatch}
	ErrUserEmailNotFound         = BackEndError{Kind: UserEmailNotFound}
	ErrUserIDNotFound            = BackEndError{Kind: UserIDNotFound}
	ErrDuplicateEmail            = BackEndError{Kind: DuplicateEmail}
	ErrNotChulaStudentEmail      = BackEndError{Kind: NotChulaStudentEmail}
	ErrAlreadyVerified           = BackEndError{Kind: AlreadyVerified}
	ErrNotVerified               = BackEndError{Kind: NotVerified}
	ErrForbidden                 = BackEndError{Kind: Forbidden}
	ErrUnauthorized              = BackEndError{Kind: Unauthorized}
	ErrRedisNotFound             = BackEndError{Kind: RedisNotFound}
	ErrMailNotSent               = BackEndError{Kind: MailNotSent}
	ErrInvalidStatus             = BackEndError{Kind: InvalidStatus}
	ErrYearMustBeGreaterThanZero = BackEndError{Kind: YearMustBeGreaterThanZero}

	ErrInternal = BackEndError{Kind: Internal}
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
