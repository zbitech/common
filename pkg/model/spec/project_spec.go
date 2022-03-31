package spec

import "github.com/zbitech/common/pkg/model/entity"

type ProjectSpec struct {
	entity.Project
	ServiceAccountName string            `json:"serviceAccountName,omitempty"`
	Namespace          string            `json:"namespace,omitempty"`
	AuthzServerImage   string            `json:"authzServerImage,omitempty"`
	Domain             string            `json:"domain"`
	CertName           string            `json:"certName"`
	Instances          []entity.Instance `json:"instances"`
}
