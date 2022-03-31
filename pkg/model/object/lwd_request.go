package object

import (
	"encoding/json"

	"github.com/zbitech/common/pkg/model/ztypes"
)

type LWDInstanceRequest struct {
	Name          string `json:"name"`
	Version       string `json:"version"`
	ZcashInstance string `json:"zcashInstance"`
	Methods       bool   `json:"allowMethods"`
}

func NewLWDInstanceRequest(obj interface{}) (*LWDInstanceRequest, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	request := LWDInstanceRequest{}
	if err = json.Unmarshal(data, &request); err != nil {
		return nil, err
	}

	return &request, nil
}

func (m LWDInstanceRequest) GetName() string {
	return m.Name
}

func (m LWDInstanceRequest) GetVersion() string {
	return m.Version
}

func (m LWDInstanceRequest) GetInstanceType() ztypes.InstanceType {
	return ztypes.LWD_INSTANCE
}

func (m LWDInstanceRequest) AllowMethods() bool {
	return m.Methods
}
