package entity

import (
	"github.com/zbitech/common/pkg/model/ztypes"
	"time"
)

type UserTeam struct {
	TeamId string `json:"teamid"`
	Key    string `json:"key"`
}

type User struct {
	UserId      string                   `json:"userid" bson:"userid"`
	Name        string                   `json:"name" bson:"name"`
	Email       string                   `json:"email" bson:"email"`
	Role        ztypes.Role              `json:"role" bson:"role"`
	Memberships []UserTeam               `json:"memberships" bson:"memberships"`
	Created     time.Time                `json:"created" bson:"created"`
	Active      bool                     `json:"active" bson:"active"`
	LastUpdate  time.Time                `json:"lastupdate" bson:"lastupdate"`
	Level       ztypes.SubscriptionLevel `json:"level" bson:"level"`
}

type UserPassword struct {
	UserId     string    `json:"userid" bson:"userid"`
	Password   string    `json:"password" bson:"password"`
	LastUpdate time.Time `json:"lastupdate" bson:"lastUpdate"`
}

func NewUserPassword(userId, password string) UserPassword {
	return UserPassword{UserId: userId, Password: password, LastUpdate: time.Now()}
}

func NewUser(userId, name, email string, role ztypes.Role) *User {
	user := &User{
		UserId: userId, Name: name, Email: email,
		Role:        role,
		Memberships: make([]UserTeam, 0),
	}

	return user
}

func (u *User) AddTeam(teamId, key string) {
	for _, t := range u.Memberships {
		if t.TeamId == teamId {
			return
		}
	}

	u.Memberships = append(u.Memberships, UserTeam{TeamId: teamId, Key: key})
}

func (u *User) GetTeam(teamId string) *UserTeam {
	for _, t := range u.Memberships {
		if t.TeamId == teamId {
			return &t
		}
	}

	return nil
}
