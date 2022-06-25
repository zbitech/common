package rctx

import (
	"context"
	"github.com/zbitech/common/pkg/errs"
	"github.com/zbitech/common/pkg/model/entity"

	"github.com/golang-jwt/jwt"
	"github.com/zbitech/common/pkg/model/object"
	"github.com/zbitech/common/pkg/vars"
)

type ContextParam string

type ContextFn func(context.Context) context.Context

const (
	RequestId ContextParam = "RequestId"
	Component ContextParam = "Component"
	StartTime ContextParam = "StartTime"
	JwtClaims ContextParam = "JwtClaims"
	User      ContextParam = "User"
)

var (
	CTX = BuildContext(context.Background(), Context(RequestId, "zbi"), Context(Component, "zbi"))
)

func Context(key ContextParam, value interface{}) ContextFn {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, key, value)
	}
}

func GetRequestId(ctx context.Context, def_value string) string {

	rawRequestId := ctx.Value(RequestId)
	requestId, ok := rawRequestId.(string)

	if !ok {
		return def_value
	}

	return requestId
}

func GetCurrentUser(ctx context.Context) *object.CurrentUser {
	claims, cErr := GetClaims(ctx)
	user, uErr := GetUser(ctx)

	if cErr != nil || uErr != nil {
		return object.GetAnonymousUser()
	}

	jwtServer := vars.ServiceFactory.GetJwtServer()

	userid := jwtServer.GetUserId(claims)
	email := jwtServer.GetEmail(claims)
	role := jwtServer.GetRole(claims)

	return object.GetCurrentUser(userid, email, role, user)
}

func GetUser(ctx context.Context) (*entity.User, error) {
	rawUser := ctx.Value(User)
	user, ok := rawUser.(*entity.User)

	if !ok {
		return nil, errs.ErrUnregisteredUser
	}

	return user, nil
}

func GetClaims(ctx context.Context) (jwt.Claims, error) {
	rawClaims := ctx.Value(JwtClaims)
	claims, ok := rawClaims.(jwt.Claims)

	if !ok {
		return nil, errs.ErrInvalidToken
	} else if err := claims.Valid(); err != nil {
		return nil, err
	}

	return claims, nil
}

func BuildContext(ctx context.Context, ctxFns ...ContextFn) context.Context {
	for _, f := range ctxFns {
		ctx = f(ctx)
	}

	return ctx
}
