package entity

import (
	"time"
)

type UserPolicy struct {
	UserId           string                 `json:"userid" bson:"userid"`
	Created          time.Time              `json:"created" bson:"created"`
	Updated          time.Time              `json:"updated" bson:"updated"`
	InstancePolicies []InstanceAccessPolicy `json:"instancePolicies"`
}

func NewUserPolicy(userId string) UserPolicy {
	return UserPolicy{
		UserId:           userId,
		Created:          time.Now(),
		InstancePolicies: make([]InstanceAccessPolicy, 0),
	}
}

func (u *UserPolicy) SetInstanceAccess(project, instance string, allow bool, methods []MethodPolicy) {
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

func (u *UserPolicy) RemoveInstanceAccess(project, instance string) {
	for index, i := range u.InstancePolicies {
		if i.Project == project && i.Instance == instance {
			u.InstancePolicies = append(u.InstancePolicies[:index], u.InstancePolicies[index+1:]...)
			return
		}
	}
}

func (u *UserPolicy) GetInstanceAccess(project, instance string) *InstanceAccessPolicy {

	for _, i := range u.InstancePolicies {
		if i.Project == project && i.Instance == instance {
			return &i
		}
	}
	return nil
}
