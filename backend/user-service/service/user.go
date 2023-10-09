package service

import (
	"context"
	"fmt"

	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	"github.com/TikhampornSky/go-auth-verifiedMail/email"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/port"
)

type userService struct {
	repo    port.UserRepoPort
	memphis port.MemphisPort
}

func NewUserService(repo port.UserRepoPort, m port.MemphisPort) port.UserServicePort {
	return &userService{
		repo:    repo,
		memphis: m,
	}
}

func (s *userService) GetStudentMe(ctx context.Context, id int64) (*pbv1.Student, error) {
	student, err := s.repo.GetStudentByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (s *userService) GetCompanyMe(ctx context.Context, id int64) (*pbv1.Company, error) {
	company, err := s.repo.GetCompanyByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (s *userService) GetStudentByID(ctx context.Context, userId, id int64) (*pbv1.Student, error) {
	_, err := s.repo.CheckUserIDExist(ctx, userId)
	if err != nil {
		return nil, domain.ErrUserIDNotFound.With("the user belonging to this token no logger exists")
	}

	student, err := s.repo.GetStudentByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (s *userService) GetCompanyByID(ctx context.Context, userId, id int64) (*pbv1.Company, error) {
	_, err := s.repo.CheckUserIDExist(ctx, userId)
	if err != nil {
		return nil, domain.ErrUserIDNotFound.With("the user belonging to this token no logger exists")
	}

	company, err := s.repo.GetCompanyByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (s *userService) GetAllCompany(ctx context.Context, userId int64) ([]*pbv1.Company, error) {
	err := s.repo.CheckIfAdmin(ctx, userId)
	if err != nil {
		return nil, domain.ErrForbidden.With("user not admin")
	}

	companies, err := s.repo.GetAllCompany(ctx)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (s *userService) GetApprovedCompany(ctx context.Context, userId int64, search string) ([]*pbv1.Company, error) {
	_, err := s.repo.CheckUserIDExist(ctx, userId)
	if err != nil {
		return nil, domain.ErrUserIDNotFound.With("the user belonging to this token no logger exists")
	}
	searchPlain := domain.RemoveSpecialChars(search)
	companies, err := s.repo.GetApprovedCompany(ctx, searchPlain)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (s *userService) UpdateStudentMe(ctx context.Context, id int64, req *pbv1.UpdatedStudent) error {
	role, err := s.repo.CheckUserIDExist(ctx, id)
	if err != nil {
		return domain.ErrUserIDNotFound.With("the user belonging to this token no logger exists")
	}
	if role != domain.StudentRole {
		return domain.ErrForbidden.With("user not student")
	}
	if req.Year <= 0 {
		return domain.ErrYearMustBeGreaterThanZero.With("year must be greater than zero")
	}
	if !domain.CheckStudentRequiredFields(req) {
		return domain.ErrFieldsAreRequired
	}

	err = s.repo.UpdateStudentByID(ctx, id, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) UpdateCompanyMe(ctx context.Context, id int64, req *pbv1.UpdatedCompany) error {
	role, err := s.repo.CheckUserIDExist(ctx, id)
	if err != nil {
		return domain.ErrUserIDNotFound.With("the user belonging to this token no logger exists")
	}
	if role != domain.CompanyRole {
		return domain.ErrForbidden.With("user not company")
	}
	if !domain.CheckCompanyRequiredFields(req) {
		return domain.ErrFieldsAreRequired
	}

	err = s.repo.UpdateCompanyByID(ctx, id, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) UpdateCompanyStatus(ctx context.Context, userId, id int64, status string) error {
	var typeEmail string
	if status == domain.ComapanyStatusApprove {
		typeEmail = domain.CompanyApproveEmail
	} else if status == domain.ComapanyStatusReject {
		typeEmail = domain.CompanyRejectEmail
	} else {
		return domain.ErrInvalidStatus.With("status must be Approve or Reject")
	}

	err := s.repo.CheckIfAdmin(ctx, userId)
	if err != nil {
		return domain.ErrForbidden.With("user not admin")
	}

	company, err := s.repo.GetCompanyByID(ctx, id)
	if err != nil {
		return err
	}

	if company.Status != domain.ComapanyStatusPending {
		return domain.ErrAlreadyVerified.With("company already approved or rejected")
	}

	err = email.SendEmail(s.memphis, typeEmail, "", status+" Company", company.Name, company.Email)
	if err != nil {
		fmt.Println("Error:", err)
		return domain.ErrMailNotSent.With("cannot send email")
	}

	err = s.repo.UpdateCompanyStatus(ctx, id, status)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteStudent(ctx context.Context, userId, id int64) error {
	role, err := s.repo.CheckUserIDExist(ctx, userId)
	if err != nil {
		return domain.ErrUserIDNotFound.With("the user belonging to this token no logger exists")
	}
	if role != domain.AdminRole {
		return domain.ErrForbidden.With("user not admin")
	}

	err = s.repo.DeleteStudent(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteCompany(ctx context.Context, userId, id int64) error {
	role, err := s.repo.CheckUserIDExist(ctx, userId)
	if err != nil {
		return domain.ErrUserIDNotFound.With("the user belonging to this token no logger exists")
	}
	if role != domain.AdminRole {
		return domain.ErrForbidden.With("user not admin")
	}

	err = s.repo.DeleteCompany(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteCompanies(ctx context.Context, userId int64) error {
	err := s.repo.CheckIfAdmin(ctx, userId)
	if err != nil {
		return domain.ErrForbidden.With("user not admin")
	}

	err = s.repo.DeleteCompanies(ctx)
	if err != nil {
		return err
	}

	return nil
}
