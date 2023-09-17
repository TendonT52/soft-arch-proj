package domain

import (
	pbUser "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
)

type IndividualSearchResult struct {
	OpenPositions  map[int64](*[]string)
	RequiredSkills map[int64](*[]string)
	Benefits       map[int64](*[]string)
}

type SummarySearchResult struct {
	Pid           int64
	OpenPosition  *[]string
	RequiredSkill *[]string
	Benefits      *[]string
}

func CheckRequireFields(post *pbv1.Post) bool {
	return checkFields(post.Topic, post.Description, post.Period, post.HowTo, len(post.OpenPositions), len(post.RequiredSkills), len(post.Benefits))
}

func CheckUpdatedFields(post *pbv1.UpdatedPost) bool {
	return checkFields(post.Topic, post.Description, post.Period, post.HowTo, len(post.OpenPositions), len(post.RequiredSkills), len(post.Benefits))
}

func checkFields(topic, description, period, howTo string, lenOpenPositions, lenRequiredSkills, lenBenefits int) bool {
	if topic == "" || description == "" || period == "" || howTo == "" {
		return false
	}
	if lenOpenPositions == 0 || lenRequiredSkills == 0 || lenBenefits == 0 {
		return false
	}
	return true
}

type CompanyInfo struct {
	Ids      *[]int64
	Profiles map[int64](*pbUser.Company)
}

func NewCompanyInfo(data []*pbUser.Company) *CompanyInfo {
	ids := make([]int64, len(data))
	profiles := make(map[int64](*pbUser.Company))
	for i, d := range data {
		ids[i] = d.Id
		profiles[d.Id] = d
	}
	return &CompanyInfo{
		Ids:      &ids,
		Profiles: profiles,
	}
}
