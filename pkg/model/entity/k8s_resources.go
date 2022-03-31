package entity

import (
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

type KubernetesResource struct {
	Id        string                       `json:"id" bson:"_id"`
	Name      string                       `json:"name,omitempty"`
	Namespace string                       `json:"namespace,omitempty"`
	Type      string                       `json:"type,omitempty"`
	GVR       *schema.GroupVersionResource `json:"gvr,omitempty"`
	State     interface{}                  `json:"state,omitempty"`
	Timestamp time.Time                    `json:"timestamp,omitempty"`
}

type KubernetesProjectResource struct {
	KubernetesResource
	Project string `json:"project"`
}

type KubernetesInstanceResource struct {
	KubernetesResource
	Project  string `json:"project"`
	Instance string `json:"instance"`
}

type Condition struct {
	Type           string
	TransitionTime time.Time
}

type ResourceState struct {
	Active  bool
	Deleted bool
}

func (r *ResourceState) IsActive() bool {
	return r.Active
}

func (r *ResourceState) IsDeleted() bool {
	return r.Deleted
}

type DeploymentState struct {
	Available           bool
	Deleted             bool
	Conditions          []string
	Replicas            int32
	AvailableReplicas   int32
	ReadyReplicas       int32
	UpdatedReplicas     int32
	UnavailableReplicas int32
}

func (r *DeploymentState) IsActive() bool {
	return r.Available
}

func (r *DeploymentState) IsDeleted() bool {
	return r.Deleted
}

type PodState struct {
	Conditions        []string
	ContainerStatuses []ContainerState
	Phase             string
	PodReady          bool
	ContainersReady   bool
	Deleted           bool
}

func (r *PodState) IsActive() bool {
	return r.Phase == "Running" && r.PodReady && r.ContainersReady
}

func (r *PodState) IsDeleted() bool {
	return r.Deleted
}

type ContainerState struct {
	Name         string
	Ready        bool
	Started      bool
	Terminated   bool
	RestartCount int32
	//	State        string
}

type PersistentVolumeState struct {
	Phase string
}

type PersistentVolumeClaimState struct {
	Phase string
}
