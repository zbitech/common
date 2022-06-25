package interfaces

import (
	"context"
	"github.com/zbitech/common/pkg/model/spec"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/zbitech/common/pkg/model/config"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/object"
	"github.com/zbitech/common/pkg/model/ztypes"
	"go.mongodb.org/mongo-driver/bson"
)

type ResourceManagerFactoryIF interface {
	Init(ctx context.Context) error
	GetAppResourceManager(ctx context.Context) AppResourceManagerIF
	GetProjectDataManager(ctx context.Context) ProjectResourceManagerIF
}

type AppResourceManagerIF interface {
	GetAppResources(version string) (*config.VersionedResourceConfig, bool)
	//	CreateControllerIngressAsset(ctx context.Context, obj *unstructured.Unstructured, project *entity.Project, action string) (*unstructured.Unstructured, error)
	//	CreateProjectIngressAsset(ctx context.Context, obj *unstructured.Unstructured, project *entity.Project, instance entity.InstanceIF, action string) (*unstructured.Unstructured, error)
	CreateSnapshotAsset(ctx context.Context, req *object.SnapshotRequest) ([]*unstructured.Unstructured, error)
	CreateSnapshotScheduleAsset(ctx context.Context, req *object.SnapshotScheduleRequest) ([]*unstructured.Unstructured, error)
	CreateVolumeAsset(ctx context.Context, volumes ...spec.VolumeSpec) ([]*unstructured.Unstructured, error)
}

type ProjectResourceManagerIF interface {
	GetProjectResources(version string) (*config.VersionedResourceConfig, bool)
	CreateProject(ctx context.Context, request *object.ProjectRequest) (*entity.Project, error)
	UpdateProject(ctx context.Context, project *entity.Project, request *object.ProjectRequest) error
	CreateProjectAssets(ctx context.Context, project *entity.Project) ([]*unstructured.Unstructured, error)
	CreateProjectIngressAsset(ctx context.Context, appIngress *unstructured.Unstructured, project *entity.Project, action ztypes.EventAction) ([]*unstructured.Unstructured, error)

	CreateInstanceRequest(ctx context.Context, iType ztypes.InstanceType, iRequest interface{}) (object.InstanceRequestIF, error)
	GetInstanceResources(iType ztypes.InstanceType, version string) (*config.VersionedResourceConfig, bool)
	CreateInstance(ctx context.Context, project *entity.Project, request object.InstanceRequestIF) (entity.InstanceIF, error)
	UpdateInstance(ctx context.Context, project *entity.Project, instance entity.InstanceIF, request object.InstanceRequestIF) error

	CreateDeploymentResourceAssets(ctx context.Context, instance entity.InstanceIF) ([]*unstructured.Unstructured, error)
	CreateStartResourceAssets(ctx context.Context, instance entity.InstanceIF) ([]*unstructured.Unstructured, error)
	CreateIngressAsset(ctx context.Context, projIngress *unstructured.Unstructured, instance entity.InstanceIF, action ztypes.EventAction) (*unstructured.Unstructured, error)
	CreateSnapshotAssets(ctx context.Context, instance entity.InstanceIF, volume string) ([]*unstructured.Unstructured, error)
	CreateSnapshotScheduleAssets(ctx context.Context, instance entity.InstanceIF, volume string, schedule ztypes.ZBIBackupScheduleType) ([]*unstructured.Unstructured, error)
	CreateRotationAssets(ctx context.Context, instance entity.InstanceIF) ([]*unstructured.Unstructured, error)

	UnmarshalBSONInstance(ctx context.Context, data bson.Raw) (entity.InstanceIF, error)
	UnmarshalBSONDetails(ctx context.Context, iType ztypes.InstanceType, value bson.Raw) (entity.InstanceIF, error)
}

type InstanceResourceManagerIF interface {
	CreateInstanceRequest(ctx context.Context, iRequest interface{}) (object.InstanceRequestIF, error)
	GetInstanceResources(version string) (*config.VersionedResourceConfig, bool)

	CreateInstance(ctx context.Context, project *entity.Project, request object.InstanceRequestIF) (entity.InstanceIF, error)
	UpdateInstance(ctx context.Context, project *entity.Project, instance entity.InstanceIF, request object.InstanceRequestIF) error

	CreateDeploymentResourceAssets(ctx context.Context, instance entity.InstanceIF) ([]*unstructured.Unstructured, error)
	CreateIngressAsset(ctx context.Context, projIngress *unstructured.Unstructured, instance entity.InstanceIF, action ztypes.EventAction) (*unstructured.Unstructured, error)
	CreateStartResourceAssets(ctx context.Context, instance entity.InstanceIF) ([]*unstructured.Unstructured, error)
	CreateSnapshotAssets(ctx context.Context, instance entity.InstanceIF, volume string) ([]*unstructured.Unstructured, error)
	CreateSnapshotScheduleAssets(ctx context.Context, instance entity.InstanceIF, volume string, scheduleType ztypes.ZBIBackupScheduleType) ([]*unstructured.Unstructured, error)
	CreateRotationAssets(ctx context.Context, instance entity.InstanceIF) ([]*unstructured.Unstructured, error)

	UnmarshalBSONDetails(ctx context.Context, value bson.Raw) (entity.InstanceIF, error)
}
