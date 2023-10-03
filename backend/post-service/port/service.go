package port

import (
	"context"

	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
)

type PostServicePort interface {
	CreatePost(ctx context.Context, token string, post *pbv1.Post) (int64, error)
	GetPost(ctx context.Context, token string, postId int64) (*pbv1.Post, error)
	GetPosts(ctx context.Context, token string, search *pbv1.SearchOptions) ([]*pbv1.Post, error)
	UpdatePost(ctx context.Context, token string, postId int64, post *pbv1.UpdatedPost) error
	DeletePost(ctx context.Context, token string, postId int64) error
	GetOpenPositions(ctx context.Context, token, search string) ([]string, error)
	GetRequiredSkills(ctx context.Context, token, search string) ([]string, error)
	GetBenefits(ctx context.Context, token, search string) ([]string, error)
}

type UserClientPort interface {
	GetCompanyProfile(ctx context.Context, req *pbv1.GetCompanyRequest) (*pbv1.GetCompanyResponse, error)
	ListApprovedCompanies(ctx context.Context, req *pbv1.ListApprovedCompaniesRequest) (*pbv1.ListApprovedCompaniesResponse, error)
}
