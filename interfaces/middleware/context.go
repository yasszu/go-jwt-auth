package middleware

import "context"

type contextKey int

const (
	ContextKeyAccountID contextKey = iota
)

func SetAccountID(ctx context.Context, accountID uint) context.Context {
	return context.WithValue(ctx, ContextKeyAccountID, accountID)
}

func GetAccountID(ctx context.Context) (uint, bool) {
	accountID, ok := ctx.Value(ContextKeyAccountID).(uint)
	return accountID, ok
}
