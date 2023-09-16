package port

import (
	"context"

	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
)

type PostRepoPort interface {
	CreatePost(ctx context.Context, userId int64, post *pbv1.Post) (int64, error)
	GetPost(ctx context.Context, postId int64) (*pbv1.Post, error)
	GetPosts(ctx context.Context, search string) ([]*pbv1.Post, error)
	GetOwner(ctx context.Context, postId int64) (int64, error)
	UpdatePost(ctx context.Context, postId int64, post *pbv1.Post) error
	DeletePost(ctx context.Context, postId int64) error
}
