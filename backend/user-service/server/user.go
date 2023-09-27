package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/TikhampornSky/go-auth-verifiedMail/domain"
	pbv1 "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-auth-verifiedMail/port"
	"github.com/TikhampornSky/go-auth-verifiedMail/utils"
)

type UserServer struct {
	UserService port.UserServicePort
	pbv1.UnimplementedUserServiceServer
}

func NewUserServer(s port.UserServicePort) *UserServer {
	return &UserServer{
		UserService: s,
	}
}

func (s *UserServer) UserHealthCheck(context.Context, *pbv1.UserHealthCheckRequest) (*pbv1.UserHealthCheckResponse, error) {
	log.Println("User HealthCheck success: ", http.StatusOK)
	return &pbv1.UserHealthCheckResponse{
		Status: http.StatusOK,
	}, nil
}

// Student Zone
func (s *UserServer) GetStudentMe(ctx context.Context, req *pbv1.GetStudentMeRequest) (*pbv1.GetStudentResponse, error) {
	payload, err := utils.ValidateAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.GetStudentResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}

	res, err := s.UserService.GetStudentMe(ctx, payload.UserId)
	if err != nil {
		log.Println("Error from get student (Me): ", err)
		return &pbv1.GetStudentResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		}, nil
	}

	log.Println("Success getting student me: ", res)
	return &pbv1.GetStudentResponse{
		Status:  http.StatusOK,
		Message: "success",
		Student: res,
	}, nil
}

func (s *UserServer) GetStudent(ctx context.Context, req *pbv1.GetStudentRequest) (*pbv1.GetStudentResponse, error) {
	payload, err := utils.ValidateAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.GetStudentResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}

	res, err := s.UserService.GetStudentByID(ctx, payload.UserId, req.Id)
	if errors.Is(err, domain.ErrUserIDNotFound) {
		log.Println("Error UserId not found: ", err)
		return &pbv1.GetStudentResponse{
			Status:  http.StatusNotFound,
			Message: "user id not found",
		}, nil
	}
	if err != nil {
		log.Println("Error from get student: ", err)
		return &pbv1.GetStudentResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		}, nil
	}

	log.Println("Success getting student: ", res)
	return &pbv1.GetStudentResponse{
		Status:  http.StatusOK,
		Message: "success",
		Student: res,
	}, nil
}

func (s *UserServer) UpdateStudent(ctx context.Context, req *pbv1.UpdateStudentRequest) (*pbv1.UpdateStudentResponse, error) {
	payload, err := utils.ValidateAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.UpdateStudentResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}

	err = s.UserService.UpdateStudentMe(ctx, payload.UserId, req.Student)
	if errors.Is(err, domain.ErrUserIDNotFound) {
		log.Println("Error UserId not found: ", err)
		return &pbv1.UpdateStudentResponse{
			Status:  http.StatusNotFound,
			Message: "user id not found",
		}, nil
	}
	if errors.Is(err, domain.ErrForbidden) {
		log.Println("Error NOT Authorize: ", err)
		return &pbv1.UpdateStudentResponse{
			Status:  http.StatusForbidden,
			Message: "You are not authorized to update this student",
		}, nil
	}
	if err != nil {
		log.Println("Error from update student: ", err)
		return &pbv1.UpdateStudentResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		}, nil
	}

	log.Println("Success updating student: ", req.Student)
	message := "Update data for " + req.Student.Name + " successfully!"
	return &pbv1.UpdateStudentResponse{
		Status:  http.StatusOK,
		Message: message,
	}, nil
}

// Company Zone
func (s *UserServer) GetCompanyMe(ctx context.Context, req *pbv1.GetCompanyMeRequest) (*pbv1.GetCompanyResponse, error) {
	payload, err := utils.ValidateAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.GetCompanyResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}

	res, err := s.UserService.GetCompanyMe(ctx, payload.UserId)
	if err != nil {
		log.Println("Error from get company (Me): ", err)
		return &pbv1.GetCompanyResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		}, nil
	}

	log.Println("Success getting company me: ", res)
	return &pbv1.GetCompanyResponse{
		Status:  http.StatusOK,
		Message: "success",
		Company: res,
	}, nil

}

func (s *UserServer) GetCompany(ctx context.Context, req *pbv1.GetCompanyRequest) (*pbv1.GetCompanyResponse, error) {
	payload, err := utils.ValidateAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.GetCompanyResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}

	res, err := s.UserService.GetCompanyByID(ctx, payload.UserId, req.Id)
	if errors.Is(err, domain.ErrUserIDNotFound) {
		log.Println("Error UserId not found: ", err)
		return &pbv1.GetCompanyResponse{
			Status:  http.StatusNotFound,
			Message: "company id not found",
		}, nil
	}
	if err != nil {
		log.Println("Error from get company: ", err)
		return &pbv1.GetCompanyResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		}, nil
	}

	log.Println("Success getting company: ", res)
	return &pbv1.GetCompanyResponse{
		Status:  http.StatusOK,
		Message: "success",
		Company: res,
	}, nil
}

