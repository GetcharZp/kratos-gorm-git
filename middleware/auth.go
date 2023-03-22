package middleware

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"kratos-gorm-git/helper"
)

func Auth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				auth := tr.RequestHeader().Get("Authorization")
				if auth == "" {
					return nil, errors.New("no auth")
				}
				userClaims, err := helper.AnalyseToken(auth)
				if err != nil {
					return nil, err
				}
				if userClaims.Identity == "" {
					return nil, errors.New("no auth")
				}
				ctx = metadata.NewServerContext(ctx, metadata.New(map[string]string{
					"username": userClaims.Name,
					"identity": userClaims.Identity,
				}))
			}
			return handler(ctx, req)
		}
	}
}
