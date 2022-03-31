package entity

import (
	"github.com/zbitech/common/pkg/id"
	"time"

	"github.com/zbitech/common/pkg/model/ztypes"
)

type TeamMember struct {
	TeamId     string                  `json:"teamid"`
	Key        string                  `json:"key"`
	Email      string                  `json:"email"`
	Role       ztypes.Role             `json:"role"`
	Status     ztypes.InvitationStatus `json:"status"`
	CreatedOn  time.Time               `json:"createdon"`
	ExpiresOn  time.Time               `json:"expireson"`
	LastUpdate time.Time               `json:"lastUpdate"`
}

func NewTeamMember(teamId, email string, role ztypes.Role) TeamMember {
	return TeamMember{TeamId: teamId, Email: email, Key: id.GenerateTeamMemberKey(), Role: role,
		Status: ztypes.NEW_INVITATION, CreatedOn: time.Now(), ExpiresOn: time.Now().Add(time.Hour * 168)}
}

func (m *TeamMember) IsJoined() bool {
	return m.Status == ztypes.ACCEPT_INVITATION
}

func (m *TeamMember) IsAdmin() bool {
	return m.Role == ztypes.ADMIN_ROLE
}

func (m *TeamMember) IsExpired() bool {
	return m.ExpiresOn.Before(time.Now())
}
