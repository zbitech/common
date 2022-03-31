package object

import (
	"github.com/golang-jwt/jwt"
	"github.com/zbitech/common/pkg/model/ztypes"
)

type ZBIBasicClaims struct {
	Email string
	Role  ztypes.Role
	jwt.StandardClaims
}
