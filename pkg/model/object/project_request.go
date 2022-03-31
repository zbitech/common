package object

import (
	"github.com/zbitech/common/pkg/model/ztypes"
)

type ProjectRequest struct {
	Network     ztypes.NetworkType `json:"network" validate:"required"`
	Version     string             `json:"version" validate:"required"`
	Name        string             `json:"name" validate:"required"`
	Description string             `json:"description" validate:"required"`
}

// func NewProjectRequest(obj interface{}) (*ProjectRequest, error) {
// 	data, err := json.Marshal(obj)
// 	if err != nil {
// 		return nil, err
// 	}

// 	request := ProjectRequest{}
// 	if err = json.Unmarshal(data, &request); err != nil {
// 		return nil, err
// 	}

// 	return &request, nil
// }
