package e2e

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/TikhampornSky/go-post-service/config"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
	"github.com/TikhampornSky/go-post-service/mock"
	"github.com/TikhampornSky/go-post-service/tools"
	"github.com/TikhampornSky/go-post-service/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetMyPost(t *testing.T) {
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

	lex := `{
		"root": {}
	  }
	`

	// Create Mock admin
	ad, err := mock.CreateMockAdmin(ctx)
	require.NoError(t, err)

	err = tools.DeleteAllPosts()
	require.NoError(t, err)

	// Create Mock companies
	myCompany, myCompanyToken, err := mock.CreateMockApprovedCompany(ctx, "Dime", ad)
	require.NoError(t, err)
	otherCompany, _, err := mock.CreateMockApprovedCompany(ctx, "Agoda", ad)
	require.NoError(t, err)

	post1 := createMockPost(t, ctx, c, otherCompany.Id, "Agoda", "Post 1", lex, "1 month", "Apply via our facebook page",
		[]string{"Software Engineer", "Data Scientist"}, []string{"Golang", "Python"}, []string{"Free lunch", "Free dinner"})
	post2 := createMockPost(t, ctx, c, myCompany.Id, "Dime", "Post 2", lex, "2 month", "Apply via our line VOOM page",
		[]string{"Data Analysts", "Full-Stack Developer"}, []string{"Python", "HTML", "CSS"}, []string{"Free lunch", "Macbook Pro"})
	post3 := createMockPost(t, ctx, c, myCompany.Id, "Dime", "Post 3", lex, "3 month", "Apply via our TikTok",
		[]string{"Backend Developer", "Data Scientist"}, []string{"Golang", "Python"}, []string{"Free dinner", "Macbook M1"})
	post4 := createMockPost(t, ctx, c, myCompany.Id, "Dime", "Post 4", lex, "4 month", "Apply via our facebook page",
		[]string{"Frontend Developer", "Data Analyst"}, []string{"HTML", "CSS", "Javascript"}, []string{"Free lunch", "Free dinner", "Macbook M1"})

	_ = post1

	tests := map[string]struct {
		req    *pbv1.GetMyPostsRequest
		expect *pbv1.GetMyPostsResponse
	}{
		"success": {
			req: &pbv1.GetMyPostsRequest{
				AccessToken: myCompanyToken,
			},
			expect: &pbv1.GetMyPostsResponse{
				Status:  200,
				Message: "My posts retrieved successfully",
				Posts: []*pbv1.Post{
					post2,
					post3,
					post4,
				},
			},
		},
		"invalid token": {
			req: &pbv1.GetMyPostsRequest{
				AccessToken: "invalid token",
			},
			expect: &pbv1.GetMyPostsResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
		"not company": {
			req: &pbv1.GetMyPostsRequest{
				AccessToken: ad,
			},
			expect: &pbv1.GetMyPostsResponse{
				Status:  403,
				Message: "Forbidden",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.GetMyPosts(ctx, tc.req)
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
				require.NotEmpty(t, p.UpdatedAt)
			}
		})
	}
}
