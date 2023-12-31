package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/JinnnDamanee/review-service/domain"
	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
	"github.com/JinnnDamanee/review-service/port"
	"github.com/JinnnDamanee/review-service/utils"
)

type ReviewServer struct {
	ReviewService port.ReviewServicePort
	pbv1.UnimplementedReviewServiceServer
}

func NewReviewServer(reviewService port.ReviewServicePort) *ReviewServer {
	return &ReviewServer{ReviewService: reviewService}
}

func (s *ReviewServer) ReviewHealthCheck(ctx context.Context, req *pbv1.ReviewHealthCheckRequest) (*pbv1.ReviewHealthCheckResponse, error) {
	log.Println("Review HealthCheck success: ", http.StatusOK)
	return &pbv1.ReviewHealthCheckResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *ReviewServer) CreateReview(ctx context.Context, req *pbv1.CreateReviewRequest) (*pbv1.CreateReviewResponse, error) {
	res, err := s.ReviewService.CreateReview(ctx, req.AccessToken, req.Review)
	if errors.Is(err, domain.ErrFieldsAreRequired) {
		log.Println("Create Review: Please fill in all required fields")
		return &pbv1.CreateReviewResponse{
			Status:  http.StatusBadRequest,
			Message: "Please fill in all required fields",
		}, nil
	}
	if errors.Is(err, domain.ErrRatingRange) {
		log.Println("Create Review: Rating must be between 1 and 5")
		return &pbv1.CreateReviewResponse{
			Status:  http.StatusBadRequest,
			Message: "Rating must be between 1 and 5",
		}, nil
	}
	if errors.Is(err, domain.ErrCompanyNotFound) {
		log.Println("Create Review: Company not found")
		return &pbv1.CreateReviewResponse{
			Status:  http.StatusBadRequest,
			Message: "Company not found",
		}, nil
	}
	if errors.Is(err, domain.ErrUnauthorize) {
		log.Println("Create Review: Your access token is invalid")
		return &pbv1.CreateReviewResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}
	if errors.Is(err, domain.ErrForbidden) {
		log.Println("Create Review: You are not allowed to create review")
		return &pbv1.CreateReviewResponse{
			Status:  http.StatusForbidden,
			Message: "You are not allowed to create review",
		}, nil
	}
	if err != nil {
		log.Println("Create Review: ", err)
		return &pbv1.CreateReviewResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}

	return &pbv1.CreateReviewResponse{
		Status:  http.StatusCreated,
		Message: "Review created successfully",
		Id:      res,
	}, nil
}

func (s *ReviewServer) ListReviewsByCompany(ctx context.Context, req *pbv1.ListReviewsByCompanyRequest) (*pbv1.ListReviewsByCompanyResponse, error) {
	res, err := s.ReviewService.GetReviewsByCompany(ctx, req.AccessToken, req.Cid)
	if errors.Is(err, domain.ErrUnauthorize) {
		log.Println("List Reviews By Company: Your access token is invalid")
		return &pbv1.ListReviewsByCompanyResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}
	if err != nil {
		log.Println("List Reviews By Company: ", err)
		return &pbv1.ListReviewsByCompanyResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}
	log.Println("List Reviews By Company success: Total is ", int32(len(res)))

	return &pbv1.ListReviewsByCompanyResponse{
		Status:  http.StatusOK,
		Message: "List reviews by company successfully",
		Reviews: res,
		Total:   int32(len(res)),
	}, nil
}

func (s *ReviewServer) GetReview(ctx context.Context, req *pbv1.GetReviewRequest) (*pbv1.GetReviewResponse, error) {
	res, err := s.ReviewService.GetReviewByID(ctx, req.AccessToken, req.Id)
	if errors.Is(err, domain.ErrUnauthorize) {
		log.Println("Get Review: Your access token is invalid")
		return &pbv1.GetReviewResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}
	if errors.Is(err, domain.ErrReviewNotFound) {
		log.Println("Get Review: Review not found")
		return &pbv1.GetReviewResponse{
			Status:  http.StatusNotFound,
			Message: "Review not found",
		}, nil
	}
	if err != nil {
		log.Println("Get Review: ", err)
		return &pbv1.GetReviewResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}

	log.Println("Get Review success: ", http.StatusOK)
	return &pbv1.GetReviewResponse{
		Status:  http.StatusOK,
		Message: "Get review successfully",
		Review:  res,
	}, nil
}

