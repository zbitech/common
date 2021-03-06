package entity

import (
	"fmt"
	"time"

	"github.com/zbitech/common/pkg/model/ztypes"
)

type Project struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Version     string             `json:"version"`
	Network     ztypes.NetworkType `json:"network"`
	Owner       string             `json:"owner"`
	Status      ztypes.StatusType  `json:"status"`
	Timestamp   time.Time          `json:"timestamp"`
	TeamId      string             `json:"team"`
	Action      string             `json:"action"`
	ActionTime  time.Time          `json:"actionTime"`
	Age         string             `json:"age"`
}

//func (p *Project) AddInstance(inst InstanceSummary) {
//	p.Instances = append(p.Instances, inst)
//}

func (p Project) GetNetwork() ztypes.NetworkType {
	return p.Network
}

func (p Project) GetOwner() string {
	return p.Owner
}

func (p *Project) SetOwner(owner string) {
	p.Owner = owner
}

func (p Project) GetStatus() ztypes.StatusType {
	return p.Status
}

func (p *Project) SetStatus(status ztypes.StatusType) {
	p.Status = status
}

func (p Project) GetTimestamp() time.Time {
	return p.Timestamp
}

func (p *Project) SetTimestamp(timestamp time.Time) {
	p.Timestamp = timestamp
}

func (p Project) GetName() string {
	return p.Name
}

func (p Project) GetVersion() string {
	return p.Version
}

func (p Project) GetNamespace() string {
	return fmt.Sprintf("%s", p.Name)
}
