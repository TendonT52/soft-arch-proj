package service

import (
	"context"

	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
	"github.com/TikhampornSky/go-post-service/port"
	"github.com/TikhampornSky/go-post-service/utils"
)

const companyRole = "company"

type postService struct {
	PostRepo port.PostRepoPort
}

func NewPostService(postRepo port.PostRepoPort) port.PostServicePort {
	return &postService{PostRepo: postRepo}
}

func (s *postService) CreatePost(ctx context.Context, token string, post *pbv1.Post) (int64, error) {
	if !domain.CheckRequireFields(post) {
		return 0, domain.ErrFieldsAreRequired
	}

	p := domain.NewPost(post)
	payload, err := utils.ValidateAccessToken(token)
	if payload.Role != companyRole {
		return 0, domain.ErrUnauthorized
	}
	
	postId, err := s.PostRepo.CreatePost(ctx, payload.UserId, p)
	if err != nil {
		return 0, err
	}

	return postId, nil
}

func (s *postService) GetPost(ctx context.Context, token string, postId int64) (*pbv1.Post, error) {
	return nil, nil
}

func (s *postService) GetPosts(ctx context.Context, token string, search string) ([]*pbv1.Post, error) {
	return nil, nil
}

func (s *postService) UpdatePost(ctx context.Context, token string, postId int64, post *pbv1.Post) error {
	return nil
}

func (s *postService) DeletePost(ctx context.Context, token string, postId int64) error {
	return nil
}
