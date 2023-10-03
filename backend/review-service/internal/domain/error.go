package domain

type ErrKind int

const (
	_ ErrKind = iota
	ReviewIDNotFound
	BadRequest
	Internal
)

var (
	ErrInternal         = &BackendError{Kind: Internal, Message: "Internal server error"}
	ErrReviewIDNotFound = &BackendError{Kind: ReviewIDNotFound, Message: "Review ID not found"}
	ErrBadRequest       = &BackendError{Kind: BadRequest, Message: "Bad request"}
)

type BackendError struct {
	Kind    ErrKind
	Message string
	Detail  map[string]string
	Err     error
}

func (e *BackendError) Error() string {
	return e.Message
}

func (e *BackendError) Unwrap() error {
	return e.Err
}
