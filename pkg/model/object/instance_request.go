package object

import "github.com/zbitech/common/pkg/model/ztypes"

type InstanceRequestIF interface {
	GetName() string
	GetVersion() string
	GetInstanceType() ztypes.InstanceType
	GetDataSourceType() ztypes.DataSourceType
	GetDataSource() string
	AllowMethods() bool
}

type InstanceRequest struct {
	Name           string                `json:"name" validate:"required"`
	Version        string                `json:"version" validate:"required"`
	Description    string                `json:"description" validate:"required"`
	Methods        bool                  `json:"allowMethods"`
	DataSourceType ztypes.DataSourceType `json:"dataSourceType"`
	DataSource     string                `json:"dataSource"`
}

func (m InstanceRequest) GetName() string {
	return m.Name
}

func (m InstanceRequest) GetVersion() string {
	return m.Version
}

func (m InstanceRequest) GetInstanceType() ztypes.InstanceType {
	return ztypes.InstanceTypeZCASH
}

func (m InstanceRequest) GetDataSourceType() ztypes.DataSourceType {
	return m.DataSourceType
}

func (m InstanceRequest) GetDataSource() string {
	return m.DataSource
}

func (m InstanceRequest) AllowMethods() bool {
	return m.Methods
}
