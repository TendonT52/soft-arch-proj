package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/TikhampornSky/go-auth-verifiedMail/config"
	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	"github.com/TikhampornSky/go-auth-verifiedMail/email"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/port"
	"github.com/TikhampornSky/go-auth-verifiedMail/utils"
)

type authService struct {
	repo    port.UserRepoPort
	memphis port.MemphisPort
	time    port.TimeProvider
}

func NewAuthService(repo port.UserRepoPort, m port.MemphisPort, t port.TimeProvider) port.AuthServicePort {
	return &authService{
		repo:    repo,
		memphis: m,
		time:    t,
	}
}

func (s *authService) SignUpStudent(ctx context.Context, req *pbv1.CreateStudentRequest) (int64, error) {
	if req.Password != req.PasswordConfirm {
		return 0, domain.ErrPasswordNotMatch
	}
	current_time := s.time.Now().Unix()

	hashedPassword := utils.HashPassword(req.Password, current_time)
	req.Password = hashedPassword

	if !email.IsChulaStudentEmail(req.Email) {
		return 0, domain.ErrNotChulaStudentEmail.With("email must be @student.chula.ac.th")
	}

	err := s.repo.CheckEmailExist(ctx, req.Email)
	if err != nil {
		return 0, domain.ErrDuplicateEmail
	}

	// Generate Verification Code
	id := email.GetStudentIDFromEmail(req.Email)
	code := utils.Encode(id, current_time)

	config, _ := config.LoadConfig("..")
	err = email.SendEmail(s.memphis, domain.StudentConfirmEmail, config.ClientOrigin+"/verifyemail/"+id+"/"+code, "Your account verification code", req.Name, req.Email)
	if err != nil {
		fmt.Println("Error:", err)
		return 0, domain.ErrMailNotSent.With("cannot send email")
	}

	sid, err := s.repo.CreateStudent(ctx, req, current_time)
	if err != nil {
		return 0, err
	}

	return sid, nil
}

func (s *authService) VerifyEmail(ctx context.Context, sid, code string) error {
	studentEmail := sid + "@student.chula.ac.th"
	salt, err := s.repo.GetSalt(ctx, studentEmail)
	if err != nil {
		return err
	}
	current_time, err := strconv.ParseInt(salt, 10, 64)
	if err != nil {
		return err
	}

	if !utils.Compare(sid, current_time, code) {
		return errors.New("invalid code")
	}

	err = s.repo.UpdateStudentStatus(ctx, studentEmail, true)
	if err != nil {
		return errors.New("cannot update student status")
	}

	return nil
}

func (s *authService) SignUpCompany(ctx context.Context, req *pbv1.CreateCompanyRequest) (int64, error) {
	if req.Password != req.PasswordConfirm {
		return 0, domain.ErrPasswordNotMatch
	}

	current_time := s.time.Now().Unix()
	hashedPassword := utils.HashPassword(req.Password, current_time)

	req.Password = hashedPassword

	err := s.repo.CheckEmailExist(ctx, req.Email)
	if err != nil {
		return 0, domain.ErrDuplicateEmail
	}

	createAt := s.time.Now().Unix()
	cid, err := s.repo.CreateCompany(ctx, req, createAt)
	if err != nil {
		return 0, err
	}

	return cid, nil
}

func (s *authService) SignUpAdmin(ctx context.Context, req *pbv1.CreateAdminRequest) (int64, error) {
	if req.Password != req.PasswordConfirm {
		return 0, domain.ErrPasswordNotMatch
	}

	current_time := s.time.Now().Unix()
	hashedPassword := utils.HashPassword(req.Password, current_time)
	req.Password = hashedPassword

	err := s.repo.CheckEmailExist(ctx, req.Email)
	if err != nil {
		return 0, domain.ErrDuplicateEmail
	}

	createAt := s.time.Now().Unix()
	aid, err := s.repo.CreateAdmin(ctx, req, createAt)
	if err != nil {
		return 0, err
	}

	return aid, nil
}

func (s *authService) SignIn(ctx context.Context, req *pbv1.LoginRequest) (string, string, error) {
	u, err := s.repo.GetUser(ctx, req)
	if err != nil {
		return "", "", err
	}
	if !u.Verified {
		return "", "", domain.ErrNotVerified.With("user not verified")
	}
	if err := utils.VerifyPassword(u.Password, req.Password, u.CreatedAt); err != nil {
		return "", "", domain.ErrPasswordNotMatch
	}

	// Generate token
	config, _ := config.LoadConfig("..")
	access_token, err := utils.CreateAccessToken(config.AccessTokenExpiresIn, &pbv1.Payload{
		UserId: u.Id,
		Role:   u.Role,
	})
	if err != nil {
		return "", "", err
	}

	fmt.Println("===> ", u.Id)
	refresh_token, err := utils.CreateRefreshToken(config.RefreshTokenExpiresIn, u.Id)
	if err != nil {
		return "", "", err
	}

	return access_token, refresh_token, nil
}

func (s *authService) RefreshAccessToken(ctx context.Context, refreshToken string) (string, error) {
	v, err := s.repo.GetValueRedis(ctx, refreshToken)
	if v != "" && err == nil {
		return "", domain.ErrInternal.With("your token has been logged out!")
	}

	config, _ := config.LoadConfig("..")
	userId, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	role, err := s.repo.CheckUserIDExist(ctx, userId)
	if err != nil {
		return "", domain.ErrUserIDNotFound.With("the user belonging to this token no logger exists")
	}

	access_token, err := utils.CreateAccessToken(config.AccessTokenExpiresIn, &pbv1.Payload{
		UserId: userId,
		Role:   role,
	})
	if err != nil {
		return "", err
	}

	return access_token, nil
}

func (s *authService) LogOut(ctx context.Context, refreshToken string) error {
	userId, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		return err
	}

	_, err = s.repo.CheckUserIDExist(ctx, userId)
	if err != nil {
		return domain.ErrUserIDNotFound.With("the user belonging to this token no logger exists")
	}

	err = s.repo.SetValueRedis(ctx, refreshToken, "logged_out")
	if err != nil {
		return err
	}

	return nil
}
