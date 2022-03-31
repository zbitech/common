package entity

import (
	"github.com/zbitech/common/pkg/id"
	//	"github.com/zbi/common/pkg/vars"
	"time"
)

type APIKey struct {
	Key     string
	UserId  string
	Created time.Time
	Expires time.Time
}

func NewAPIKey(userid string, expire int) APIKey {
	return APIKey{
		Key:     id.GenerateAPIKey(),
		UserId:  userid,
		Created: time.Now(),
		Expires: time.Now().Add(time.Duration(expire) * time.Hour),
	}
}

func (a APIKey) IsExpired() bool {
	return time.Now().After(a.Expires)
}

type APIKeyPolicy struct {
	Key              string                 `json:"key" bson:"key"`
	Allow            bool                   `json:"allow" bson:"allow"`
	Created          time.Time              `json:"created" bson:"created"`
	Updated          time.Time              `json:"updated" bson:"updated"`
	Addresses        []string               `json:"addresses,omitempty"`
	InstancePolicies []InstanceAccessPolicy `json:"instancePolicies,omitempty"`
}

func NewAPIKeyPolicy(key string, allow bool) APIKeyPolicy {
	return APIKeyPolicy{
		Key:              key,
		Allow:            allow,
		Created:          time.Now(),
		Updated:          time.Now(),
		InstancePolicies: make([]InstanceAccessPolicy, 0),
		Addresses:        nil,
	}
}

func (u *APIKeyPolicy) SetInstanceAccess(project, instance string, allow bool, methods []MethodPolicy) {
	for _, i := range u.InstancePolicies {
		if i.Project == project && i.Instance == instance {
			i.Allow = allow
			i.SetMethods(methods)
		}
	}

	p := NewInstanceAccessPolicy(project, instance, allow)
	p.SetMethods(methods)
	u.InstancePolicies = append(u.InstancePolicies, p)
}

func (u *APIKeyPolicy) RemoveInstanceAccess(project, instance string) {
	for index, i := range u.InstancePolicies {
		if i.Project == project && i.Instance == instance {
			u.InstancePolicies = append(u.InstancePolicies[:index], u.InstancePolicies[index+1:]...)
			return
		}
	}
}

func (u *APIKeyPolicy) GetInstanceAccess(project, instance string) *InstanceAccessPolicy {

	for _, i := range u.InstancePolicies {
		if i.Project == project && i.Instance == instance {
			return &i
		}
	}
	return nil
}

func (a *APIKeyPolicy) IsAddressAllowed(addr string) bool {
	if a.Addresses == nil {
		return true
	}

	for _, address := range a.Addresses {
		if address == addr {
			return true
		}
	}

	return false
}
