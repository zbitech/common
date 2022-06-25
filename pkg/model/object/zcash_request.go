package object

import (
	"github.com/zbitech/common/pkg/model/ztypes"
)

type ZcashNodeInstanceRequest struct {
	InstanceRequest
	TransactionIndex bool     `json:"transaction"`
	Miner            bool     `json:"miner"`
	Peers            []string `json:"peers"`
}

func (m ZcashNodeInstanceRequest) GetInstanceType() ztypes.InstanceType {
	return ztypes.InstanceTypeZCASH
}
