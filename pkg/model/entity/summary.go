package entity

import "github.com/zbitech/common/pkg/model/ztypes"

type ResourceSummary struct {
	TotalTeams   int `json:"totalTeams"`
	TotalMembers int `json:"totalMembers"`
	TotalAPIKeys int `json:"totalAPIKeys"`
	//	TotalProjects     int `json:"totalProjects"`
	//	TotalVolumeClaims int `json:"totalVolumeClaims"`
	//	TotalInstances    int `json:"totalInstances"`
}

type ProjectSummary struct {
	TotalProjects  int                         `json:"totalProjects"`
	TotalInstances int                         `json:"totalInstances"`
	Instances      map[ztypes.InstanceType]int `json:"instances"`
}

type ResourceObjectSummary struct {
	TotalDeployments     int
	AvailableDeployments int

	Deployments struct {
		Total     int
		Available int
	} `json:"deployments"`
	Pods struct {
		Total   int
		Running int
	} `json:"pods"`
	Services struct {
		Total int
	} `json:"services"`
	ConfigMaps struct {
		Total int
	} `json:"configMaps"`
	Secrets struct {
		Total int
	} `json:"secrets"`
	Volumes struct {
		Total int
	} `json:"volumes"`
	SnapShots struct {
		Total int
	} `json:"snapShots"`
	SnapshotSchedules struct {
		Total int
	} `json:"snapshotSchedules"`
	PlatformIngress struct {
		Total int
	} `json:"platformIngress,omitempty"`
	ProjectIngress struct {
		Total  int
		Routes int
	} `json:"projectIngress,omitempty"`
}
