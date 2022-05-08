package auth

import (
	"context"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"

	"github.com/karamaru-alpha/melt/pkg/mcontext"
	"github.com/karamaru-alpha/melt/pkg/merrors"
)

var skipAuthMethods = []string{"/api.Auth/Signup", "/api.Auth/RefreshToken"}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return grpc_auth.UnaryServerInterceptor(authFunc)
}

func authFunc(ctx context.Context) (context.Context, error) {
	// 認証スキップするエンドポイントは飛ばす
	method, _ := grpc.Method(ctx)
	for _, v := range skipAuthMethods {
		if method == v {
			return ctx, nil
		}
	}
	// メタデータからアクセストークン文字列を取り出す
	tokenString, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, merrors.Wrap(err, merrors.Unauthenticated)
	}
	// アクセストークンからユーザーIDを取り出しcontextに詰める
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, merrors.Newf(merrors.Unauthenticated, "token method is invalid")
		}
		return []byte(os.Getenv("TOKEN_SIGNED_STRING")), nil
	})
	if err != nil {
		return nil, merrors.Wrap(err, merrors.Unauthenticated)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, merrors.Newf(merrors.Unauthenticated, "token claim is invalid")
	}
	userID := claims["sub"].(string)
	mctx := mcontext.New(userID)
	return mcontext.SetInContext(ctx, mctx), nil
}
