package e2e

import (
	"context"
	"testing"
	"time"

	"github.com/TikhampornSky/go-post-service/config"
	"github.com/TikhampornSky/go-post-service/domain"
	pbv1 "github.com/TikhampornSky/go-post-service/gen/v1"
	"github.com/TikhampornSky/go-post-service/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestUpdatePost(t *testing.T) {
	conn, err := grpc.Dial(":8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("could not connect to grpc server: %v", err)
	}
	defer conn.Close()

	c := pbv1.NewPostServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	config, _ := config.LoadConfig("..")
	token, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 1,
		Role:   "company",
	})
	require.NoError(t, err)

	tokenStudent, err := mock.GenerateAccessToken(config.AccessTokenExpiredInTest, &domain.Payload{
		UserId: 2,
		Role:   "student",
	})
	require.NoError(t, err)

	resDelete, err := c.DeletePosts(ctx, &pbv1.DeletePostsRequest{
		AccessToken: token,
	})
	require.NoError(t, err)
	require.Equal(t, int64(200), resDelete.Status)

	lex := `{ "root": {} }`

	updatedLex := `{
		"root": {
		  "children": ["child"]
		}
	  }
	`

	topic := "What to expect from here on out"
	description := lex
	period := "1 month"
	howTo := lex
	openPositions := []string{"Software Engineer", "Data Scientist"}
	requiredSkills := []string{"Golang", "Python"}
	benefits := []string{"Free lunch", "Free dinner"}

	updatedTopic := "NEW What to expect from here on out"
	updatedDescription := updatedLex
	updatedPeriod := "NEW 1 month"
	updatedHowTo := updatedLex
	updatedOpenPositions := []*pbv1.Element{
		{ Value: "NEW Software Engineer", Action: pbv1.ElementStatus_ADD, },
		{ Value: "NEW Data Scientist", Action: pbv1.ElementStatus_ADD, },
		{ Value: "Software Engineer", Action: pbv1.ElementStatus_REMOVE, },
		{ Value: "Data Scientist", Action: pbv1.ElementStatus_REMOVE, },
	}
	updatedRequiredSkills := []*pbv1.Element{
		{ Value: "Golang", Action: pbv1.ElementStatus_REMOVE, },
		{ Value: "Python", Action: pbv1.ElementStatus_SAME, },
		{ Value: "NEW Java", Action: pbv1.ElementStatus_ADD, },
	}
	updatedBenefits := []*pbv1.Element{
		{ Value: "NEW Free lunch", Action: pbv1.ElementStatus_ADD, },
		{ Value: "Free lunch", Action: pbv1.ElementStatus_REMOVE, },
		{ Value: "Free dinner", Action: pbv1.ElementStatus_REMOVE, },
	}

	CreateRes, err := c.CreatePost(ctx, &pbv1.CreatePostRequest{
		AccessToken: token,
		Post: &pbv1.Post{
			Topic:          topic,
			Description:    description,
			Period:         period,
			HowTo:          howTo,
			OpenPositions:  openPositions,
			RequiredSkills: requiredSkills,
			Benefits:       benefits,
		},
	})
	require.NoError(t, err)
	require.Equal(t, int64(201), CreateRes.Status)
	p := &pbv1.UpdatedPost{
		Topic:          updatedTopic,
		Description:    updatedDescription,
		Period:         updatedPeriod,
		HowTo:          updatedHowTo,
		OpenPositions:  updatedOpenPositions,
		RequiredSkills: updatedRequiredSkills,
		Benefits:       updatedBenefits,
	}

	tests := map[string]struct {
		req    *pbv1.UpdatePostRequest
		expect *pbv1.UpdatePostResponse
	}{
		"success": {
			req: &pbv1.UpdatePostRequest{
				AccessToken: token,
				Id:          CreateRes.Id,
				Post:        p,
			},
			expect: &pbv1.UpdatePostResponse{
				Status:  200,
				Message: "Post updated successfully",
			},
		},
		"fail: unauthorized": {
			req: &pbv1.UpdatePostRequest{
				AccessToken: tokenStudent,
				Id:          CreateRes.Id,
				Post:        p,
			},
			expect: &pbv1.UpdatePostResponse{
				Status:  401,
				Message: "Unauthorized",
			},
		},
		"some field is empty": {
			req: &pbv1.UpdatePostRequest{
				AccessToken: token,
				Id:          CreateRes.Id,
				Post: &pbv1.UpdatedPost{
					Topic:          "",
					Description:    updatedDescription,
					Period:         updatedPeriod,
					HowTo:          updatedHowTo,
					OpenPositions:  updatedOpenPositions,
					RequiredSkills: updatedRequiredSkills,
					Benefits:       updatedBenefits,
				},
			},
			expect: &pbv1.UpdatePostResponse{
				Status:  400,
				Message: "Please fill all required fields",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := c.UpdatePost(ctx, tc.req)
			require.NoError(t, err)
			require.Equal(t, tc.expect.Status, res.Status)
			require.Equal(t, tc.expect.Message, res.Message)
			if tc.expect.Status == 200 {
				g, err := c.GetPost(ctx, &pbv1.GetPostRequest{
					AccessToken: token,
					Id:          CreateRes.Id,
				})
				require.NoError(t, err)
				require.Equal(t, updatedTopic, g.Post.Topic)
				require.Equal(t, updatedDescription, g.Post.Description)
				require.Equal(t, updatedPeriod, g.Post.Period)
				require.Equal(t, updatedHowTo, g.Post.HowTo)
				checkOpen := 0
				for _, v := range g.Post.OpenPositions {
					if v == "NEW Software Engineer" {
						checkOpen += 1
					}
					if v == "NEW Data Scientist" {
						checkOpen += 1
					}
				}
				require.Equal(t, 2, checkOpen)
				require.Equal(t, 2, len(g.Post.OpenPositions))

				checkRequired := 0
				for _, v := range g.Post.RequiredSkills {
					if v == "Python" {
						checkRequired += 1
					}
					if v == "NEW Java" {
						checkRequired += 1
					}
				}
				require.Equal(t, 2, checkRequired)

				checkBenefits := 0
				for _, v := range g.Post.Benefits {
					if v == "NEW Free lunch" {
						checkBenefits += 1
					}
				}
				require.Equal(t, 1, checkBenefits)
			}
		})
	}
}