func (s *UserServer) UpdateCompany(ctx context.Context, req *pbv1.UpdateCompanyRequest) (*pbv1.UpdateCompanyResponse, error) {
	payload, err := utils.ValidateAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.UpdateCompanyResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}

	err = s.UserService.UpdateCompanyMe(ctx, payload.UserId, req.Company)
	if errors.Is(err, domain.ErrUserIDNotFound) {
		log.Println("Error UserId not found: ", err)
		return &pbv1.UpdateCompanyResponse{
			Status:  http.StatusNotFound,
			Message: "user id not found",
		}, nil
	}
	if errors.Is(err, domain.ErrForbidden) {
		log.Println("Error NOT Authorize: ", err)
		return &pbv1.UpdateCompanyResponse{
			Status:  http.StatusForbidden,
			Message: "You are not authorized to update this company",
		}, nil
	}
	if err != nil {
		log.Println("Error from update company: ", err)
		return &pbv1.UpdateCompanyResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		}, nil
	}

	log.Println("Success updating company: ", req.Company)
	message := "Update data for " + req.Company.Name + " successfully!"
	return &pbv1.UpdateCompanyResponse{
		Status:  http.StatusOK,
		Message: message,
	}, nil
}

func (s *UserServer) ListCompanies(ctx context.Context, req *pbv1.ListCompaniesRequest) (*pbv1.ListCompaniesResponse, error) {
	payload, err := utils.ValidateAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.ListCompaniesResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}

	res, err := s.UserService.GetAllCompany(ctx, payload.UserId)
	if errors.Is(err, domain.ErrForbidden) {
		log.Println("Error NOT Authorize: ", err)
		return &pbv1.ListCompaniesResponse{
			Status:  http.StatusForbidden,
			Message: "Only admin can view",
		}, nil
	}
	if err != nil {
		log.Println("Error from list companies: ", err)
		return &pbv1.ListCompaniesResponse{
			Status:  http.StatusBadRequest,
			Message: "Something went wrong",
		}, nil
	}

	log.Println("Success listing all companies: ", len(res))
	return &pbv1.ListCompaniesResponse{
		Status:    http.StatusOK,
		Message:   "success",
		Companies: res,
		Total:     int64(len(res)),
	}, nil
}

func (s *UserServer) ListApprovedCompanies(ctx context.Context, req *pbv1.ListApprovedCompaniesRequest) (*pbv1.ListApprovedCompaniesResponse, error) {
	payload, err := utils.ValidateAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.ListApprovedCompaniesResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}

	res, err := s.UserService.GetApprovedCompany(ctx, payload.UserId, req.Search)
	if err != nil {
		log.Println("Error from list approved companies: ", err)
		return &pbv1.ListApprovedCompaniesResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}, nil
	}

	log.Println("Success listing all approved companies: ", len(res))
	return &pbv1.ListApprovedCompaniesResponse{
		Status:    http.StatusOK,
		Message:   "success",
		Companies: res,
		Total:     int64(len(res)),
	}, nil
}

func (s *UserServer) UpdateCompanyStatus(ctx context.Context, req *pbv1.UpdateCompanyStatusRequest) (*pbv1.UpdateCompanyStatusResponse, error) {
	payload, err := utils.ValidateAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.UpdateCompanyStatusResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}

	err = s.UserService.UpdateCompanyStatus(ctx, payload.UserId, req.Id, req.Status)
	if errors.Is(err, domain.ErrForbidden) {
		return &pbv1.UpdateCompanyStatusResponse{
			Status:  http.StatusForbidden,
			Message: "Only admin can approve",
		}, nil
	}
	if errors.Is(err, domain.ErrAlreadyVerified) || errors.Is(err, domain.ErrInvalidStatus) || errors.Is(err, domain.ErrMailNotSent) {
		log.Println("Error from update company status: ", err)
		return &pbv1.UpdateCompanyStatusResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}
	if err != nil {
		log.Println("Error internal when update company status: ", err)
		return &pbv1.UpdateCompanyStatusResponse{
			Status:  http.StatusInternalServerError,
			Message: "Something went wrong",
		}, nil
	}

	log.Println("Success updating company status: ", req.Id)
	message := "Update status for company id " + strconv.FormatInt(req.Id, 10) + " successfully!"
	return &pbv1.UpdateCompanyStatusResponse{
		Status:  http.StatusOK,
		Message: message,
	}, nil
}
