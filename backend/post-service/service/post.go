package service

import (
	"context"

	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
	"github.com/TikhampornSky/go-post-service/port"
)

const companyRole = "company"

type postService struct {
	PostRepo     port.PostRepoPort
	TokenService port.TokenServicePort
}

func NewPostService(postRepo port.PostRepoPort, tokenService port.TokenServicePort) port.PostServicePort {
	return &postService{
		PostRepo:     postRepo,
		TokenService: tokenService,
	}
}

func (s *postService) CreatePost(ctx context.Context, token string, post *pbv1.Post) (int64, error) {
	if !domain.CheckRequireFields(post) {
		return 0, domain.ErrFieldsAreRequired
	}

	payload, err := s.TokenService.ValidateAccessToken(token)
	if err != nil {
		return 0, err
	}
	if payload.Role != companyRole {
		return 0, domain.ErrUnauthorized
	}

	postId, err := s.PostRepo.CreatePost(ctx, payload.UserId, post)
	if err != nil {
		return 0, err
	}

	return postId, nil
}

func (s *postService) GetPost(ctx context.Context, token string, postId int64) (*pbv1.Post, error) { // Internal use only
	_, err := s.TokenService.ValidateAccessToken(token)
	if err != nil {
		return nil, err
	}
	post, err := s.PostRepo.GetPost(ctx, postId)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *postService) GetPosts(ctx context.Context, token string, search string) ([]*pbv1.Post, error) {
	return nil, nil
}

func (s *postService) UpdatePost(ctx context.Context, token string, postId int64, post *pbv1.Post) error {
	if !domain.CheckRequireFields(post) {
		return domain.ErrFieldsAreRequired
	}

	owner, err := s.PostRepo.GetOwner(ctx, postId)
	if err != nil {
		return err
	}
	payload, err := s.TokenService.ValidateAccessToken(token)
	if err != nil {
		return err
	}
	if payload.Role != companyRole || payload.UserId != owner {
		return domain.ErrUnauthorized
	}

	err = s.PostRepo.UpdatePost(ctx, postId, post)
	if err != nil {
		return err
	}
	return nil
}

func (s *postService) DeletePost(ctx context.Context, token string, postId int64) error {
	owner, err := s.PostRepo.GetOwner(ctx, postId)
	if err != nil {
		return err
	}
	payload, err := s.TokenService.ValidateAccessToken(token)
	if err != nil {
		return err
	}
	if payload.Role != companyRole || payload.UserId != owner {
		return domain.ErrUnauthorized
	}

	err = s.PostRepo.DeletePost(ctx, postId)
	if err != nil {
		return err
	}

	return nil
}
