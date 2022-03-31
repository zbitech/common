package entity

import (
	"github.com/zbitech/common/pkg/id"
	"time"
)

type Team struct {
	TeamId     string    `json:"id" bson:"_id"`
	Name       string    `json:"name" bson:"name"`
	Email      string    `json:"email" bson:"email"`
	Owner      string    `json:"owner" bson:"owner"`
	Created    time.Time `json:"created" bson:"created"`
	LastUpdate time.Time `json:"lastupdate" bson:"lastupdate"`
}

func NewTeam(name, email, owner string) Team {
	return Team{TeamId: id.GenerateTeamID(), Name: name, Email: email, Owner: owner, Created: time.Now(), LastUpdate: time.Now()}
}

func (t *Team) IsOwner(userid string) bool {
	return t.Owner == userid
}