func (s *ReviewServer) UpdateReview(ctx context.Context, req *pbv1.UpdateReviewRequest) (*pbv1.UpdateReviewResponse, error) {
	err := s.ReviewService.UpdateReview(ctx, req.AccessToken, req.Review, req.Id)
	if errors.Is(err, domain.ErrFieldsAreRequired) {
		log.Println("Update Review: Please fill in all required fields")
		return &pbv1.UpdateReviewResponse{
			Status:  http.StatusBadRequest,
			Message: "Please fill in all required fields",
		}, nil
	}
	if errors.Is(err, domain.ErrRatingRange) {
		log.Println("Update Review: Rating must be between 1 and 5")
		return &pbv1.UpdateReviewResponse{
			Status:  http.StatusBadRequest,
			Message: "Rating must be between 1 and 5",
		}, nil
	}
	if errors.Is(err, domain.ErrUnauthorize) {
		log.Println("Update Review: Your access token is invalid")
		return &pbv1.UpdateReviewResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}
	if errors.Is(err, domain.ErrForbidden) {
		log.Println("Update Review: You are not allowed to update review")
		return &pbv1.UpdateReviewResponse{
			Status:  http.StatusForbidden,
			Message: "You are not allowed to update review",
		}, nil
	}
	if err != nil {
		log.Println("Update Review: ", err)
		return &pbv1.UpdateReviewResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}

	return &pbv1.UpdateReviewResponse{
		Status:  http.StatusOK,
		Message: "Update review successfully",
	}, nil
}

func (s *ReviewServer) ListReviewsByUser(ctx context.Context, req *pbv1.ListReviewsByUserRequest) (*pbv1.ListReviewsByUserResponse, error) {
	payload, err := utils.ValidateAccessToken(req.AccessToken)
	if err != nil {
		log.Println("List Reviews By User: Your access token is invalid")
		return &pbv1.ListReviewsByUserResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}

	myReviews, err := s.ReviewService.GetReviewsByUser(ctx, req.AccessToken, payload.UserId)
	if errors.Is(err, domain.ErrUnauthorize) {
		log.Println("List Reviews By User: Your access token is invalid")
		return &pbv1.ListReviewsByUserResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}
	if errors.Is(err, domain.ErrForbidden) {
		log.Println("List Reviews By User: You are not allowed to get reviews")
		return &pbv1.ListReviewsByUserResponse{
			Status:  http.StatusForbidden,
			Message: "You are not allowed to view these reviews",
		}, nil
	}
	if err != nil {
		log.Println("List Review by User : ", err)
		return &pbv1.ListReviewsByUserResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}
	return &pbv1.ListReviewsByUserResponse{
		Status:  http.StatusOK,
		Message: "List reviews by user successfully",
		Reviews: myReviews,
		Total:   int32(len(myReviews)),
	}, nil
}

func (s *ReviewServer) DeleteReview(ctx context.Context, req *pbv1.DeleteReviewRequest) (*pbv1.DeleteReviewResponse, error) {
	err := s.ReviewService.DeleteReview(ctx, req.AccessToken, req.Id)
	if errors.Is(err, domain.ErrUnauthorize) {
		log.Println("Delete Review: Your access token is invalid")
		return &pbv1.DeleteReviewResponse{
			Status:  http.StatusUnauthorized,
			Message: "Your access token is invalid",
		}, nil
	}
	if errors.Is(err, domain.ErrForbidden) {
		log.Println("Delete Review: You are not allowed to delete review")
		return &pbv1.DeleteReviewResponse{
			Status:  http.StatusForbidden,
			Message: "You are not allowed to delete review",
		}, nil
	}
	if errors.Is(err, domain.ErrReviewNotFound) {
		log.Println("Delete Review: Review not found")
		return &pbv1.DeleteReviewResponse{
			Status:  http.StatusNotFound,
			Message: "Review not found",
		}, nil
	}
	if err != nil {
		log.Println("Delete Review: ", err)
		return &pbv1.DeleteReviewResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
		}, nil
	}

	return &pbv1.DeleteReviewResponse{
		Status:  http.StatusOK,
		Message: "Delete review successfully",
	}, nil
}
