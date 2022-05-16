package interfaces

import (
	"context"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/zbitech/common/pkg/model/config"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/object"
	"github.com/zbitech/common/pkg/model/ztypes"
	"go.mongodb.org/mongo-driver/bson"
)

type ResourceManagerFactoryIF interface {
	Init(ctx context.Context) error
	GetIngressResourceManager(ctx context.Context) IngressResourceManagerIF
	GetProjectDataManager(ctx context.Context) ProjectResourceManagerIF
}

type IngressResourceManagerIF interface {
	GetIngressResources(version string) (config.VersionedResourceConfig, bool)
	CreateProjectIngressSpec(ctx context.Context, obj *unstructured.Unstructured, version string, project *entity.Project, instance *entity.Instance) (*unstructured.Unstructured, error)
	CreateControllerIngressSpec(ctx context.Context, obj *unstructured.Unstructured, version string, project *entity.Project) (*unstructured.Unstructured, error)
}

type ProjectResourceManagerIF interface {
	GetProjectResources(version string) (config.VersionedResourceConfig, bool)
	ValidateProjectRequest(ctx context.Context, request *object.ProjectRequest) error
	CreateProject(ctx context.Context, request *object.ProjectRequest) (*entity.Project, error)
	CreateProjectAssets(ctx context.Context, project *entity.Project) ([]*unstructured.Unstructured, error)

	GetInstanceResources(iType ztypes.InstanceType, version string) (*config.VersionedResourceConfig, bool)
	ValidateInstanceRequest(ctx context.Context, request ztypes.InstanceRequestIF) error
	CreateInstance(ctx context.Context, project *entity.Project, request ztypes.InstanceRequestIF) (*entity.Instance, error)
	CreateInstanceAssets(ctx context.Context, project *entity.Project, instance *entity.Instance) ([]*unstructured.Unstructured, error)

	UnmarshalBSONInstance(ctx context.Context, data []byte) (*entity.Instance, error)
	UnmarshalBSONDetails(ctx context.Context, iType ztypes.InstanceType, value bson.RawValue) (ztypes.InstanceDetailIF, error)
}

type InstanceResourceManagerIF interface {
	GetInstanceResources(version string) (*config.VersionedResourceConfig, bool)
	ValidateInstanceRequest(ctx context.Context, request ztypes.InstanceRequestIF) error
	CreateInstance(ctx context.Context, project *entity.Project, request ztypes.InstanceRequestIF) (*entity.Instance, error)
	CreateInstanceAssets(ctx context.Context, project *entity.Project, instance *entity.Instance) ([]*unstructured.Unstructured, error)
	UnmarshalBSONDetails(ctx context.Context, value bson.RawValue) (ztypes.InstanceDetailIF, error)
}
