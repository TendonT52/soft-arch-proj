package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"github.com/tendont52/api-gateway/config"
	postService "github.com/tendont52/api-gateway/gen/post-service/v1"
	reportService "github.com/tendont52/api-gateway/gen/report-service/v1"
	reviewService "github.com/tendont52/api-gateway/gen/review-service/v1"
	userService "github.com/tendont52/api-gateway/gen/user-service/v1"
	_ "github.com/tendont52/api-gateway/statik"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Serve(conf *config.Config) error {

	fmt.Println("Starting API Gateway...")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux := runtime.NewServeMux(
		runtime.WithForwardResponseOption(TranformOutgoingResponse),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := userService.RegisterUserServiceHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf("%s:%s", conf.UserServiceHost, conf.UserServicePort), opts)
	if err != nil {
		log.Fatalf("cannot register user service: %v", err)
	}
	err = userService.RegisterAuthServiceHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf("%s:%s", conf.UserServiceHost, conf.UserServicePort), opts)
	if err != nil {
		log.Fatalf("cannot register user service: %v", err)
	}
	err = postService.RegisterPostServiceHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf("%s:%s", conf.PostServiceHost, conf.PostServicePort), opts)
	if err != nil {
		log.Fatalf("cannot register post service: %v", err)
	}
	err = reportService.RegisterReportServiceHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf("%s:%s", conf.ReportServiceHost, conf.ReportServicePort), opts)
	if err != nil {
		log.Fatalf("cannot register report service: %v", err)
	}
	err = reviewService.RegisterReviewServiceHandlerFromEndpoint(ctx, gwmux, fmt.Sprintf("%s:%s", conf.ReviewServiceHost, conf.ReviewServicePort), opts)
	if err != nil {
		log.Fatalf("cannot register review service: %v", err)
	}

	gwServer := &http.Server{
		Addr: fmt.Sprintf(":%s", conf.RESTPort),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if strings.HasPrefix(r.URL.Path, "/v") {
				TransformIncomingRequest(w, r)
				gwmux.ServeHTTP(w, r)
			} else {
				SwaggerHandler().ServeHTTP(w, r)
			}

		}),
	}

	log.Printf("Listening on port %v ...", conf.RESTPort)
	return gwServer.ListenAndServe()
}

func TransformIncomingRequest(w http.ResponseWriter, r *http.Request) {

	refreshToken, err := r.Cookie("refreshToken")
	if err == nil {
		var body map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("cannot read body"))
			return
		}
		body["refreshToken"] = refreshToken.Value
		json, err := json.Marshal(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("cannot read body"))
			return
		}
		r.Body = io.NopCloser(strings.NewReader(string(json)))
	}
}

func TranformOutgoingResponse(ctx context.Context, w http.ResponseWriter, resp proto.Message) error {
	resp.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.Name() == "refresh_token" {
			resp.ProtoReflect().Clear(fd)
			http.SetCookie(w, &http.Cookie{
				Name:     "refreshToken",
				Value:    v.String(),
				HttpOnly: true,
				Path:     "/v1/refresh",
			})
		}
		return true
	})
	resp.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.Name() == "status" {
			w.WriteHeader(int(v.Int()))
		}
		return true
	})

	return nil
}

func SwaggerHandler() http.Handler {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	return http.FileServer(statikFS)
}
