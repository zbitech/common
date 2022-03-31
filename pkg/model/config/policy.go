package config

import "github.com/zbitech/common/pkg/model/ztypes"

type ResourcePolicy struct {
	MaxStorage string `json:"maxStorage,omitempty"`
	MaxCPU     string `json:"maxCPU,omitempty"`
	MaxMemory  string `json:"maxMemory,omitempty"`
}

type GlobalPolicy struct {
	ImageRegistries      []string             `json:"imageRegistries"`
	StorageClasses       []string             `json:"storageClasses"`
	Domain               string               `json:"domain"`
	CertName             string               `json:"certName"`
	ServiceAccount       string               `json:"serviceAccount"`
	SubscriptionPolicies []SubscriptionPolicy `json:"subscriptionPolicies"`
}

type SubscriptionPolicy struct {
	Level           ztypes.SubscriptionLevel `json:"level" bson:"level"`
	MaxKeys         int                      `json:"maxKeys"`
	MaxProjects     int                      `json:"maxProjects"`
	MaxInstances    int                      `json:"maxInstances"`
	MaxTeams        int                      `json:"maxTeams"`
	MaxAPIKeys      int                      `json:"maxAPIKeys"`
	MaxVolumeClaims int                      `json:"maxVolumeClaims"`
	ResourceLimit   ResourcePolicy           `json:"resourceLimit"`
	InstanceTypes   []ztypes.InstanceType    `json:"instanceTypes"`
}

func (s *SubscriptionPolicy) IsInstanceTypeAllowed(i_type ztypes.InstanceType) bool {
	for _, _type := range s.InstanceTypes {
		if _type == i_type {
			return true
		}
	}
	return false
}

func (g *GlobalPolicy) SetSubscriptionPolicy(p SubscriptionPolicy) {
	for _, sp := range g.SubscriptionPolicies {
		if sp.Level == p.Level {
			sp.InstanceTypes = p.InstanceTypes
			sp.MaxInstances = p.MaxInstances
			sp.MaxKeys = p.MaxKeys
			sp.MaxProjects = p.MaxProjects
			sp.MaxVolumeClaims = p.MaxVolumeClaims
			sp.ResourceLimit = p.ResourceLimit

			return
		}
	}

	g.SubscriptionPolicies = append(g.SubscriptionPolicies, p)
}

func (g *GlobalPolicy) GetSubscriptionPolicy(level ztypes.SubscriptionLevel) *SubscriptionPolicy {

	for _, sp := range g.SubscriptionPolicies {
		if sp.Level == level {
			return &sp
		}
	}

	return nil
}
