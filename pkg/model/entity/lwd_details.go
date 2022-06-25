package entity

type LWDInstanceDetail struct {
	// Project      string              `json:"project"`
	// Name         string              `json:"name"`
	// Version      string              `json:"version"`
	// Network      ztypes.NetworkType  `json:"network"`
	// Description  string              `json:"description"`
	// Owner        string              `json:"owner"`
	// Object       string              `json:"status"`
	// Timestamp    time.Time           `json:"timestamp"`
	// InstanceType ztypes.InstanceType `json:"type"`

	ZcashInstance string `json:"instanceId,omitempty"`
}

// func (z *LWDInstanceDetail) GetName() string {
// 	return z.Name
// }

// func (z *LWDInstanceDetail) GetProject() string {
// 	return z.Project
// }

// func (z *LWDInstanceDetail) GetVersion() string {
// 	return z.Version
// }

// func (z *LWDInstanceDetail) GetNetwork() ztypes.NetworkType {
// 	return z.Network
// }

// func (z *LWDInstanceDetail) GetOwner() string {
// 	return z.Owner
// }

// func (z *LWDInstanceDetail) GetInstanceType() ztypes.InstanceType {
// 	return z.InstanceType
// }
