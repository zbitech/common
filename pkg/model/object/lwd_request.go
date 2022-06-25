package object

import (
	"github.com/zbitech/common/pkg/model/ztypes"
)

type LWDInstanceRequest struct {
	InstanceRequest
	ZcashInstance string `json:"zcashInstance"`
}

func (m LWDInstanceRequest) GetInstanceType() ztypes.InstanceType {
	return ztypes.InstanceTypeLWD
}
