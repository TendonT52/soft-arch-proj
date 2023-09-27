package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/port"
	"github.com/TikhampornSky/go-auth-verifiedMail/utils"
)

// Same idea as handler/auth.go but this is for gRPC
type AuthServer struct {
	AuthService port.AuthServicePort
	pbv1.UnimplementedAuthServiceServer
}

func NewAuthServer(s port.AuthServicePort) *AuthServer {
	return &AuthServer{
		AuthService: s,
	}
}

func (s *AuthServer) AuthHealthCheck(context.Context, *pbv1.AuthHealthCheckRequest) (*pbv1.AuthHealthCheckResponse, error) {
	log.Println("Auth HealthCheck success: ", http.StatusOK)
	return &pbv1.AuthHealthCheckResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *AuthServer) CreateStudent(ctx context.Context, req *pbv1.CreateStudentRequest) (*pbv1.CreateStudentResponse, error) {
	id, err := s.AuthService.SignUpStudent(ctx, req)
	if errors.Is(err, domain.ErrYearMustBeGreaterThanZero) {
		log.Printf("Year must be greater than zero: %v", err)
		return &pbv1.CreateStudentResponse{
			Status:  http.StatusBadRequest,
			Message: "Year must be greater than zero",
		}, nil
	}
	if errors.Is(err, domain.ErrPasswordNotMatch) {
		log.Printf("Passwords do not match: %v", err)
		return &pbv1.CreateStudentResponse{
			Status:  http.StatusBadRequest,
			Message: "Passwords do not match",
		}, nil
	}
	if errors.Is(err, domain.ErrNotChulaStudentEmail) {
		log.Printf("Email must be studentID with @student.chula.ac.th: %v", err)
		return &pbv1.CreateStudentResponse{
			Status:  http.StatusBadRequest,
			Message: "Email must be studentID with @student.chula.ac.th",
		}, nil
	}
	if errors.Is(err, domain.ErrDuplicateEmail) {
		log.Printf("Email already exists: %v", err)
		return &pbv1.CreateStudentResponse{
			Status:  http.StatusBadRequest,
			Message: "Email already exists",
		}, nil
	}
	if err != nil {
		log.Printf("Error: %v", err)
		return &pbv1.CreateStudentResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		}, nil
	}

	log.Println("CreateStudent success: ", http.StatusCreated)
	return &pbv1.CreateStudentResponse{
		Status:  http.StatusCreated,
		Message: "Your account has been created. Please verify your email",
		Id:      id,
	}, nil
}

func (s *AuthServer) CreateCompany(ctx context.Context, req *pbv1.CreateCompanyRequest) (*pbv1.CreateCompanyResponse, error) {
	id, err := s.AuthService.SignUpCompany(ctx, req)
	if errors.Is(err, domain.ErrPasswordNotMatch) {
		log.Printf("Passwords do not match: %v", err)
		return &pbv1.CreateCompanyResponse{
			Status:  http.StatusBadRequest,
			Message: "Passwords do not match",
		}, nil
	}
	if errors.Is(err, domain.ErrDuplicateEmail) {
		log.Printf("Email already exists: %v", err)
		return &pbv1.CreateCompanyResponse{
			Status:  http.StatusBadRequest,
			Message: "Email already exists",
		}, nil
	}
	if err != nil {
		log.Printf("Error: %v", err)
		return &pbv1.CreateCompanyResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		}, nil
	}

	log.Println("CreateCompany success: ", http.StatusCreated)
	return &pbv1.CreateCompanyResponse{
		Status:  http.StatusCreated,
		Message: "The Approval process will take 1-2 days. Thank you for your patience",
		Id:      id,
	}, nil
}

