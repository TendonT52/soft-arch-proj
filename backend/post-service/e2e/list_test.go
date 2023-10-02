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
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetOpenPositions(t *testing.T) {
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

	token, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 1,
		Role:   "company",
	})
	require.NoError(t, err)
	tokenStudent, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 2,
		Role:   "student",
	})

	err = tools.DeleteAllPosts()
	require.NoError(t, err)

	ad, err := mock.CreateMockAdmin(ctx)
	require.NoError(t, err)
	com1Res, _, err := mock.CreateMockApprovedCompany(ctx, "Grab", ad)
	require.NoError(t, err)
	com2Res, _, err := mock.CreateMockApprovedCompany(ctx, "Agoda", ad)
	require.NoError(t, err)
	com3Res, _, err := mock.CreateMockApprovedCompany(ctx, "Lineman", ad)
	require.NoError(t, err)

	post1 := createMockPost(t, ctx, c, com1Res.Id, "Grab", "Post 1", lex, "1 month", lex,
		[]string{"Software Engineer", "Data Scientist"}, []string{"Golang", "Python"}, []string{"Free lunch", "Free dinner"})
	post2 := createMockPost(t, ctx, c, com2Res.Id, "Agoda", "Post 2", lex, "2 month", lex,
		[]string{"Data Analysts", "Full-Stack Developer"}, []string{"Python", "HTML", "CSS"}, []string{"Free lunch", "Macbook Pro"})
	post3 := createMockPost(t, ctx, c, com2Res.Id, "Agoda", "Post 3", lex, "3 month", lex,
		[]string{"Backend Developer", "Data Scientist"}, []string{"Golang", "Python"}, []string{"Free dinner", "Macbook M1"})
	post4 := createMockPost(t, ctx, c, com2Res.Id, "Agoda", "Post 4", lex, "4 month", lex,
		[]string{"Frontend Developer", "Data Analyst"}, []string{"HTML", "CSS", "Javascript"}, []string{"Free lunch", "Free dinner", "Macbook M1"})
	post5 := createMockPost(t, ctx, c, com3Res.Id, "Lineman", "Post 5", lex, "5 month", lex,
		[]string{"Frontend Developer", "Data Analyst"}, []string{"HTML", "CSS", "Javascript"}, []string{"Free lunch", "Free dinner", "Macbook Pro"})

	_ = post1
	_ = post2
	_ = post3
	_ = post4
	_ = post5

	tests := map[string]struct {
		req    *pbv1.GetOpenPositionsRequest
		expect *pbv1.GetOpenPositionsResponse
	}{
		"success": {
			req: &pbv1.GetOpenPositionsRequest{
				AccessToken: token,
				Search:      "Frontend",
			},
			expect: &pbv1.GetOpenPositionsResponse{
				Status:        200,
				Message:       "Open positions retrieved successfully",
				OpenPositions: []string{"Frontend Developer", "Backend Developer", "Full-Stack Developer"},
			},
		},
		"empty search": {
			req: &pbv1.GetOpenPositionsRequest{
				AccessToken: token,
				Search:      "",
			},
			expect: &pbv1.GetOpenPositionsResponse{
				Status:        200,
				Message:       "Open positions retrieved successfully",
				OpenPositions: []string{"Backend Developer", "Data Analyst", "Data Analysts", "Data Scientist", "Frontend Developer", "Full-Stack Developer", "Software Engineer"},
			},
		},
		"not foind any match": {
			req: &pbv1.GetOpenPositionsRequest{
				AccessToken: token,
				Search:      "xxyyyyzzz",
			},
			expect: &pbv1.GetOpenPositionsResponse{
				Status:        200,
				Message:       "Open positions retrieved successfully",
				OpenPositions: []string(nil),
			},
		},
		"not company": {
			req: &pbv1.GetOpenPositionsRequest{
				AccessToken: tokenStudent,
				Search:      "Developer",
			},
			expect: &pbv1.GetOpenPositionsResponse{
				Status:  403,
				Message: "Forbidden",
			},
		},
		"invalid token": {
			req: &pbv1.GetOpenPositionsRequest{
				AccessToken: "invalid token",
				Search:      "Developer",
			},
			expect: &pbv1.GetOpenPositionsResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.GetOpenPositions(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			require.Equal(t, tc.expect.OpenPositions, res.OpenPositions)
		})
	}
}

