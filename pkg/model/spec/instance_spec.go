package spec

import (
	"time"

	"github.com/zbitech/common/pkg/model/ztypes"
)

type InstanceSpec struct {
	Project            string              `json:"project,omitempty"`
	Namespace          string              `json:"namespace,omitempty"`
	Name               string              `json:"name,omitempty"`
	Version            string              `json:"version,omitempty"`
	Network            ztypes.NetworkType  `json:"network,omitempty"`
	Owner              string              `json:"owner,omitempty"`
	Status             string              `json:"status,omitempty"`
	Timestamp          time.Time           `json:"timestamp,omitempty"`
	InstanceType       ztypes.InstanceType `json:"type,omitempty"`
	ServiceAccountName string              `json:"serviceAccountName,omitempty"`
	StorageClass       string              `json:"storageClass,omitempty"`
	Timeout            string              `json:"timeout,omitempty"`
}

func (m InstanceSpec) GetInstance() ztypes.InstanceIF {
	return nil
}
