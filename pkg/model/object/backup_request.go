package object

import "github.com/zbitech/common/pkg/model/ztypes"

type SnapshotRequest struct {
	Version    string `json:"version" validate:"required"`
	Volume     string `json:"volume" validate:"required"`
	VolumeName string `json:"volumeName"`
	Namespace  string `json:"namespace"`
	Labels     map[string]string
}

type SnapshotScheduleRequest struct {
	Version    string                       `json:"version" validate:"required"`
	Schedule   ztypes.ZBIBackupScheduleType `json:"schedule" validate:"required"`
	Volume     string                       `json:"volume" validate:"required"`
	VolumeName string                       `json:"volumeName"`
	Namespace  string                       `json:"namespace"`
	Labels     map[string]string
}
