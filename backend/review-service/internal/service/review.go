package service

import (
	"fmt"
	"jindamanee2544/review-service/internal/model"
	"jindamanee2544/review-service/internal/repo"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReviewService struct {
	Repo *repo.ReviewRepository
}

func NewReviewService(repo *repo.ReviewRepository) *ReviewService {
	return &ReviewService{Repo: repo}
}

func (s *ReviewService) CreateReview(c echo.Context) error {
	req := model.ReviewCreateRequest{}
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	review := model.Review{
		ReviewerID:   req.ReviewerID,
		ReviewTitle:  req.ReviewTitle,
		ReviewDetail: req.ReviewDetail,
	}

	r, err := s.Repo.Create(&review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	c.JSON(http.StatusOK, r)
	return nil
}

func (s *ReviewService) GetAllReviews(c echo.Context) error {
	reviews, err := s.Repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	c.JSON(http.StatusOK, reviews)
	return nil
}

func (s *ReviewService) GetReviewByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	review, err := s.Repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Review with id %d not found", id))
		return err
	}
	c.JSON(http.StatusOK, review)
	return nil
}

func (s *ReviewService) UpdateReviewByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}
	req := model.ReviewUpdateRequest{}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}

	_, err = s.Repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Review with id %d not found", id))
		return err
	}

	r := model.Review{
		ReviewerID:   req.ReviewerID,
		ReviewTitle:  req.ReviewTitle,
		ReviewDetail: req.ReviewDetail,
	}
	r.ID = uint(id)
	updatedReview, err := s.Repo.Update(&r)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	c.JSON(http.StatusOK, updatedReview)
	return nil
}

func (s *ReviewService) DeleteReviewByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return err
	}

	r, err := s.Repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, fmt.Sprintf("Review with id %d not found", id))
		return err
	}

	if err = s.Repo.Delete(r); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	c.JSON(http.StatusOK, "Deleted")
	return nil
}
