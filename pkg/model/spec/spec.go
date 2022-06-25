package spec

import (
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/ztypes"
)

type ProjectSpec struct {
	Namespace string
	Labels    map[string]string
	//Data      string
	//DomainName   string
	//DomainSecret string
	//Envoy        EnvoySpec
}

type ControllerIngressSpec struct {
	ControllerDomain  string
	CertSecret        string
	ControllerService string
	ControllerPort    int32
	Projects          []ProjectSpec
	Labels            map[string]string
}

type ProjectIngressSpec struct {
	entity.Project
	Namespace string
	Instances []InstanceIngressSpec
	Labels    map[string]string
}

type InstanceIngressSpec struct {
	Name          string
	Version       string
	ServicePrefix string
	ServicePort   int32
}

type InstanceSpec struct {
	Name               string                `json:"name"`
	Project            string                `json:"project"`
	Version            string                `json:"version"`
	Namespace          string                `json:"namespace"`
	ServiceAccountName string                `json:"serviceAccountName"`
	DataSourceType     ztypes.DataSourceType `json:"dataSourceType"`
	DataSource         string                `json:"dataSource"`
	DomainName         string
	DomainSecret       string
	Labels             map[string]string
}

type EnvoySpec struct {
	Image                 string
	Command               string
	Port                  int32
	Timeout               float32
	AccessAuthorization   bool
	AuthServerURL         string
	AuthServerPort        int32
	AuthenticationEnabled bool
}

type VolumeSource struct {
	Group string `json:"group,omitempty"`
	Kind  string `json:"kind,omitempty"`
	Name  string `json:"name,omitempty"`
}

type ZcashNodeInstanceSpec struct {
	InstanceSpec
	Username          string
	Password          string
	ZcashConf         string
	ZcashImage        string
	MetricsImage      string
	Port              int32
	MetricsPort       int32
	DataVolume        string
	ParamsVolume      string
	ParamVolumeSource *VolumeSource
	DataVolumeSource  *VolumeSource
	Envoy             EnvoySpec
}

type LWDInstanceSpec struct {
	InstanceSpec
	ZcashInstanceName string
	ZcashInstanceUrl  string
	ZcashPort         int32
	LightwalletImage  string
	Port              int32
	HttpPort          int32
	LogLevel          int32
	DataVolume        string
	DataVolumeSource  *VolumeSource
	Envoy             EnvoySpec
	TLSEnabled        bool
}

type VolumeSpec struct {
	Volume             string
	VolumeName         string
	StorageClass       string
	Namespace          string
	SourceName         string
	VolumeDataSource   bool
	SnapshotDataSource bool
	Size               float32
	Labels             map[string]string
}

type SnapshotSpec struct {
	Name string
	//	Project   string
	Namespace string
	//	Version   string
	//	Network          ztypes.NetworkType
	Owner            string
	Volume           string
	BackupExpiration string
	Schedule         string
	SnapshotClass    string
	MaxBackupCount   int
	Labels           map[string]string
}

type SnapshotScheduleSpec struct {
	Name             string
	Namespace        string
	Volume           string
	Schedule         string
	ScheduleType     ztypes.ZBIBackupScheduleType
	SnapshotClass    string
	BackupExpiration string
	MaxBackupCount   int
	Labels           map[string]string
	ClaimLabels      map[string]string
	SnapshotLabels   map[string]string
}
