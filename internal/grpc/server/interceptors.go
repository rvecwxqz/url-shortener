package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UserIDInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	var id string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("user_id")
		if len(values) > 0 {
			id = values[0]
		}
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	if len(id) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing user_id")
	}

	return handler(ctx, req)
}
