package entity

import "time"

type InstanceAccessPolicy struct {
	Project  string         `json:"project"`
	Instance string         `json:"instance"`
	Allow    bool           `json:"allow"`
	Created  time.Time      `json:"created"`
	Updated  time.Time      `json:"updated"`
	Methods  []MethodPolicy `json:"methods"`
}

func NewInstanceAccessPolicy(project, instance string, allow bool) InstanceAccessPolicy {
	return InstanceAccessPolicy{
		Project:  project,
		Instance: instance,
		Allow:    allow,
		Created:  time.Now(),
		Updated:  time.Now(),
		Methods:  make([]MethodPolicy, 0),
	}
}

func (a *InstanceAccessPolicy) SetMethods(methods []MethodPolicy) {
	for _, method := range methods {
		a.SetMethod(method.Name, method.Category, method.Allow)
	}
	a.Updated = time.Now()
}

func (a *InstanceAccessPolicy) SetMethod(name, category string, allow bool) {

	for _, m := range a.Methods {
		if m.Name == name && m.Category == category {
			m.Allow = allow
			return
		}
	}

	a.Methods = append(a.Methods, MethodPolicy{Name: name, Category: category, Allow: allow})
	a.Updated = time.Now()
}

func (a *InstanceAccessPolicy) IsMethodAllowed(name string) bool {

	if a.Methods == nil {
		return true
	}

	for _, method := range a.Methods {
		if method.Name == name {
			return method.Allow
		}
	}

	return false
}

func (a *InstanceAccessPolicy) GetMethodsByCategory(category string) []MethodPolicy {

	policies := make([]MethodPolicy, 0)
	for _, method := range a.Methods {
		if method.Category == category {
			policies = append(policies, method)
		}
	}

	return policies
}

func (a *InstanceAccessPolicy) GetMethods(names []string) []MethodPolicy {
	policies := make([]MethodPolicy, len(names))
	for index, name := range names {
		for _, method := range a.Methods {
			if name == method.Name {
				policies[index] = method
			}
		}
	}
	return policies
}

func (a *InstanceAccessPolicy) GetMethodByName(name string) *MethodPolicy {

	for _, method := range a.Methods {
		if method.Name == name {
			return &method
		}
	}

	return nil
}
