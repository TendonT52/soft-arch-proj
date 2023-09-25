package domain

import (
	"regexp"

	pbUser "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
)

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

func RemoveSpecialChars(input *pbv1.SearchOptions) *pbv1.SearchOptions {
	// In this pattern, [^a-zA-Z0-9 ] matches any character that is not a letter, digit, or space.
	re := regexp.MustCompile("[^a-zA-Z0-9 ]")
	resultSearchCompany := re.ReplaceAllString(input.SearchCompany, "")
	resultSearchOpenPosition := re.ReplaceAllString(input.SearchOpenPosition, "")
	resultSearchRequiredSkill := re.ReplaceAllString(input.SearchRequiredSkill, "")
	resultSearchBenefit := re.ReplaceAllString(input.SearchBenefit, "")

	result := &pbv1.SearchOptions{
		SearchCompany:        resultSearchCompany,
		SearchOpenPosition:   resultSearchOpenPosition,
		SearchRequiredSkill:  resultSearchRequiredSkill,
		SearchBenefit:        resultSearchBenefit,
	}
	
	return result
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
