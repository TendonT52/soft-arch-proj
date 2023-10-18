package domain

import (
	pbv1 "github.com/TikhampornSky/report-service/gen/v1"
)

const (
	REPORT_TYPE_SCAM_LIST       = "Scam And Fraudulent Listing"
	REPORT_TYPE_FAKE_REVIEW     = "Fake Review"
	REPORT_TYPE_SUSPICIOUS_USER = "Suspicious User"
	REPORT_TYPE_WEBSITE_BUGS    = "Website Bug"
	REPORT_TYPE_SUGGESTION      = "Suggestion"
	REPORT_TYPE_OTHER           = "Other"
)

func CheckRequireFields(report *pbv1.CreatedReport) bool {
	if report.Topic == "" || report.Type == "" || report.Description == "" {
		return false
	}
	return true
}