func (s *AuthServer) CreateAdmin(ctx context.Context, req *pbv1.CreateAdminRequest) (*pbv1.CreateAdminResponse, error) {
	payload, err := utils.ValidateAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.CreateAdminResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}
	if payload.Role != "admin" {
		log.Println("Error: ", err)
		return &pbv1.CreateAdminResponse{
			Status:  http.StatusForbidden,
			Message: "You are not admin",
		}, nil
	}
	id, err := s.AuthService.SignUpAdmin(ctx, req)
	if errors.Is(err, domain.ErrPasswordNotMatch) {
		log.Printf("Passwords do not match: %v", err)
		return &pbv1.CreateAdminResponse{
			Status:  http.StatusBadRequest,
			Message: "Passwords do not match",
		}, nil
	}
	if errors.Is(err, domain.ErrDuplicateEmail) {
		log.Printf("Email already exists: %v", err)
		return &pbv1.CreateAdminResponse{
			Status:  http.StatusBadRequest,
			Message: "Email already exists",
		}, nil
	}
	if err != nil {
		log.Printf("Error: %v", err)
		return &pbv1.CreateAdminResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		}, nil
	}

	log.Println("CreateAdmin success: ", http.StatusCreated)
	return &pbv1.CreateAdminResponse{
		Status:  http.StatusCreated,
		Message: "Welcome to admin world!",
		Id:      id,
	}, nil
}

func (s *AuthServer) SignIn(ctx context.Context, req *pbv1.LoginRequest) (*pbv1.LoginResponse, error) {
	access_token, refresh_token, err := s.AuthService.SignIn(ctx, req)
	if errors.Is(err, domain.ErrNotVerified) {
		log.Printf("Error: %v", err)
		return &pbv1.LoginResponse{
			Status:  http.StatusBadRequest,
			Message: "Your account is not verified",
		}, nil
	}
	if errors.Is(err, domain.ErrPasswordNotMatch) {
		log.Printf("Passwords do not match: %v", err)
		return &pbv1.LoginResponse{
			Status:  http.StatusBadRequest,
			Message: "Passwords do not match",
		}, nil
	}
	if err != nil {
		log.Printf("Error: %v", err)
		return &pbv1.LoginResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		}, nil
	}

	log.Println("SignIn success: ", http.StatusOK)
	return &pbv1.LoginResponse{
		Status:       http.StatusOK,
		Message:      "Login success",
		AccessToken:  access_token,
		RefreshToken: refresh_token,
	}, nil
}

func (s *AuthServer) RefreshToken(ctx context.Context, req *pbv1.RefreshTokenRequest) (*pbv1.RefreshTokenResponse, error) {
	access_token, err := s.AuthService.RefreshAccessToken(ctx, req.RefreshToken)
	if errors.Is(err, domain.ErrUserIDNotFound) {
		log.Printf("the user belonging to this token no logger exists: %v", err)
		return &pbv1.RefreshTokenResponse{
			Status:  http.StatusForbidden,
			Message: "the user belonging to this token no logger exists",
		}, nil
	}

	if err != nil {
		log.Printf("Error: %v", err)
		return &pbv1.RefreshTokenResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	log.Println("RefreshToken success: ", http.StatusOK)
	return &pbv1.RefreshTokenResponse{
		Status:      http.StatusOK,
		Message:     "Refresh token success",
		AccessToken: access_token,
	}, nil
}

func (s *AuthServer) LogOut(ctx context.Context, req *pbv1.LogOutRequest) (*pbv1.LogOutResponse, error) {
	err := s.AuthService.LogOut(ctx, req.RefreshToken)
	if err != nil {
		log.Printf("Error: %v", err)
		return &pbv1.LogOutResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		}, nil
	}

	log.Println("Logout success: ", http.StatusOK)
	return &pbv1.LogOutResponse{
		Status:  http.StatusOK,
		Message: "Logout success",
	}, nil
}

func (s *AuthServer) VerifyEmailCode(ctx context.Context, req *pbv1.VerifyEmailCodeRequest) (*pbv1.VerifyEmailCodeResponse, error) {
	err := s.AuthService.VerifyEmail(ctx, req.StudentId, req.Code)
	if errors.Is(err, domain.ErrAlreadyVerified) {
		log.Printf("Your account has already been verified: %v", err)
		return &pbv1.VerifyEmailCodeResponse{
			Status:  http.StatusBadRequest,
			Message: "Your account has already been verified",
		}, nil
	}
	if err != nil {
		log.Printf("Error: %v", err)
		return &pbv1.VerifyEmailCodeResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid verification code or user doesn't exists",
		}, nil
	}

	log.Println("VerifyEmailCode success: ", http.StatusCreated)
	return &pbv1.VerifyEmailCodeResponse{
		Status:  http.StatusOK,
		Message: "verify success",
	}, nil
}
