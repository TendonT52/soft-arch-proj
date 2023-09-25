package service

import (
	"context"

	pbUser "github.com/TikhampornSky/go-auth-verifiedMail/gen/v1"
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

func (s *postService) CreatePost(ctx context.Context, token string, post *pbv1.Post) (int64, error) {
	if !domain.CheckRequireFields(post) {
		return 0, domain.ErrFieldsAreRequired
	}

	payload, err := utils.ValidateAccessToken(token)
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

func (s *postService) GetPost(ctx context.Context, token string, postId int64) (*pbv1.Post, error) {
	_, err := utils.ValidateAccessToken(token)
	if err != nil {
		return nil, err
	}
	post, err := s.PostRepo.GetPost(ctx, postId)
	if err != nil {
		return nil, err
	}

	req := &pbUser.GetCompanyRequest{
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
		return nil, err
	}

	u, err := s.UserService.ListApprovedCompanies(ctx, &pbUser.ListApprovedCompaniesRequest{
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
	payload, err := utils.ValidateAccessToken(token)
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

func (s *postService) DeleteAllPosts(ctx context.Context, token string) error {
	_, err := utils.ValidateAccessToken(token)
	if err != nil {
		return err
	}

	err = s.PostRepo.DeleteAllPosts(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *postService) GetOpenPositions(ctx context.Context, token, search string) ([]string, error) {
	payload, err := utils.ValidateAccessToken(token)
	if payload.Role != companyRole {
		return nil, domain.ErrUnauthorized
	}
	if err != nil {
		return nil, err
	}

	openPositions, err := s.PostRepo.GetOpenPositions(ctx, search)
	if err != nil {
		return nil, err
	}

	return openPositions, nil
}

func (s *postService) GetRequiredSkills(ctx context.Context, token, search string) ([]string, error) {
	payload, err := utils.ValidateAccessToken(token)
	if payload.Role != companyRole {
		return nil, domain.ErrUnauthorized
	}
	if err != nil {
		return nil, err
	}

	requiredSkills, err := s.PostRepo.GetRequiredSkills(ctx, search)
	if err != nil {
		return nil, err
	}

	return requiredSkills, nil
}

func (s *postService) GetBenefits(ctx context.Context, token, search string) ([]string, error) {
	payload, err := utils.ValidateAccessToken(token)
	if payload.Role != companyRole {
		return nil, domain.ErrUnauthorized
	}
	if err != nil {
		return nil, err
	}

	benefits, err := s.PostRepo.GetBenefits(ctx, search)
	if err != nil {
		return nil, err
	}

	return benefits, nil
}
