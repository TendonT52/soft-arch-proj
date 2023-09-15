package domain

const (
	StudentRole = "student"
	AdminRole   = "admin"
	CompanyRole = "company"
)

func IsStudentRole(role string) bool {
	return role == StudentRole
}

func IsAdminRole(role string) bool {
	return role == AdminRole
}

func IsCompanyRole(role string) bool {
	return role == CompanyRole
}
