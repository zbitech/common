package jwtutil

import (
	"context"
	"github.com/zbitech/common/pkg/id"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/object"
	"github.com/zbitech/common/pkg/rctx"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/zbitech/common/pkg/errs"
	"github.com/zbitech/common/pkg/logger"
	"github.com/zbitech/common/pkg/vars"
)

var (
//	SECRET_KEY = "KaPdSgVkYp3s6v9y$B&E)H@McQeThWmZ"
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

	iamSvc := vars.ServiceFactory.GetIAMService()
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

func GenerateJwtToken(user entity.User) (*string, error) {

	now := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, object.ZBIBasicClaims{
		StandardClaims: jwt.StandardClaims{
			Audience:  "ZBI",
			ExpiresAt: now.Add(time.Hour * time.Duration(vars.AppConfig.Policy.TokenExpirationPolicy)).Unix(),
			Id:        id.GenerateRequestID(),
			IssuedAt:  now.Unix(),
			Issuer:    "ZBI",
			NotBefore: now.Unix(),
			Subject:   user.UserId,
		},
		Role:  user.Role,
		Email: user.Email,
	})

	signedToken, err := token.SignedString([]byte(vars.AppConfig.Repository.JwtConfig.SecretKey))
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}

func ParseWithClaims(tokenString string, claims jwt.Claims, validateToken func(token *jwt.Token) error) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		err := validateToken(token)
		if err != nil {
			return nil, err
		}

		return []byte(vars.AppConfig.Repository.JwtConfig.SecretKey), nil
	})

	return token, err
}
