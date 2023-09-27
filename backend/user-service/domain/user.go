package domain

import "regexp"

const (
	StudentRole = "student"
	AdminRole   = "admin"
	CompanyRole = "company"
)

const (
	ComapanyStatusPending = "Pending"
	ComapanyStatusApprove = "Approve"
	ComapanyStatusReject  = "Reject"
)

type UserStatus struct {
	Verified bool
	Role     string
}

func IsStudentRole(role string) bool {
	return role == StudentRole
}

func IsAdminRole(role string) bool {
	return role == AdminRole
}

func IsCompanyRole(role string) bool {
	return role == CompanyRole
}

func RemoveSpecialChars(input string) string {
	// In this pattern, [^a-zA-Z0-9 ] matches any character that is not a letter, digit, or space.
	re := regexp.MustCompile("[^a-zA-Z0-9 ]")
	result := re.ReplaceAllString(input, "")

	return result
}
