package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math"

	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	"github.com/TikhampornSky/go-auth-verifiedMail/email"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/initializers"
	"github.com/TikhampornSky/go-auth-verifiedMail/port"
	"github.com/TikhampornSky/go-auth-verifiedMail/utils"
	"github.com/thanhpk/randstr"
)

type authService struct {
	repo    port.UserRepoPort
	memphis port.MemphisPort
}

func NewAuthService(repo port.UserRepoPort, m port.MemphisPort) port.AuthServicePort {
	return &authService{
		repo:    repo,
		memphis: m,
	}
}

func (s *authService) SignUpStudent(ctx context.Context, req *pbv1.CreateStudentRequest) error {
	if req.Password != req.PasswordConfirm {
		return domain.ErrPasswordNotMatch
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	req.Password = hashedPassword

	if !email.IsChulaStudentEmail(req.Email) {
		return domain.ErrNotChulaStudentEmail.With("email must be @student.chula.ac.th")
	}

	err = s.repo.CheckEmailExist(ctx, req.Email)
	if err != nil {
		return domain.ErrDuplicateEmail
	}

	// Generate Verification Code
	code := randstr.String(20)
	verification_code := utils.Encode(code)

	// Send Email
	config, _ := initializers.LoadConfig(".")
	emailData := domain.EmailData{
		URL:     config.ClientOrigin + "/verifyemail/" + code,
		Subject: "Your account verification code",
		Name:    req.Name,
		Email:   req.Email,
	}

	jsonData, err := json.Marshal(emailData)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	err = email.SendEmail(s.memphis, domain.StudentConfirmEmail, jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return domain.ErrMailNotSent.With("cannot send email")
	}

	err = s.repo.CreateStudent(ctx, req, verification_code)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) VerifyEmail(ctx context.Context, code string) error {
	verification_code := utils.Encode(code)
	err := s.repo.UpdateVerificationCode(ctx, verification_code)
	if err != nil {
		return err
	}
	return nil
}

func (s *authService) SignUpCompany(ctx context.Context, req *pbv1.CreateCompanyRequest) error {
	if req.Password != req.PasswordConfirm {
		return domain.ErrPasswordNotMatch
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	req.Password = hashedPassword

	err = s.repo.CheckEmailExist(ctx, req.Email)
	if err != nil {
		return domain.ErrDuplicateEmail
	}

	err = s.repo.CreateCompany(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) SignUpAdmin(ctx context.Context, req *pbv1.CreateAdminRequest) error {
	if req.Password != req.PasswordConfirm {
		return domain.ErrPasswordNotMatch
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	req.Password = hashedPassword

	err = s.repo.CheckEmailExist(ctx, req.Email)
	if err != nil {
		return domain.ErrDuplicateEmail
	}

	err = s.repo.CreateAdmin(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) SignIn(ctx context.Context, req *pbv1.LoginRequest) (string, string, error) {
	id, password, err := s.repo.GetPassword(ctx, req)
	if err != nil {
		return "", "", err
	}
	if err := utils.VerifyPassword(password, req.Password); err != nil {
		return "", "", domain.ErrPasswordNotMatch
	}

	// Generate token
	config, _ := initializers.LoadConfig("..")
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, id, config.AccessTokenPrivateKey)
	if err != nil {
		return "", "", err
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, id, config.RefreshTokenPrivateKey)
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

	config, _ := initializers.LoadConfig("..")
	sub, err := utils.ValidateToken(refreshToken, config.RefreshTokenPublicKey)
	if err != nil {
		return "", err
	}

	val, ok := sub.(float64)
	if !ok {
		return "", domain.ErrInternal.With("cannot convert sub to float64")
	}
	_, err = s.repo.CheckUserIDExist(ctx, int64(math.Round(val)))
	if err != nil {
		return "", domain.ErrUserIDNotFound.With("the user belonging to this token no logger exists")
	}

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, int64(math.Round(val)), config.AccessTokenPrivateKey)
	if err != nil {
		return "", err
	}

	return access_token, nil
}

func (s *authService) LogOut(ctx context.Context, refreshToken string) error {
	config, _ := initializers.LoadConfig("..")
	sub, err := utils.ValidateToken(refreshToken, config.RefreshTokenPublicKey)
	if err != nil {
		return err
	}

	val, ok := sub.(float64)
	if !ok {
		return domain.ErrInternal.With("cannot convert sub to float64")
	}
	_, err = s.repo.CheckUserIDExist(ctx, int64(math.Round(val)))
	if err != nil {
		return domain.ErrUserIDNotFound.With("the user belonging to this token no logger exists")
	}

	// MUST USE REDIS TO LOGOUT
	// err = s.repo.SetValueRedis(ctx, refreshToken, "logged_out")
	// if err != nil {
	// 	return err
	// }

	return nil
}
