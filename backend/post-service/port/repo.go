package port

import (
	"context"

	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
)

type PostRepoPort interface {
	CreatePost(ctx context.Context, userId int64, post *pbv1.CreatedPost) (int64, error)
	GetPost(ctx context.Context, postId int64) (*pbv1.Post, error)
	GetPosts(ctx context.Context, search *pbv1.SearchOptions, cids *domain.CompanyInfo) ([]*pbv1.Post, error)
	GetMyPosts(ctx context.Context, userId int64) ([]*pbv1.Post, error)
	GetOwner(ctx context.Context, postId int64) (int64, error)
	UpdatePost(ctx context.Context, postId int64, post *pbv1.UpdatedPost) error
	DeletePost(ctx context.Context, postId int64) error
	GetOpenPositions(ctx context.Context, search string) ([]string, error)
	GetRequiredSkills(ctx context.Context, search string) ([]string, error)
	GetBenefits(ctx context.Context, search string) ([]string, error)
}
