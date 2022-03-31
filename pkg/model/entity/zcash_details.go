package entity

type ZcashInstanceDetail struct {
	TransactionIndex bool     `json:"transactionIndex"`
	Miner            bool     `json:"miner"`
	Peers            []string `json:"peers"`
}

// func (z *ZcashInstanceDetail) GetName() string {
// 	return z.Name
// }

// func (z *ZcashInstanceDetail) GetProject() string {
// 	return z.Project
// }

// func (z *ZcashInstanceDetail) GetVersion() string {
// 	return z.Version
// }

// func (z *ZcashInstanceDetail) GetNetwork() ztypes.NetworkType {
// 	return z.Network
// }

// func (z *ZcashInstanceDetail) GetOwner() string {
// 	return z.Owner
// }

// func (z *ZcashInstanceDetail) GetInstanceType() ztypes.InstanceType {
// 	return z.InstanceType
// }
