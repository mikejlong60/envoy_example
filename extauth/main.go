package main

import (
	"context"
	"fmt"
	"net"
	"strings"

	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/googleapis/google/rpc"
	status "google.golang.org/genproto/googleapis/rpc/status"

	"google.golang.org/grpc"
)

type authorizationServer struct{}

func (a *authorizationServer) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	headers := req.GetAttributes().GetRequest().GetHttp().GetHeaders()
	authHeader := headers["authorization"]
	userDN := headers["user_dn"]
	bar := headers["bar"]
	fmt.Printf("authtoken in header: %v", authHeader)
	fmt.Println("\n==============")
	fmt.Printf("userDN in header: %v", userDN)
	fmt.Println("\n==============")
	fmt.Printf("Bar in header: %v", bar)
	fmt.Println("\n==============")

	var splitToken []string
	if authHeader != "" {
		splitToken = strings.Split(authHeader, "Bearer ")
	}

	if len(splitToken) == 2 {
		token := splitToken[1]

		if len(token) == 3 {
			return &auth.CheckResponse{
				Status: &status.Status{
					Code: int32(rpc.OK),
				},
				HttpResponse: &auth.CheckResponse_OkResponse{
					OkResponse: &auth.OkHttpResponse{
						Headers: []*core.HeaderValueOption{
							{
								Header: &core.HeaderValue{
									Key:   "my-credential-header",
									Value: "permission6,permission9",
								},
							},
						},
					},
				},
			}, nil
		}
	}

	return &auth.CheckResponse{
		Status: &status.Status{
			Code: int32(rpc.UNAUTHENTICATED),
		},
		HttpResponse: &auth.CheckResponse_DeniedResponse{
			DeniedResponse: &auth.DeniedHttpResponse{
				Status: &envoy_type.HttpStatus{Code: 401},
				Body:   "Authorization header must be formatted as Authorization: Bearer TOKEN, where TOKEN is a three-character string",
			},
		},
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":4040")
	grpcServer := grpc.NewServer()
	authServer := &authorizationServer{}
	auth.RegisterAuthorizationServer(grpcServer, authServer)
	grpcServer.Serve(lis)
}
