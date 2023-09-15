package domain

import (
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
)

type Post struct {
	PostId         int64
	UserId         int64
	Topic          string
	Description    Lexical
	Period         string
	HowTo          Lexical
	OpenPositions  []string
	RequiredSkills []string
	Benefits       []string
}

func CheckRequireFields(post *pbv1.Post) bool {
	if post.Topic == "" || post.Description == "" || post.Period == "" || post.HowTo == "" {
		return false
	}
	if len(post.OpenPositions) == 0 || len(post.RequiredSkills) == 0 || len(post.Benefits) == 0 {
		return false
	}
	return true
}

func NewPost(post *pbv1.Post) *Post {
	return &Post{
		Topic:          post.Topic,
		Description:    CreateLexical(post.Description),
		Period:         post.Period,
		HowTo:          CreateLexical(post.HowTo),
		OpenPositions:  post.OpenPositions,
		RequiredSkills: post.RequiredSkills,
		Benefits:       post.Benefits,
	}
}
