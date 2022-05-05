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
		Status: ztypes.NEW_INVITATION, CreatedOn: time.Now(), ExpiresOn: time.Now().Add(time.Hour * 168)}

	if admin {
		mbr.Role = ztypes.ADMIN_ROLE
	} else {
		mbr.Role = ztypes.USER_ROLE
	}

	return mbr
}

func (m *TeamMember) SetAdmin(admin bool) {
	if admin {
		m.Role = ztypes.ADMIN_ROLE
	} else {
		m.Role = ztypes.USER_ROLE
	}
}

func (m *TeamMember) Accept() {
	m.Status = ztypes.ACCEPT_INVITATION
	m.LastUpdate = time.Now()
}

func (m *TeamMember) Reject() {
	m.Status = ztypes.REJECT_INVITATION
	m.LastUpdate = time.Now()
}

func (m *TeamMember) Expire() {
	m.Status = ztypes.EXPIRED_INVITATION
	m.LastUpdate = time.Now()
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
