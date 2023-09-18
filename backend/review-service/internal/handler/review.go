package handler

import (
	"JinnnDamanee/review-service/internal/domain"
	"JinnnDamanee/review-service/internal/model"
	"JinnnDamanee/review-service/internal/service"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReviewHandler struct {
	Service *service.ReviewService
}

func NewReviewHandler(service *service.ReviewService) *ReviewHandler {
	return &ReviewHandler{Service: service}
}

func (h *ReviewHandler) GetAllReview(c echo.Context) error {
	reviews, err := h.Service.GetAllReviews()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.RespInternal)
	}
	return c.JSON(http.StatusOK, domain.ResponseWithData{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    reviews,
	})
}

func (h *ReviewHandler) GetReviewByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BadRequest)
	}

	review, err := h.Service.GetReviewByID(id)
	if err != nil {
		if errors.Is(err, domain.ErrReviewIDNotFound) {
			return c.JSON(http.StatusNotFound, domain.Response{
				Code:    http.StatusNotFound,
				Message: domain.ErrReviewIDNotFound.Message,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, domain.Response{
				Code:    http.StatusInternalServerError,
				Message: domain.ErrInternal.Message,
			})
		}
	}
	return c.JSON(http.StatusOK, domain.ResponseWithData{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    review,
	})
}

func (h *ReviewHandler) CreateReview(c echo.Context) error {
	req := model.ReviewCreateRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, domain.BadRequest)
	}

	review := model.Review{
		ReviewerID:   req.ReviewerID,
		ReviewTitle:  req.ReviewTitle,
		ReviewDetail: req.ReviewDetail,
	}
	r, err := h.Service.CreateReview(&review)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.RespInternal)
	}
	return c.JSON(http.StatusOK, domain.ResponseWithData{
		Code:    http.StatusOK,
		Message: "Create successful",
		Data:    r,
	})
}

func (h *ReviewHandler) UpdateReviewByID(c echo.Context) error {
	reviewId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BadRequest)

	}
	req := model.ReviewUpdateRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, domain.BadRequest)

	}

	r := model.Review{
		ReviewerID:   req.ReviewerID,
		ReviewTitle:  req.ReviewTitle,
		ReviewDetail: req.ReviewDetail,
	}

	updatedReview, err := h.Service.UpdateReviewByID(reviewId, &r)
	if err != nil {
		if errors.Is(err, domain.ErrReviewIDNotFound) {
			return c.JSON(http.StatusNotFound, domain.Response{
				Code:    http.StatusNotFound,
				Message: domain.ErrReviewIDNotFound.Message,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, domain.Response{
				Code:    http.StatusInternalServerError,
				Message: domain.ErrInternal.Message,
			})
		}

	}
	return c.JSON(http.StatusOK, domain.ResponseWithData{
		Code:    http.StatusOK,
		Message: "Update successfully",
		Data:    updatedReview,
	})
}

func (h *ReviewHandler) DeleteReviewByID(c echo.Context) error {
	d, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.BadRequest)
	}

	if err := h.Service.DeleteReviewByID(d); err != nil {
		if errors.Is(err, domain.ErrReviewIDNotFound) {
			return c.JSON(http.StatusNotFound, domain.Response{
				Code:    http.StatusNotFound,
				Message: domain.ErrReviewIDNotFound.Message,
			})
		} else {
			return c.JSON(http.StatusInternalServerError, domain.Response{
				Code:    http.StatusInternalServerError,
				Message: domain.ErrInternal.Message,
			})
		}

	}
	return c.JSON(http.StatusOK, domain.Response{
		Code:    http.StatusOK,
		Message: "Delete successfully",
	})
}
