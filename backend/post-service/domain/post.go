package domain

import (
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
)

type SubSearch struct {
	Id    int64
	Title string
}

type IndividualSearchResult struct {
	OpenPosition  map[int64](*[]SubSearch)
	RequiredSkill map[int64](*[]SubSearch)
	Benefits      map[int64](*[]SubSearch)
}

type SummarySearchResult struct {
	Pid 		 int64
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
