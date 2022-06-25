package object

import (
	"github.com/zbitech/common/pkg/model/ztypes"
)

type ProjectRequest struct {
	Network     ztypes.NetworkType `json:"network" validate:"required"`
	Version     string             `json:"version" validate:"required"`
	Name        string             `json:"name" validate:"required"`
	Team        string             `json:"team"`
	Description string             `json:"description" validate:"required"`
	owner       string
}

func (req *ProjectRequest) SetOwner(owner string) {
	req.owner = owner
}

func (req *ProjectRequest) GetOwner() string {
	return req.owner
}
