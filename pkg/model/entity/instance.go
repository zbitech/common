package entity

import (
	"fmt"
	"time"

	"github.com/zbitech/common/pkg/model/ztypes"
)

type DataVolume struct {
	Name   string  `json:"name"`
	Volume string  `json:"volume"`
	Size   float32 `json:"size"`
}

type InstanceIF interface {
	GetName() string
	GetNamespace() string
	GetInstanceType() ztypes.InstanceType
	GetProject() string
	GetVersion() string
	GetOwner() string
	SetStatus(status ztypes.StatusType)
	SetTimestamp(ts time.Time)
}

type Instance struct {
	Project        string                `json:"project"`
	Name           string                `json:"name"`
	Version        string                `json:"version"`
	Network        ztypes.NetworkType    `json:"network"`
	Description    string                `json:"description"`
	Owner          string                `json:"owner"`
	Status         ztypes.StatusType     `json:"status"`
	Timestamp      time.Time             `json:"timestamp"`
	DataSourceType ztypes.DataSourceType `json:"dataSourceType"`
	DataSource     string                `json:"dataSource"`
	InstanceType   ztypes.InstanceType   `json:"type"`
	Action         string                `json:"action"`
	ActionTime     time.Time             `json:"actionTime"`
	Age            string                `json:"age"`
}

type ZcashInstance struct {
	Instance     `json:",inline" bson:",inline"`
	ZcashDetails `json:",inline" bson:"details"`
}

type ZcashDetails struct {
	TransactionIndex bool       `json:"transactionIndex"`
	Miner            bool       `json:"miner"`
	Peers            []string   `json:"peers"`
	DataVolume       DataVolume `json:"dataVolume"`
	ParamsVolume     DataVolume `json:"paramsVolume"`
}

type LWDInstance struct {
	Instance   `json:",inline" bson:",inline"`
	LWDDetails `json:",inline" bson:"details"`
}

type LWDDetails struct {
	ZcashInstance string     `json:"zcashInstance"`
	DataVolume    DataVolume `json:"dataVolume"`
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

func (m Instance) GetStatus() ztypes.StatusType {
	return m.Status
}

func (m *Instance) SetStatus(status ztypes.StatusType) {
	m.Status = status
}

func (m Instance) GetTimestamp() time.Time {
	return m.Timestamp
}

func (m *Instance) SetTimestamp(timestamp time.Time) {
	m.Timestamp = timestamp
	m.Age = time.Now().Sub(timestamp).String()
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
