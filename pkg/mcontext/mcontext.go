package mcontext

import (
	"context"
)

type Context interface {
	GetUserID() string
}

type contextKey struct{}

type meltContext struct {
	userID string
}

func (c *meltContext) GetUserID() string {
	if c == nil {
		return ""
	}
	return c.userID
}

func New(userID string) Context {
	return &meltContext{
		userID: userID,
	}
}

// SetInContext contextにmeltContextを設定する
func SetInContext(ctx context.Context, meltContext Context) context.Context {
	return context.WithValue(ctx, contextKey{}, meltContext)
}

// Extract contextからmeltContextを取得する
func Extract(ctx context.Context) Context {
	return ctx.Value(contextKey{}).(*meltContext)
}
