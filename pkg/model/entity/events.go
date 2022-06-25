package entity

import (
	"github.com/zbitech/common/pkg/model/ztypes"
	"time"
)

type Event struct {
	Action    ztypes.EventAction `json:"action,omitempty"`
	Actor     string             `json:"actor,omitempty"`
	Timestamp time.Time          `json:"timestamp,omitempty"`
	Success   bool               `json:"success"`
	Reason    string             `json:"reason,omitempty"`
}

type ProjectEvent struct {
	Project string `json:"project"`
	Event   `json:",inline" bson:",inline"`
}

type InstanceEvent struct {
	Project  string `json:"project"`
	Instance string `json:"instance"`
	Event    `json:",inline" bson:",inline"`
}

type ProfileEvent struct {
	UserId string `json:"userid"`
	Event  `json:",inline" bson:",inline"`
}

type TeamEvent struct {
	TeamId  string   `json:"teamid"`
	Targets []string `json:"targets,omitempty"`
	Event   `json:",inline" bson:",inline"`
}

func NewEvent(actor string, eventAction ztypes.EventAction, evtErr error) Event {
	var reason string
	var success bool

	if evtErr != nil {
		reason = evtErr.Error()
		success = false
	} else {
		success = true
	}

	return Event{
		Action:    eventAction,
		Actor:     actor,
		Timestamp: time.Now(),
		Success:   success,
		Reason:    reason,
	}
}

func NewProjectEvent(project, actor string, eventAction ztypes.EventAction, evtErr error) ProjectEvent {
	return ProjectEvent{
		Project: project,
		Event:   NewEvent(actor, eventAction, evtErr),
	}
}

func NewInstanceEvent(project, instance, actor string, eventAction ztypes.EventAction, evtErr error) InstanceEvent {
	return InstanceEvent{
		Project:  project,
		Instance: instance,
		Event:    NewEvent(actor, eventAction, evtErr),
	}
}

func NewProfileEvent(userId, actor string, eventAction ztypes.EventAction, evtErr error) ProfileEvent {
	return ProfileEvent{
		UserId: userId,
		Event:  NewEvent(actor, eventAction, evtErr),
	}
}

func NewTeamEvent(teamId, actor string, targets []string, eventAction ztypes.EventAction, evtErr error) TeamEvent {
	return TeamEvent{
		TeamId:  teamId,
		Targets: targets,
		Event:   NewEvent(actor, eventAction, evtErr),
	}
}
