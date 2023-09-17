package port

import (
	"context"

	pbUser "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
)

type PostServicePort interface {
	CreatePost(ctx context.Context, token string, post *pbv1.Post) (int64, error)
	GetPost(ctx context.Context, token string, postId int64) (*pbv1.Post, error)
	GetPosts(ctx context.Context, token string, search *pbv1.SearchOptions) ([]*pbv1.Post, error)
	UpdatePost(ctx context.Context, token string, postId int64, post *pbv1.UpdatedPost) error
	DeletePost(ctx context.Context, token string, postId int64) error
	DeleteAllPosts(ctx context.Context, token string) error // for testing
}

type TokenServicePort interface {
	ValidateAccessToken(token string) (*domain.Payload, error)
}

type UserClientPort interface {
	GetCompanyProfile(ctx context.Context, req *pbUser.GetCompanyRequest) (*pbUser.GetCompanyResponse, error)
	ListApprovedCompanies(ctx context.Context, req *pbUser.ListApprovedCompaniesRequest) (*pbUser.ListApprovedCompaniesResponse, error)
}
