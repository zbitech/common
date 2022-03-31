package entity

import (
	"fmt"
	"time"

	"github.com/zbitech/common/pkg/model/ztypes"
)

type Instance struct {
	Project        string                  `json:"project"`
	Name           string                  `json:"name"`
	Version        string                  `json:"version"`
	Network        ztypes.NetworkType      `json:"network"`
	Description    string                  `json:"description"`
	Owner          string                  `json:"owner"`
	Status         string                  `json:"status"`
	Timestamp      time.Time               `json:"timestamp"`
	InstanceType   ztypes.InstanceType     `json:"type"`
	InstanceDetail ztypes.InstanceDetailIF `json:"details,omitempty,inline"`
}

func (m Instance) GetNetwork() ztypes.NetworkType {
	return m.Network
}

func (m Instance) GetOwner() string {
	return m.Owner
}

func (m *Instance) SetOwner(owner string) {
	m.Owner = owner
}

func (m Instance) GetStatus() string {
	return m.Status
}

func (m *Instance) SetStatus(status string) {
	m.Status = status
}

func (m Instance) GetTimestamp() time.Time {
	return m.Timestamp
}

func (m *Instance) SetTimestamp(timestamp time.Time) {
	m.Timestamp = timestamp
}

func (m Instance) GetName() string {
	return m.Name
}

func (m Instance) GetInstanceType() ztypes.InstanceType {
	return m.InstanceType
}

func (m Instance) GetProject() string {
	return m.Project
}

func (m *Instance) SetProject(name string) {
	m.Project = name
}

func (m Instance) GetVersion() string {
	return m.Version
}

func (m Instance) GetNamespace() string {
	return fmt.Sprintf("%s", m.Project)
}
