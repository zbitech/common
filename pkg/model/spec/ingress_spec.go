package spec

import "github.com/zbitech/common/pkg/model/entity"

type ControllerIngressSpec struct {
	ControllerDomain  string
	CertSecret        string
	ControllerService string
	ControllerPort    int32
	Projects          []entity.Project
}

type ProjectIngressSpec struct {
	entity.Project
	Instances []InstanceIngressSpec
}

type InstanceIngressSpec struct {
	Name          string
	Version       string
	ServicePrefix string
	ServicePort   int32
}