func TestGetRequiredSkills(t *testing.T) {
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

	token, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 1,
		Role:   "company",
	})
	require.NoError(t, err)
	tokenStudent, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 2,
		Role:   "student",
	})

	err = tools.DeleteAllPosts()
	require.NoError(t, err)

	ad, err := mock.CreateMockAdmin(ctx)
	require.NoError(t, err)
	com1Res, _, err := mock.CreateMockApprovedCompany(ctx, "Grab", ad)
	require.NoError(t, err)
	com2Res, _, err := mock.CreateMockApprovedCompany(ctx, "KBTG", ad)
	require.NoError(t, err)
	com3Res, _, err := mock.CreateMockApprovedCompany(ctx, "SCB", ad)
	require.NoError(t, err)

	post1 := createMockPost(t, ctx, c, com1Res.Id, "Grab", "Post 1", lex, "1 month", lex,
		[]string{"Software Engineer", "Data Scientist"}, []string{"Golang", "Python"}, []string{"Free lunch", "Free dinner"})
	post2 := createMockPost(t, ctx, c, com2Res.Id, "KBTG", "Post 2", lex, "2 month", lex,
		[]string{"Data Analysts", "Full-Stack Developer"}, []string{"Python", "HTML", "CSS"}, []string{"Free lunch", "Macbook Pro"})
	post3 := createMockPost(t, ctx, c, com2Res.Id, "KBTG", "Post 3", lex, "3 month", lex,
		[]string{"Backend Developer", "Data Scientist"}, []string{"Golang", "Python"}, []string{"Free dinner", "Macbook M1"})
	post4 := createMockPost(t, ctx, c, com2Res.Id, "KBTG", "Post 4", lex, "4 month", lex,
		[]string{"Frontend Developer", "Data Analyst"}, []string{"HTML", "CSS", "Javascript"}, []string{"Free lunch", "Free dinner", "Macbook M1"})
	post5 := createMockPost(t, ctx, c, com3Res.Id, "SCB", "Post 5", lex, "5 month", lex,
		[]string{"Frontend Developer", "Data Analyst"}, []string{"HTML", "CSS", "Javascript"}, []string{"Free lunch", "Free dinner", "Macbook Pro"})

	_ = post1
	_ = post2
	_ = post3
	_ = post4
	_ = post5

	tests := map[string]struct {
		req    *pbv1.GetRequiredSkillsRequest
		expect *pbv1.GetRequiredSkillsResponse
	}{
		"success": {
			req: &pbv1.GetRequiredSkillsRequest{
				AccessToken: token,
				Search:      "Go",
			},
			expect: &pbv1.GetRequiredSkillsResponse{
				Status:         200,
				Message:        "Required skills retrieved successfully",
				RequiredSkills: []string{"Golang"},
			},
		},
		"empty search": {
			req: &pbv1.GetRequiredSkillsRequest{
				AccessToken: token,
				Search:      "",
			},
			expect: &pbv1.GetRequiredSkillsResponse{
				Status:         200,
				Message:        "Required skills retrieved successfully",
				RequiredSkills: []string{"CSS", "Golang", "HTML", "Javascript", "Python"},
			},
		},
		"not foind any match": {
			req: &pbv1.GetRequiredSkillsRequest{
				AccessToken: token,
				Search:      "xxyyyyzzz",
			},
			expect: &pbv1.GetRequiredSkillsResponse{
				Status:         200,
				Message:        "Required skills retrieved successfully",
				RequiredSkills: []string(nil),
			},
		},
		"not Company": {
			req: &pbv1.GetRequiredSkillsRequest{
				AccessToken: tokenStudent,
				Search:      "Go",
			},
			expect: &pbv1.GetRequiredSkillsResponse{
				Status:  403,
				Message: "Forbidden",
			},
		},
		"invalid token": {
			req: &pbv1.GetRequiredSkillsRequest{
				AccessToken: "invalid token",
				Search:      "Go",
			},
			expect: &pbv1.GetRequiredSkillsResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.GetRequiredSkills(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			require.Equal(t, tc.expect.RequiredSkills, res.RequiredSkills)
		})
	}
}

