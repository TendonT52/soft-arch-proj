package domain

import (
	pbv1 "github.com/JinnnDamanee/review-service/gen/v1"
)

func CheckCreatedRequireField(review *pbv1.CreatedReview) bool {
	if review.Rating == 0 || review.Title == "" || review.Description == "" {
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

func CheckRatingRange(rating int32) bool {
	if rating < 1 || rating > 5 {
		return false
	}
	return true
}