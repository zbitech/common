package object

import (
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/ztypes"
)

type CurrentUser struct {
	Authenticated bool
	UserId        string
	Email         string
	Role          ztypes.Role
	User          *entity.User
}

func GetAnonymousUser() *CurrentUser {
	return &CurrentUser{
		Authenticated: false,
	}
}

func GetCurrentUser(userid, email string, role ztypes.Role, user *entity.User) *CurrentUser {
	return &CurrentUser{
		Authenticated: true,
		UserId:        userid,
		Email:         email,
		Role:          role,
		User:          user,
	}
}

func (c CurrentUser) IsAuthenticated() bool {
	return c.Authenticated
}

func (c CurrentUser) IsAnonymous() bool {
	return !c.Authenticated
}

func (c CurrentUser) GetUserId() string {
	return c.UserId
}

func (c CurrentUser) GetEmail() string {
	return c.Email
}

func (c CurrentUser) IsAdmin() bool {
	return c.Role == ztypes.RoleAdmin
}

func (c CurrentUser) IsOwner() bool {
	return c.Role == ztypes.RoleOwner
}

func (c CurrentUser) IsUser() bool {
	return c.Role == ztypes.RoleUser
}
