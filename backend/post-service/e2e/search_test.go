package e2e

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/TikhampornSky/go-post-service/config"
	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
	"github.com/TikhampornSky/go-post-service/mock"
	"github.com/TikhampornSky/go-post-service/tools"
	"github.com/TikhampornSky/go-post-service/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	lex = `{"root": {}}`
)

func createMockPost(t *testing.T, ctx context.Context, c pbv1.PostServiceClient, ownerId int64, companyName, topic, description, period, howTo string, ops, rss, bs []string) *pbv1.Post {
	p := &pbv1.CreatedPost{
		Topic:          topic,
		Description:    description,
		Period:         period,
		HowTo:          howTo,
		OpenPositions:  ops,
		RequiredSkills: rss,
		Benefits:       bs,
	}

	config, _ := config.LoadConfig("..")
	token, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: ownerId,
		Role:   "company",
	})
	require.NoError(t, err)

	res, err := c.CreatePost(ctx, &pbv1.CreatePostRequest{
		AccessToken: token,
		Post:        p,
	})
	require.NoError(t, err)
	require.Equal(t, int64(201), res.Status)

	return &pbv1.Post{
		PostId:         res.Id,
		Topic:          p.Topic,
		Description:    p.Description,
		Period:         p.Period,
		HowTo:          p.HowTo,
		OpenPositions:  p.OpenPositions,
		RequiredSkills: p.RequiredSkills,
		Benefits:       p.Benefits,
		Owner: &pbv1.PostOwner{
			Id:   ownerId,
			Name: companyName,
		},
		UpdatedAt: time.Now().Unix(),
	}
}

