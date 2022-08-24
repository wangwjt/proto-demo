package interceptor

import (
	"context"
	"google.golang.org/grpc"
)

type Token struct {
}

func (t *Token) UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	return nil, nil
}
