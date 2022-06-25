package entity

import (
	"github.com/zbitech/common/pkg/model/ztypes"
)

type KubernetesResource struct {
	Name      string                    `json:"name,omitempty"`
	Namespace string                    `json:"namespace,omitempty"`
	Type      ztypes.ResourceObjectType `json:"type,omitempty"`
	Status    string                    `json:"status,omitempty"`
	//	Timestamp time.Time                 `json:"timestamp,omitempty"`
	Age string `json:"age,omitempty"`
}

//type Condition struct {
//	Type           string
//	TransitionTime time.Time
//}

//type ResourceState struct {
//	Active  bool
//	Deleted bool
//}

//func (r *ResourceState) IsActive() bool {
//	return r.Active
//}

//func (r *ResourceState) IsDeleted() bool {
//	return r.Deleted
//}

//type PersistentVolumeState struct {
//	Phase string
//}

//type PersistentVolumeClaimState struct {
//	Phase string
//}
