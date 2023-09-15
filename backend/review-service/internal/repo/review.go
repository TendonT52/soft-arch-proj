package repo

import (
	"jindamanee2544/review-service/internal/model"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	DB *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	db.AutoMigrate(&model.Review{})
	return &ReviewRepository{DB: db}
}

func (r *ReviewRepository) Create(review *model.Review) (*model.Review, error) {
	if err := r.DB.Create(review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func (r *ReviewRepository) FindAll() ([]*model.Review, error) {
	reviews := []*model.Review{}
	if err := r.DB.Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *ReviewRepository) FindByID(id int) (*model.Review, error) {
	review := &model.Review{}
	if err := r.DB.First(review, id).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func (r *ReviewRepository) Update(review *model.Review) (*model.Review, error) {
	if err := r.DB.Save(&review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func (r *ReviewRepository) Delete(review *model.Review) error {
	if err := r.DB.Unscoped().Delete(&review).Error; err != nil {
		return err
	}
	return nil
}
