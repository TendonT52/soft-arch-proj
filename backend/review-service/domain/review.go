package domain

import (
	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
)

func CheckCreatedRequireField(review *pbv1.CreatedReview) bool {
	if review.Cid <= 0 || review.Rating == 0 || review.Title == "" || review.Description == "" {
		return false
	}
	return true
}

func CheckUpdatedRequireField(review *pbv1.UpdatedReview) bool {
	if review.Rating == 0 || review.Title == "" || review.Description == "" {
		return false
	}
	return true
}