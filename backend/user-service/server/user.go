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

// Student Zone
func (s *UserServer) GetStudentMe(ctx context.Context, req *pbv1.GetStudentMeRequest) (*pbv1.GetStudentResponse, error) {
	id, err := utils.ExtractUserIDFromAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.GetStudentResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}
	res, err := s.UserService.GetStudentMe(ctx, id)
	if err != nil {
		log.Println("Error from get student (Me): ", err)
		return &pbv1.GetStudentResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
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
	userId, err := utils.ExtractUserIDFromAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.GetStudentResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}
	res, err := s.UserService.GetStudentByID(ctx, userId, req.Id)
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
			Message: err.Error(),
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
	userId, err := utils.ExtractUserIDFromAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.UpdateStudentResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}

	err = s.UserService.UpdateStudentMe(ctx, userId, req.Student)
	if errors.Is(err, domain.ErrUserIDNotFound) {
		log.Println("Error UserId not found: ", err)
		return &pbv1.UpdateStudentResponse{
			Status:  http.StatusNotFound,
			Message: "user id not found",
		}, nil
	}
	if errors.Is(err, domain.ErrNotAuthorized) {
		log.Println("Error NOT Authorize: ", err)
		return &pbv1.UpdateStudentResponse{
			Status:  http.StatusUnauthorized,
			Message: "You are not authorized to update this student",
		}, nil
	}
	if err != nil {
		log.Println("Error from update student: ", err)
		return &pbv1.UpdateStudentResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
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
	id, err := utils.ExtractUserIDFromAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.GetCompanyResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}
	res, err := s.UserService.GetCompanyMe(ctx, id)
	if err != nil {
		log.Println("Error from get company (Me): ", err)
		return &pbv1.GetCompanyResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
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
	userId, err := utils.ExtractUserIDFromAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.GetCompanyResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}

	res, err := s.UserService.GetCompanyByID(ctx, userId, req.Id)
	if errors.Is(err, domain.ErrUserIDNotFound) {
		log.Println("Error UserId not found: ", err)
		return &pbv1.GetCompanyResponse{
			Status:  http.StatusNotFound,
			Message: "user id not found",
		}, nil
	}
	if err != nil {
		log.Println("Error from get company: ", err)
		return &pbv1.GetCompanyResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
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
	userId, err := utils.ExtractUserIDFromAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.UpdateCompanyResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}

	err = s.UserService.UpdateCompanyMe(ctx, userId, req.Company)
	if errors.Is(err, domain.ErrUserIDNotFound) {
		log.Println("Error UserId not found: ", err)
		return &pbv1.UpdateCompanyResponse{
			Status:  http.StatusNotFound,
			Message: "user id not found",
		}, nil
	}
	if errors.Is(err, domain.ErrNotAuthorized) {
		log.Println("Error NOT Authorize: ", err)
		return &pbv1.UpdateCompanyResponse{
			Status:  http.StatusUnauthorized,
			Message: "You are not authorized to update this company",
		}, nil
	}
	if err != nil {
		log.Println("Error from update company: ", err)
		return &pbv1.UpdateCompanyResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
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
	userId, err := utils.ExtractUserIDFromAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.ListCompaniesResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}

	res, err := s.UserService.GetAllCompany(ctx, userId)
	if errors.Is(err, domain.ErrNotAuthorized) {
		log.Println("Error NOT Authorize: ", err)
		return &pbv1.ListCompaniesResponse{
			Status:  http.StatusUnauthorized,
			Message: "Only admin can view",
		}, nil
	}
	if err != nil {
		log.Println("Error from list companies: ", err)
		return &pbv1.ListCompaniesResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
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
	userId, err := utils.ExtractUserIDFromAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.ListApprovedCompaniesResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}

	res, err := s.UserService.GetApprovedCompany(ctx, userId, req.Search)
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
	userId, err := utils.ExtractUserIDFromAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.UpdateCompanyStatusResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}

	err = s.UserService.UpdateCompanyStatus(ctx, userId, req.Id, req.Status)
	if errors.Is(err, domain.ErrNotAuthorized) {
		return &pbv1.UpdateCompanyStatusResponse{
			Status:  http.StatusUnauthorized,
			Message: "Only admin can approve",
		}, nil
	}
	if err != nil {
		log.Println("Error from update company status: ", err)
		return &pbv1.UpdateCompanyStatusResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}

	log.Println("Success updating company status: ", req.Id)
	message := "Update status for company id " + strconv.FormatInt(req.Id, 10) + " successfully!"
	return &pbv1.UpdateCompanyStatusResponse{
		Status:  http.StatusOK,
		Message: message,
	}, nil
}

func (s *UserServer) DeleteCompanies(ctx context.Context, req *pbv1.DeleteCompaniesRequest) (*pbv1.DeleteCompaniesResponse, error) {
	userId, err := utils.ExtractUserIDFromAccessToken(req.AccessToken)
	if err != nil {
		log.Println("Error in extract userID: ", err)
		return &pbv1.DeleteCompaniesResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}

	err = s.UserService.DeleteCompanies(ctx, userId)
	if errors.Is(err, domain.ErrNotAuthorized) {
		return &pbv1.DeleteCompaniesResponse{
			Status:  http.StatusUnauthorized,
			Message: "Only admin can delete",
		}, nil
	}
	if err != nil {
		log.Println("Error from delete company: ", err)
		return &pbv1.DeleteCompaniesResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}

	log.Println("Success deleting companies!")
	message := "Delete companies successfully!"
	return &pbv1.DeleteCompaniesResponse{
		Status:  http.StatusOK,
		Message: message,
	}, nil
}