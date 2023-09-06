package domain

type EmailData struct {
	URL     string
	Subject string
	Name    string
	Email   string
}

const (
	StudentConfirmEmail = "student_confirm_email"
	CompanyApproveEmail = "company_approve_email"
	CompanyRejectEmail = "company_reject_email"
)