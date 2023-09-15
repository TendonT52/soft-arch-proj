package port

import (
	"context"

	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
)

type PostServicePort interface {
	CreatePost(ctx context.Context, userId int64, post *pbv1.Post) (int64, error)
	GetPost(ctx context.Context, userId, postId int64) (*pbv1.Post, error)
	GetPosts(ctx context.Context, userId int64, search string) ([]*pbv1.Post, error)
	UpdatePost(ctx context.Context, userId, postId int64, post *pbv1.Post) error
	DeletePost(ctx context.Context, userId, postId int64) error
}
