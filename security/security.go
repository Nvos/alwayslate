package security

import (
	"alwayslate/x/errorsx"
	"context"
)

const ErrUnauthenticated = errorsx.Sentinel("unauthenticated")
var authCtxKey = &contextKey{"auth"}

type contextKey struct {
	name string
}

type User struct {
	ID int
	Username string
	Role string
}

type Credentials struct {
	Username string
	Password string
}

type Hasher interface {
	Encode(value string) (string, error)
	Decode(value string) (string, error)
	Verify(value string, against string) (bool, error)
}

func WithUser(ctx context.Context, u User) context.Context {
	return context.WithValue(ctx, authCtxKey, u)
}

func FromContext(ctx context.Context) (User, error) {
	value := ctx.Value(authCtxKey)
	if value == nil {
		return User{}, ErrUnauthenticated
	}

	if user, ok := value.(User); ok {
		return user, nil
	}

	return User{}, ErrUnauthenticated
}