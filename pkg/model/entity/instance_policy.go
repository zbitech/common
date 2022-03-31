package entity

type InstancePolicy struct {
	Project  string         `json:"project"`
	Instance string         `json:"instance"`
	Methods  []MethodPolicy `json:"methods"`
	Allow    bool           `json:"allow"`
}

func NewInstancePolicy(project, instance string, method_map map[string][]string, allow bool) InstancePolicy {

	i_policy := InstancePolicy{
		Project:  project,
		Instance: instance,
		Methods:  nil,
		Allow:    allow,
	}

	if method_map != nil {
		i_policy.SetMethodMap(method_map, allow)
	}

	return i_policy
}

func (i_p *InstancePolicy) IsMethodAllowed(methodName string) bool {
	if i_p.Methods == nil {
		return i_p.Allow
	}

	for _, method := range i_p.Methods {
		if method.Name == methodName {
			return method.Allow
		}
	}

	return false
}

func (i_p *InstancePolicy) SetMethods(methods []MethodPolicy) {
	for _, method := range methods {
		i_p.SetMethod(method.Name, method.Category, method.Allow)
	}
	//	a.Updated = time.Now()
}

func (i_p *InstancePolicy) SetMethodMap(method_map map[string][]string, allow bool) {
	if method_map != nil {
		i_p.Methods = make([]MethodPolicy, 0)
		for cat, methods := range method_map {
			for _, method := range methods {
				i_p.Methods = append(i_p.Methods, MethodPolicy{Name: method, Category: cat, Allow: allow})
			}
		}
	}
}

func (i_p *InstancePolicy) SetMethod(method, category string, allow bool) {
	for _, m := range i_p.Methods {
		if m.Category == category && m.Name == method {
			m.Allow = allow
			return
		}
	}

	i_p.Methods = append(i_p.Methods, MethodPolicy{Name: method, Category: category, Allow: allow})
}

func (i_p *InstancePolicy) GetMethodsByCategory(category string) []MethodPolicy {

	policies := make([]MethodPolicy, 0)
	for _, method := range i_p.Methods {
		if method.Category == category {
			policies = append(policies, method)
		}
	}

	return policies
}

func (i_p *InstancePolicy) GetMethods(names []string) []MethodPolicy {
	policies := make([]MethodPolicy, len(names))
	for index, name := range names {
		for _, method := range i_p.Methods {
			if name == method.Name {
				policies[index] = method
			}
		}
	}
	return policies
}

func (i_p *InstancePolicy) GetMethodByName(name string) *MethodPolicy {

	for _, method := range i_p.Methods {
		if method.Name == name {
			return &method
		}
	}

	return nil
}
