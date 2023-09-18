package service

import (
	"JinnnDamanee/review-service/internal/domain"
	"JinnnDamanee/review-service/internal/model"
	"JinnnDamanee/review-service/internal/repo"
)

type ReviewService struct {
	Repo *repo.ReviewRepository
}

func NewReviewService(repo *repo.ReviewRepository) *ReviewService {
	return &ReviewService{Repo: repo}
}

func (s *ReviewService) CreateReview(review *model.Review) (*model.Review, error) {
	r, err := s.Repo.Create(review)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *ReviewService) GetAllReviews() ([]*model.Review, error) {
	reviews, err := s.Repo.FindAll()
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (s *ReviewService) GetReviewByID(id int) (*model.Review, error) {
	review, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, domain.ErrReviewIDNotFound
	}
	return review, nil
}

func (s *ReviewService) UpdateReviewByID(id int, newReview *model.Review) (*model.Review, error) {
	review, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, domain.ErrReviewIDNotFound
	}

	review.ReviewerID = newReview.ReviewerID
	review.ReviewTitle = newReview.ReviewTitle
	review.ReviewDetail = newReview.ReviewDetail

	r, err := s.Repo.Update(review)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (s *ReviewService) DeleteReviewByID(id int) error {
	r, err := s.Repo.FindByID(id)
	if err != nil {
		return domain.ErrReviewIDNotFound
	}

	if err = s.Repo.Delete(r); err != nil {
		return err
	}
	return nil
}
