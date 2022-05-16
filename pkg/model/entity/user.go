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
	LastName    string                   `json:"lastname" bson:"lastname"`
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

func NewUser(userId, name, lastName, email string, role ztypes.Role, level ztypes.SubscriptionLevel) *User {
	user := &User{
		UserId: userId, Name: name, LastName: lastName,
		Email:       email,
		Role:        role,
		Level:       level,
		Active:      true,
		Memberships: make([]UserTeam, 0),
		Created:     time.Now(),
		LastUpdate:  time.Now(),
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

func (u *User) RemoveTeam(key string) {
	for index, t := range u.Memberships {
		if t.Key == key {
			u.Memberships = append(u.Memberships[:index], u.Memberships[index+1:]...)
			return
		}
	}
}

func (u *User) GetTeam(teamId string) *UserTeam {
	for _, t := range u.Memberships {
		if t.TeamId == teamId {
			return &t
		}
	}

	return nil
}