func TestSearchPosts(t *testing.T) {
	config, _ := config.LoadConfig("..")
	target := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewPostServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ad, err := mock.CreateMockAdmin(ctx)
	userRes, _, err := mock.CreateMockApprovedCompany(ctx, "Some cute company", ad)
	require.NoError(t, err)
	token, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: userRes.Id,
		Role:   "company",
	})
	require.NoError(t, err)

	err = tools.DeleteAllPosts()
	require.NoError(t, err)

	// Create Mock companies
	com1Res, _, err := mock.CreateMockApprovedCompany(ctx, "Agoda", ad)
	require.NoError(t, err)
	com2Res, _, err := mock.CreateMockApprovedCompany(ctx, "Dime", ad)
	require.NoError(t, err)
	com3Res, _, err := mock.CreateMockApprovedCompany(ctx, "Grab", ad)
	require.NoError(t, err)

	post1 := createMockPost(t, ctx, c, com1Res.Id, "Agoda", "Post 1", lex, "1 month", "Apply via our facebook page",
		[]string{"Software Engineer", "Data Scientist"}, []string{"Golang", "Python"}, []string{"Free lunch", "Free dinner"})
	post2 := createMockPost(t, ctx, c, com2Res.Id, "Dime", "Post 2", lex, "2 month", "Apply via our line VOOM page",
		[]string{"Data Analysts", "Full-Stack Developer"}, []string{"Python", "HTML", "CSS"}, []string{"Free lunch", "Macbook Pro"})
	post3 := createMockPost(t, ctx, c, com2Res.Id, "Dime", "Post 3", lex, "3 month", "Apply via our TikTok",
		[]string{"Backend Developer", "Data Scientist"}, []string{"Golang", "Python"}, []string{"Free dinner", "Macbook M1"})
	post4 := createMockPost(t, ctx, c, com2Res.Id, "Dime", "Post 4", lex, "4 month", "Apply via our facebook page",
		[]string{"Frontend Developer", "Data Analyst"}, []string{"HTML", "CSS", "Javascript"}, []string{"Free lunch", "Free dinner", "Macbook M1"})
	post5 := createMockPost(t, ctx, c, com3Res.Id, "Grab", "Post 5", lex, "5 month", "Apply via our facebook page",
		[]string{"Frontend Developer", "Data Analyst"}, []string{"HTML", "CSS", "Javascript"}, []string{"Free lunch", "Free dinner", "Macbook Pro"})

	_ = post1
	_ = post2
	_ = post3
	_ = post4
	_ = post5

	tests := map[string]struct {
		req    *pbv1.ListPostsRequest
		expect *pbv1.ListPostsResponse
	}{
		"success": {
			req: &pbv1.ListPostsRequest{
				AccessToken: token,
				SearchOptions: &pbv1.SearchOptions{
					SearchCompany:       "Dime !!!",
					SearchOpenPosition:  "Developer",
					SearchRequiredSkill: "HTML CSS",
					SearchBenefit:       "Macbook M2",
				},
			},
			expect: &pbv1.ListPostsResponse{
				Status:  200,
				Message: "Posts retrieved successfully",
				Posts: []*pbv1.Post{
					post4,
					post2,
				},
			},
		},
		"success with empty search options": {
			req: &pbv1.ListPostsRequest{
				AccessToken: token,
				SearchOptions: &pbv1.SearchOptions{
					SearchCompany:       "Dime",
					SearchOpenPosition:  "",
					SearchRequiredSkill: "",
					SearchBenefit:       "",
				},
			},
			expect: &pbv1.ListPostsResponse{
				Status:  200,
				Message: "Posts retrieved successfully",
				Posts: []*pbv1.Post{
					post2,
					post3,
					post4,
				},
			},
		},
		"success with all empty search options": {
			req: &pbv1.ListPostsRequest{
				AccessToken: token,
				SearchOptions: &pbv1.SearchOptions{
					SearchCompany:       "",
					SearchOpenPosition:  "",
					SearchRequiredSkill: "",
					SearchBenefit:       "",
				},
			},
			expect: &pbv1.ListPostsResponse{
				Status:  200,
				Message: "Posts retrieved successfully",
				Posts: []*pbv1.Post{
					post1,
					post2,
					post3,
					post4,
					post5,
				},
			},
		},
		"success with NIL empty search options": {
			req: &pbv1.ListPostsRequest{
				AccessToken: token,
			},
			expect: &pbv1.ListPostsResponse{
				Status:  200,
				Message: "Posts retrieved successfully",
				Posts: []*pbv1.Post{
					post1,
					post2,
					post3,
					post4,
					post5,
				},
			},
		},
		"success with only require_skills search options": {
			req: &pbv1.ListPostsRequest{
				AccessToken: token,
				SearchOptions: &pbv1.SearchOptions{
					SearchCompany:       "Dime",
					SearchOpenPosition:  "",
					SearchRequiredSkill: "Golang Javascript",
					SearchBenefit:       "",
				},
			},
			expect: &pbv1.ListPostsResponse{
				Status:  200,
				Message: "Posts retrieved successfully",
				Posts: []*pbv1.Post{
					post3,
					post4,
				},
			},
		},
		"success with only open_positions search options": {
			req: &pbv1.ListPostsRequest{
				AccessToken: token,
				SearchOptions: &pbv1.SearchOptions{
					SearchCompany:       "Dime",
					SearchOpenPosition:  "Data Analyst",
					SearchRequiredSkill: "",
					SearchBenefit:       "",
				},
			},
			expect: &pbv1.ListPostsResponse{
				Status:  200,
				Message: "Posts retrieved successfully",
				Posts: []*pbv1.Post{
					post4,
					post2,
					post3,
				},
			},
		},
		"success with only benefits search options": {
			req: &pbv1.ListPostsRequest{
				AccessToken: token,
				SearchOptions: &pbv1.SearchOptions{
					SearchCompany:       "Dime",
					SearchOpenPosition:  "",
					SearchRequiredSkill: "",
					SearchBenefit:       "Macbook Pro",
				},
			},
			expect: &pbv1.ListPostsResponse{
				Status:  200,
				Message: "Posts retrieved successfully",
				Posts: []*pbv1.Post{
					post2,
					post3,
					post4,
				},
			},
		},

		"not found": {
			req: &pbv1.ListPostsRequest{
				AccessToken: token,
				SearchOptions: &pbv1.SearchOptions{
					SearchCompany:       "Dime",
					SearchOpenPosition:  "mock-search-open-position",
					SearchRequiredSkill: "mock-search-required-skill",
					SearchBenefit:       "mock-search-benefit",
				},
			},
			expect: &pbv1.ListPostsResponse{
				Status:  404,
				Message: "Posts not found",
			},
		},
		"token wrong": {
			req: &pbv1.ListPostsRequest{
				AccessToken:   "wrong",
				SearchOptions: &pbv1.SearchOptions{},
			},
			expect: &pbv1.ListPostsResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.ListPosts(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			require.Equal(t, int64(len(tc.expect.Posts)), res.Total)
			require.Equal(t, len(tc.expect.Posts), len(res.Posts))
			for i, p := range res.Posts {
				require.Equal(t, tc.expect.Posts[i].Topic, p.Topic)
				require.Equal(t, tc.expect.Posts[i].Description, p.Description)
				require.Equal(t, tc.expect.Posts[i].Period, p.Period)
				require.Equal(t, tc.expect.Posts[i].HowTo, p.HowTo)
				require.Equal(t, true, utils.CheckArrayEqual(&tc.expect.Posts[i].OpenPositions, &p.OpenPositions))
				require.Equal(t, true, utils.CheckArrayEqual(&tc.expect.Posts[i].RequiredSkills, &p.RequiredSkills))
				require.Equal(t, true, utils.CheckArrayEqual(&tc.expect.Posts[i].Benefits, &p.Benefits))
				require.NotEmpty(t, p.Owner.Id)
				require.Equal(t, tc.expect.Posts[i].Owner.Name, p.Owner.Name)
				require.NotEmpty(t, p.UpdatedAt)
			}
		})
	}
}
