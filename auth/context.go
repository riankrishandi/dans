package auth

import (
	"context"
	"errors"
)

type contextKey string

var (
	claimsKey contextKey = "claims"

	errClaimsNotFound = errors.New("claims not found")
)

func NewContext(ctx context.Context, claims CustomClaims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

func FromContext(ctx context.Context) (*CustomClaims, error) {
	claims, ok := ctx.Value(claimsKey).(CustomClaims)
	if !ok {
		return nil, errClaimsNotFound
	}

	return &claims, nil
}
