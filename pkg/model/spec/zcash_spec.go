package spec

type VolumeSource struct {
	Group string `json:"group,omitempty"`
	Kind  string `json:"kind,omitempty"`
	Name  string `json:"name,omitempty"`
}

type ZcashNodeInstanceSpec struct {
	InstanceSpec
	Username          string
	Password          string
	Authentication    string
	ZcashConf         string
	ZcashImage        string
	MetricsImage      string
	Port              int32
	MetricsPort       int32
	ParamVolumeSource *VolumeSource
	DataVolumeSource  *VolumeSource
}
