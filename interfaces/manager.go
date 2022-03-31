package interfaces

import (
	"context"

	"github.com/zbitech/common/pkg/model/config"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/object"
	"github.com/zbitech/common/pkg/model/ztypes"
	"go.mongodb.org/mongo-driver/bson"
)

type ResourceManagerFactoryIF interface {
	Init(ctx context.Context) error
	GetProjectDataManager(ctx context.Context) ProjectResourceManagerIF
}

type ProjectResourceManagerIF interface {
	AddInstanceManager(ctx context.Context, i_type ztypes.InstanceType, manager InstanceResourceManagerIF)

	GetProjectResources(ctx context.Context, version string) (config.VersionedResourceConfig, bool)
	ValidateProjectRequest(ctx context.Context, request *object.ProjectRequest) error
	CreateProject(ctx context.Context, request *object.ProjectRequest) (*entity.Project, error)
	CreateProjectSpec(ctx context.Context, project *entity.Project) ([]string, error)
	CreateProjectIngressSpec(ctx context.Context, project *entity.Project, instances []entity.Instance) ([]string, error)

	GetInstanceResources(ctx context.Context, iType ztypes.InstanceType, version string) (*config.VersionedResourceConfig, bool)
	ValidateInstanceRequest(ctx context.Context, request ztypes.InstanceRequestIF) error
	CreateInstance(ctx context.Context, project *entity.Project, request ztypes.InstanceRequestIF) (*entity.Instance, error)
	CreateInstanceSpec(ctx context.Context, project *entity.Project, instance *entity.Instance) ([]string, error)

	UnmarshalBSONInstance(ctx context.Context, data []byte) (*entity.Instance, error)
	UnmarshalBSONKubernetesInstanceResource(ctx context.Context, data []byte) (*entity.KubernetesInstanceResource, error)
	UnmarshalBSONKubernetesProjectResource(ctx context.Context, data []byte) (*entity.KubernetesProjectResource, error)
	UnmarshalBSONDetails(ctx context.Context, iType ztypes.InstanceType, value bson.RawValue) (ztypes.InstanceDetailIF, error)
}

type InstanceResourceManagerIF interface {
	GetInstanceResources(ctx context.Context, version string) (*config.VersionedResourceConfig, bool)
	ValidateInstanceRequest(ctx context.Context, request ztypes.InstanceRequestIF) error
	CreateInstance(ctx context.Context, project *entity.Project, request ztypes.InstanceRequestIF) (*entity.Instance, error)
	CreateInstanceSpec(ctx context.Context, project *entity.Project, instance *entity.Instance) ([]string, error)
	UnmarshalBSONDetails(ctx context.Context, value bson.RawValue) (ztypes.InstanceDetailIF, error)
}
