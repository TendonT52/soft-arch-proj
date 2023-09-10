package domain

type EmailData struct {
	URL     string `json:"url"`
	Subject string `json:"subject"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}

const (
	StudentConfirmEmail = "student_confirm_email"
	CompanyApproveEmail = "company_approve_email"
	CompanyRejectEmail = "company_reject_email"
)