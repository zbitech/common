package jwtutil

import (
	"context"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/rctx"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/zbitech/common/pkg/errs"
	"github.com/zbitech/common/pkg/logger"
	"github.com/zbitech/common/pkg/vars"
)

func ValidateAuthorizationToken(ctx context.Context, authorization string) (jwt.Claims, *entity.User, error) {

	ctx = rctx.BuildContext(ctx, rctx.Context(rctx.Component, "ValidateAuthorizationToken"), rctx.Context(rctx.StartTime, time.Now()))
	defer logger.LogComponentTime(ctx)

	if len(authorization) == 0 {
		return nil, nil, nil
	}

	authHeaderParts := strings.Split(authorization, " ")
	method, credentials := authHeaderParts[0], authHeaderParts[1]

	if method != "Bearer" {
		return nil, nil, errs.ErrInvalidAuthMethod
	}

	iamSvc := vars.AuthorizationFactory.GetIAMService()
	claims, user, err := iamSvc.ValidateAuthToken(ctx, credentials) // tokenVerifier.ParseWithClaims(credentials)

	if err != nil {
		logger.Errorf(ctx, "Could not verify bearer token %s due to error %s", credentials, err)
		return nil, nil, errs.ErrTokenValidation
	}

	if err := claims.Valid(); err != nil {
		logger.Errorf(ctx, "Token %s is no longer valid", credentials)
		return nil, nil, errs.ErrInvalidToken
	}

	return claims, user, nil

}
