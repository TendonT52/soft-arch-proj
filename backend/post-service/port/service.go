package port

import (
	"context"

	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
)

type PostServicePort interface {
	CreatePost(ctx context.Context, token string, post *pbv1.Post) (int64, error)
	GetPost(ctx context.Context, token string, postId int64) (*pbv1.Post, error)
	GetPosts(ctx context.Context, token string, search string) ([]*pbv1.Post, error)
	UpdatePost(ctx context.Context, token string, postId int64, post *pbv1.Post) error
	DeletePost(ctx context.Context, token string, postId int64) error
}

type TokenServicePort interface {
	ValidateAccessToken(token string) (*domain.Payload, error)
}