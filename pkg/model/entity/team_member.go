package entity

import (
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

func NewTeamMember(teamId, key, email string, admin bool) TeamMember {
	mbr := TeamMember{TeamId: teamId, Email: email, Key: key,
		Status: ztypes.InvitationPending, CreatedOn: time.Now(), ExpiresOn: time.Now().Add(time.Hour * 168)}

	if admin {
		mbr.Role = ztypes.RoleAdmin
	} else {
		mbr.Role = ztypes.RoleUser
	}

	return mbr
}

func (m *TeamMember) SetAdmin(admin bool) {
	if admin {
		m.Role = ztypes.RoleAdmin
	} else {
		m.Role = ztypes.RoleUser
	}
}

func (m *TeamMember) Accept() {
	m.Status = ztypes.InvitationAccept
	m.LastUpdate = time.Now()
}

func (m *TeamMember) Reject() {
	m.Status = ztypes.InvitationReject
	m.LastUpdate = time.Now()
}

func (m *TeamMember) Expire() {
	m.Status = ztypes.InvitationExpired
	m.LastUpdate = time.Now()
}

func (m *TeamMember) IsJoined() bool {
	return m.Status == ztypes.InvitationAccept
}

func (m *TeamMember) IsAdmin() bool {
	return m.Role == ztypes.RoleAdmin
}

func (m *TeamMember) IsExpired() bool {
	return m.ExpiresOn.Before(time.Now())
}
