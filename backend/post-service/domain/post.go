package domain

type Post struct {
	PostId         int64
	UserId         int64
	Topic          string
	Description    Lexical
	Period         string
	HowTo         Lexical
	OpenPositions  []string
	RequiredSkills []string
	Benefits       []string
}
