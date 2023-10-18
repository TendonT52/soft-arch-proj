package service

import (
	"context"
	"database/sql"
	"sort"

	"github.com/JinnnDamanee/review-service/domain"
	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
	"github.com/JinnnDamanee/review-service/port"
	"github.com/JinnnDamanee/review-service/utils"
)

const (
	StudentRole   = "student"
	AnonymousName = "Anonymous"
)

type reviewService struct {
	repo        port.ReviewRepoPort
	userService port.UserClientPort
}

func NewReviewService(repo port.ReviewRepoPort, userService port.UserClientPort) port.ReviewServicePort {
	return &reviewService{repo: repo, userService: userService}
}

func (s *reviewService) CreateReview(ctx context.Context, token string, review *pbv1.CreatedReview) (int64, error) {
	if !domain.CheckCreatedRequireField(review) {
		return 0, domain.ErrFieldsAreRequired
	}
	if !domain.CheckRatingRange(review.Rating) {
		return 0, domain.ErrRatingRange
	}

	payload, err := utils.ValidateAccessToken(token)
	if err != nil {
		return 0, domain.ErrUnauthorize
	}

	if payload.Role != StudentRole {
		return 0, domain.ErrForbidden
	}

	// Check company exist
	reqCompany := &pbv1.GetCompanyRequest{
		AccessToken: token,
		Id:          review.Cid,
	}
	result, err := s.userService.GetCompanyProfile(ctx, reqCompany)
	if err != nil {
		return 0, err
	}
	if result.Status == 404 {
		return 0, domain.ErrCompanyNotFound
	}

	userID := payload.UserId
	reviewId, err := s.repo.CreateReview(ctx, userID, review)
	if err != nil {
		return 0, err
	}

	return reviewId, nil
}

func (s *reviewService) GetReviewByID(ctx context.Context, token string, reviewID int64) (*pbv1.Review, error) {
	_, err := utils.ValidateAccessToken(token)
	if err != nil {
		return nil, domain.ErrUnauthorize
	}

	review, err := s.repo.GetReviewByID(ctx, reviewID)
	if err != nil {
		return nil, err
	}

	// Get company info
	reqCompany := &pbv1.GetCompanyRequest{
		AccessToken: token,
		Id:          review.Company.Id,
	}
	result, err := s.userService.GetCompanyProfile(ctx, reqCompany)
	if err != nil {
		return nil, err
	}
	review.Company.Name = result.Company.Name

	// Get user info
	if review.Owner.Id == 0 {
		review.Owner.Name = AnonymousName
		return review, nil
	}

	reqStudent := &pbv1.GetStudentRequest{
		AccessToken: token,
		Id:          review.Owner.Id,
	}
	res, err := s.userService.GetStudentProfile(ctx, reqStudent)
	if err != nil {
		return nil, err
	}
	review.Owner.Name = res.Student.Name
	return review, nil
}

func (s *reviewService) GetReviewsByCompany(ctx context.Context, token string, companyID int64) ([]*pbv1.ReviewCompany, error) {
	_, err := utils.ValidateAccessToken(token)
	if err != nil {
		return nil, domain.ErrUnauthorize
	}

	reviews, err := s.repo.GetReviewsByCompany(ctx, companyID)
	if err != nil {
		return nil, err
	}

	m := make(map[int64]string) // map[userID]Name

	for _, review := range reviews {
		if review.Owner.Id == 0 {
			continue
		}
		m[review.Owner.Id] = ""
	}

	keys := make([]int64, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	res, err := s.userService.GetStudentProfiles(ctx, &pbv1.GetStudentsRequest{
		AccessToken: token,
		Ids:         keys,
	})
	if err != nil {
		return nil, err
	}

	for _, student := range res.Students {
		m[student.Id] = student.Name
	}

	for _, review := range reviews {
		if review.Owner.Id == 0 {
			review.Owner.Name = AnonymousName
			continue
		}
		review.Owner.Name = m[review.Owner.Id]
	}

	return reviews, nil
}

func (s *reviewService) UpdateReview(ctx context.Context, token string, review *pbv1.UpdatedReview, rid int64) error {
	if !domain.CheckUpdatedRequireField(review) {
		return domain.ErrFieldsAreRequired
	}
	if !domain.CheckRatingRange(review.Rating) {
		return domain.ErrRatingRange
	}

	payload, err := utils.ValidateAccessToken(token)
	if err != nil {
		return domain.ErrUnauthorize
	}

	if payload.Role != StudentRole {
		return domain.ErrForbidden
	}

	// Check review owner
	ownerID := payload.UserId
	uid, err := s.repo.GetReviewOwner(ctx, rid)
	if err == sql.ErrNoRows && ownerID != uid {
		return domain.ErrForbidden
	}
	if err != nil {
		return err
	}

	err = s.repo.UpdateReview(ctx, review, rid)
	if err != nil {
		return err
	}

	return nil
}

func (s *reviewService) GetReviewsByUser(ctx context.Context, token string, userID int64) ([]*pbv1.MyReview, error) {
	// Similar to GetReviewsByCompany function

	payload, err := utils.ValidateAccessToken(token)
	if err != nil {
		return nil, domain.ErrUnauthorize
	}
	if payload.Role != StudentRole {
		return nil, domain.ErrForbidden
	}

	myReviews, err := s.repo.GetReviewsByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	// fetch company name
	for _, review := range myReviews {
		req := &pbv1.GetCompanyRequest{
			AccessToken: token,
			Id:          review.Company.Id,
		}

		cProfile, err := s.userService.GetCompanyProfile(ctx, req)
		if err != nil {
			return nil, err
		}
		review.Company.Name = cProfile.Company.Name
	}
	return myReviews, nil
}

func (s *reviewService) DeleteReview(ctx context.Context, token string, reviewID int64) error {
	// Similar to UpdateReview function

	// Check request user
	payload, err := utils.ValidateAccessToken(token)
	if err != nil {
		return domain.ErrUnauthorize
	}
	if payload.Role != StudentRole {
		return domain.ErrForbidden
	}

	// Check review owner
	ownerID := payload.UserId
	uid, err := s.repo.GetReviewOwner(ctx, reviewID)
	if err == sql.ErrNoRows && ownerID != uid {
		return domain.ErrForbidden
	}
	if err != nil {
		return err
	}

	// Delete review
	err = s.repo.DeleteReview(ctx, reviewID)
	if err != nil {
		return err
	}

	return nil
}
