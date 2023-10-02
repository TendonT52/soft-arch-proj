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
	PostRepo    port.PostRepoPort
	UserService port.UserClientPort
}

func NewPostService(postRepo port.PostRepoPort, userService port.UserClientPort) port.PostServicePort {
	return &postService{
		PostRepo:    postRepo,
		UserService: userService,
	}
}

func (s *postService) CreatePost(ctx context.Context, token string, post *pbv1.CreatedPost) (int64, error) {
	if !domain.CheckRequireFields(post) {
		return 0, domain.ErrFieldsAreRequired
	}

	payload, err := utils.ValidateAccessToken(token)
	if err != nil {
		return 0, domain.ErrUnauthorize
	}
	if payload.Role != companyRole {
		return 0, domain.ErrForbidden
	}

	postId, err := s.PostRepo.CreatePost(ctx, payload.UserId, post)
	if err != nil {
		return 0, err
	}

	return postId, nil
}

func (s *postService) GetPost(ctx context.Context, token string, postId int64) (*pbv1.Post, error) {
	_, err := utils.ValidateAccessToken(token)
	if err != nil {
		return nil, domain.ErrUnauthorize
	}
	post, err := s.PostRepo.GetPost(ctx, postId)
	if err != nil {
		return nil, err
	}

	req := &pbv1.GetCompanyRequest{
		AccessToken: token,
		Id:          post.Owner.Id,
	}
	res, err := s.UserService.GetCompanyProfile(ctx, req)
	if err != nil {
		return nil, err
	}
	post.Owner.Name = res.Company.Name

	return post, nil
}

func (s *postService) GetPosts(ctx context.Context, token string, search *pbv1.SearchOptions) ([]*pbv1.Post, error) {
	_, err := utils.ValidateAccessToken(token)
	if err != nil {
		return nil, domain.ErrUnauthorize
	}
	if search == nil {
		search = &pbv1.SearchOptions{}
	}

	u, err := s.UserService.ListApprovedCompanies(ctx, &pbv1.ListApprovedCompaniesRequest{
		AccessToken: token,
		Search:      search.SearchCompany,
	})
	if err != nil {
		return nil, err
	}

	companyInfo := domain.NewCompanyInfo(u.Companies)
	search = domain.RemoveSpecialChars(search)
	posts, err := s.PostRepo.GetPosts(ctx, search, companyInfo)
	if posts == nil {
		return nil, domain.ErrPostNotFound
	}
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *postService) UpdatePost(ctx context.Context, token string, postId int64, post *pbv1.UpdatedPost) error {
	if !domain.CheckUpdatedFields(post) {
		return domain.ErrFieldsAreRequired
	}

	owner, err := s.PostRepo.GetOwner(ctx, postId)
	if err != nil {
		return err
	}
	payload, err := utils.ValidateAccessToken(token)
	if err != nil {
		return domain.ErrUnauthorize
	}
	if payload.Role != companyRole || payload.UserId != owner {
		return domain.ErrForbidden
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
	payload, err := utils.ValidateAccessToken(token)
	if err != nil {
		return domain.ErrUnauthorize
	}
	if payload.Role != companyRole || payload.UserId != owner {
		return domain.ErrForbidden
	}

	err = s.PostRepo.DeletePost(ctx, postId)
	if err != nil {
		return err
	}

	return nil
}

func (s *postService) GetOpenPositions(ctx context.Context, token, search string) ([]string, error) {
	payload, err := utils.ValidateAccessToken(token)
	if err != nil {
		return nil, domain.ErrUnauthorize
	}
	if payload.Role != companyRole {
		return nil, domain.ErrForbidden
	}

	openPositions, err := s.PostRepo.GetOpenPositions(ctx, search)
	if err != nil {
		return nil, err
	}

	return openPositions, nil
}

func (s *postService) GetRequiredSkills(ctx context.Context, token, search string) ([]string, error) {
	payload, err := utils.ValidateAccessToken(token)
	if err != nil {
		return nil, domain.ErrUnauthorize
	}
	if payload.Role != companyRole {
		return nil, domain.ErrForbidden
	}

	requiredSkills, err := s.PostRepo.GetRequiredSkills(ctx, search)
	if err != nil {
		return nil, err
	}

	return requiredSkills, nil
}

func (s *postService) GetBenefits(ctx context.Context, token, search string) ([]string, error) {
	payload, err := utils.ValidateAccessToken(token)
	if err != nil {
		return nil, domain.ErrUnauthorize
	}
	if payload.Role != companyRole {
		return nil, domain.ErrForbidden
	}

	benefits, err := s.PostRepo.GetBenefits(ctx, search)
	if err != nil {
		return nil, err
	}

	return benefits, nil
}
