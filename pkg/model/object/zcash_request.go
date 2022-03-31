package object

import (
	"encoding/json"

	"github.com/zbitech/common/pkg/model/ztypes"
)

type ZcashNodeInstanceRequest struct {
	Name             string   `json:"name" validate:"required"`
	Version          string   `json:"version" validate:"required"`
	Description      string   `json:"description" validate:"required"`
	TransactionIndex bool     `json:"transaction"`
	Miner            bool     `json:"miner"`
	Peers            []string `json:"peers"`
	Methods          bool     `json:"allowMethods"`
}

func NewZcashInstanceRequest(obj interface{}) (*ZcashNodeInstanceRequest, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	request := ZcashNodeInstanceRequest{}
	if err = json.Unmarshal(data, &request); err != nil {
		return nil, err
	}

	return &request, nil
}

// func NewZcashInstanceRequest(name, version string, txIndex, miner, allowMethods bool, peers []string) ZcashNodeInstanceRequest {
// 	return ZcashNodeInstanceRequest{Name: name, Version: version, TransactionIndex: txIndex, Miner: miner, Methods: allowMethods, Peers: peers}
// }

func (m ZcashNodeInstanceRequest) GetName() string {
	return m.Name
}

func (m ZcashNodeInstanceRequest) GetVersion() string {
	return m.Version
}

func (m ZcashNodeInstanceRequest) GetInstanceType() ztypes.InstanceType {
	return ztypes.ZCASH_INSTANCE
}

func (m ZcashNodeInstanceRequest) AllowMethods() bool {
	return m.Methods
}
