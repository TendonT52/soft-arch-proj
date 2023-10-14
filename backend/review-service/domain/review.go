package domain

import (
	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
)

const (
	REPORT_TYPE_SCAM_LIST       = "Scam And Fraudulent Listing"
	REPORT_TYPE_FAKE_REVIEW     = "Fake Review"
	REPORT_TYPE_SUSPICIOUS_USER = "Suspicious User"
	REPORT_TYPE_WEBSITE_BUGS    = "Website Bug"
	REPORT_TYPE_SUGGESTION      = "Suggestion"
	REPORT_TYPE_OTHER           = "Other"
)

func CheckCreatedRequireField(review *pbv1.CreatedReview) bool {
	if review.Cid <= 0 || review.Rating == 0 || review.Title == "" || review.Description == "" {
		return false
	}
	return true
}

func CheckUpdatedRequireField(review *pbv1.UpdatedReview) bool {
	if review.Rating == 0 || review.Title == "" || review.Description == "" {
		return false
	}
	return true
}