func TestGetBenefits(t *testing.T) {
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

	token, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 1,
		Role:   "company",
	})
	require.NoError(t, err)
	tokenStudent, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 2,
		Role:   "student",
	})

	err = tools.DeleteAllPosts()
	require.NoError(t, err)

	ad, err := mock.CreateMockAdmin(ctx)
	require.NoError(t, err)
	com1Res, _, err := mock.CreateMockApprovedCompany(ctx, "TTB", ad)
	require.NoError(t, err)
	com2Res, _, err := mock.CreateMockApprovedCompany(ctx, "Wongnai", ad)
	require.NoError(t, err)
	com3Res, _, err := mock.CreateMockApprovedCompany(ctx, "KBTG", ad)
	require.NoError(t, err)

	post1 := createMockPost(t, ctx, c, com1Res.Id, "TTB", "Post 1", lex, "1 month", lex,
		[]string{"Software Engineer", "Data Scientist"}, []string{"Golang", "Python"}, []string{"Free lunch", "Free dinner"})
	post2 := createMockPost(t, ctx, c, com2Res.Id, "Wongnai", "Post 2", lex, "2 month", lex,
		[]string{"Data Analysts", "Full-Stack Developer"}, []string{"Python", "HTML", "CSS"}, []string{"Free lunch", "Macbook Pro"})
	post3 := createMockPost(t, ctx, c, com2Res.Id, "Wongnai", "Post 3", lex, "3 month", lex,
		[]string{"Backend Developer", "Data Scientist"}, []string{"Golang", "Python"}, []string{"Free dinner", "Macbook M1"})
	post4 := createMockPost(t, ctx, c, com2Res.Id, "Wongnai", "Post 4", lex, "4 month", lex,
		[]string{"Frontend Developer", "Data Analyst"}, []string{"HTML", "CSS", "Javascript"}, []string{"Free lunch", "Free dinner", "Macbook M1"})
	post5 := createMockPost(t, ctx, c, com3Res.Id, "KBTG", "Post 5", lex, "5 month", lex,
		[]string{"Frontend Developer", "Data Analyst"}, []string{"HTML", "CSS", "Javascript"}, []string{"Free lunch", "Free dinner", "Macbook Pro"})

	_ = post1
	_ = post2
	_ = post3
	_ = post4
	_ = post5

	tests := map[string]struct {
		req    *pbv1.GetBenefitsRequest
		expect *pbv1.GetBenefitsResponse
	}{
		"success": {
			req: &pbv1.GetBenefitsRequest{
				AccessToken: token,
				Search:      "free",
			},
			expect: &pbv1.GetBenefitsResponse{
				Status:   200,
				Message:  "Benefits retrieved successfully",
				Benefits: []string{"Free lunch", "Free dinner"},
			},
		},
		"empty search": {
			req: &pbv1.GetBenefitsRequest{
				AccessToken: token,
				Search:      "",
			},
			expect: &pbv1.GetBenefitsResponse{
				Status:   200,
				Message:  "Benefits retrieved successfully",
				Benefits: []string{"Free dinner", "Free lunch", "Macbook M1", "Macbook Pro"},
			},
		},
		"not foind any match": {
			req: &pbv1.GetBenefitsRequest{
				AccessToken: token,
				Search:      "xxyyyyzzz",
			},
			expect: &pbv1.GetBenefitsResponse{
				Status:   200,
				Message:  "Benefits retrieved successfully",
				Benefits: []string(nil),
			},
		},
		"not Company": {
			req: &pbv1.GetBenefitsRequest{
				AccessToken: tokenStudent,
				Search:      "free",
			},
			expect: &pbv1.GetBenefitsResponse{
				Status:  403,
				Message: "Forbidden",
			},
		},
		"invalid token": {
			req: &pbv1.GetBenefitsRequest{
				AccessToken: "invalid token",
				Search:      "free",
			},
			expect: &pbv1.GetBenefitsResponse{
				Status:  401,
				Message: "Your access token is invalid",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.GetBenefits(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			require.Equal(t, tc.expect.Benefits, res.Benefits)
		})
	}
}
