package spec

type LWDInstanceSpec struct {
	InstanceSpec
	ZcashInstance    string
	LightwalletImage string
	Port             int32
	HttPort          int32
	LogLevel         int32
	DataVolumeSource *VolumeSource
}
