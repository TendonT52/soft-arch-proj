package model

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ReviewerID   int    `json:"reviewer_id"`
	ReviewTitle  string `json:"review_title"`
	ReviewDetail string `json:"review_detail"`
}

type ReviewCreateRequest struct {
	ReviewerID   int    `json:"reviewer_id"`
	ReviewTitle  string `json:"review_title"`
	ReviewDetail string `json:"review_detail"`
}

type ReviewUpdateRequest struct {
	ReviewerID   int    `json:"reviewer_id"`
	ReviewTitle  string `json:"review_title"`
	ReviewDetail string `json:"review_detail"`
}
