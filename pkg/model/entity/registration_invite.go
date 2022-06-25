package entity

import (
	"github.com/zbitech/common/pkg/id"
	"github.com/zbitech/common/pkg/model/ztypes"
	"time"
)

type RegistrationInvite struct {
	Key         string
	Email       string
	Role        ztypes.Role
	Status      ztypes.InvitationStatus
	Created     time.Time
	LastUpdate  time.Time
	Level       ztypes.SubscriptionLevel
	EmailStatus struct {
		Attempts        int
		LastAttemptDate time.Time
		Sent            bool
		SentDate        time.Time
	}
}

func NewRegistrationInvite(email string, role ztypes.Role, level ztypes.SubscriptionLevel) RegistrationInvite {
	return RegistrationInvite{
		Key:        id.GenerateKey(),
		Email:      email,
		Role:       role,
		Status:     ztypes.InvitationPending,
		Level:      level,
		Created:    time.Now(),
		LastUpdate: time.Now(),
	}
}

func NewTeamMemberRegistrationInvite(email string, role ztypes.Role) RegistrationInvite {
	return RegistrationInvite{
		Key:        id.GenerateKey(),
		Email:      email,
		Role:       role,
		Status:     ztypes.InvitationPending,
		Level:      ztypes.SubscriptionTeamMember,
		Created:    time.Now(),
		LastUpdate: time.Now(),
	}
}

func (i *RegistrationInvite) Accept() {
	i.Status = ztypes.InvitationAccept
	i.LastUpdate = time.Now()
}

func (i *RegistrationInvite) Reject() {
	i.Status = ztypes.InvitationReject
	i.LastUpdate = time.Now()
}

func (i *RegistrationInvite) Expire() {
	i.Status = ztypes.InvitationExpired
	i.LastUpdate = time.Now()
}

func (i *RegistrationInvite) SetEmailStatus(attemptTime time.Time, success bool) {
	i.EmailStatus.Attempts++
	i.EmailStatus.LastAttemptDate = attemptTime
	if success {
		i.EmailStatus.Sent = true
		i.EmailStatus.SentDate = attemptTime
	} else {
		i.EmailStatus.Sent = false
	}
}
