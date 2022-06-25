package interfaces

import (
	"context"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/k8s"
	"github.com/zbitech/common/pkg/model/object"
	"github.com/zbitech/common/pkg/model/ztypes"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type KlientFactoryIF interface {
	Init(ctx context.Context, namespace string, rtypes ...ztypes.ResourceObjectType) error
	GetZBIClient() ZBIClientIF
	GetKubernesClient() KlientIF
	StartInformer(ctx context.Context) error
}

type KlientInformerControllerIF interface {
	AddInformer(ctx context.Context, rType ztypes.ResourceObjectType)
	Run(ctx context.Context)
	Close()
}

type ZBIClientIF interface {
	RunInformer(ctx context.Context)

	//	CreateIngress(ctx context.Context, project *entity.Project, instance entity.InstanceIF) ([]entity.KubernetesResource, error)
	//	UpdateControllerIngress(ctx context.Context, project *entity.Project, action string) ([]entity.KubernetesResource, error)
	//	UpdateProjectIngress(ctx context.Context, project *entity.Project, instance entity.InstanceIF, action string) ([]entity.KubernetesResource, error)
	GetStorageClasses(ctx context.Context) []k8s.StorageClass
	GetSnapshotClasses(ctx context.Context) []k8s.SnapshotClass

	ValidateProjectRequest(ctx context.Context, request *object.ProjectRequest) (bool, map[string]string)
	CreateProject(ctx context.Context, request *object.ProjectRequest) (*entity.Project, error)
	CreateProjectResources(ctx context.Context, project *entity.Project) ([]entity.KubernetesResource, error)

	UpdateProject(ctx context.Context, project *entity.Project, request *object.ProjectRequest) error
	DeleteProject(ctx context.Context, name string) error
	GetAllProjects(ctx context.Context) ([]entity.Project, error)
	GetProject(ctx context.Context, name string) (*entity.Project, error)
	GetProjectsByOwner(ctx context.Context, owner string) ([]entity.Project, error)
	GetProjectsByTeam(ctx context.Context, team string) ([]entity.Project, error)
	GetProjectResources(ctx context.Context, project *entity.Project) ([]entity.KubernetesResource, error)

	ValidateInstanceRequest(ctx context.Context, project *entity.Project, request object.InstanceRequestIF) (bool, map[string]string)
	CreateInstance(ctx context.Context, project *entity.Project, request object.InstanceRequestIF) (entity.InstanceIF, error)
	UpdateInstance(ctx context.Context, project *entity.Project, instance entity.InstanceIF, request object.InstanceRequestIF) error
	DeleteInstance(ctx context.Context, project *entity.Project, instance entity.InstanceIF) error

	CreateInstanceResources(ctx context.Context, project *entity.Project, instance entity.InstanceIF) ([]entity.KubernetesResource, error)
	StopInstance(ctx context.Context, project *entity.Project, instance entity.InstanceIF) error
	StartInstance(ctx context.Context, project *entity.Project, instance entity.InstanceIF) ([]entity.KubernetesResource, error)
	CreateInstanceBackup(ctx context.Context, instance entity.InstanceIF, req *object.SnapshotRequest) (*entity.KubernetesResource, error)
	ScheduleInstanceBackup(ctx context.Context, instance entity.InstanceIF, req *object.SnapshotScheduleRequest) (*entity.KubernetesResource, error)
	RotateInstanceCredentials(ctx context.Context, project *entity.Project, instance entity.InstanceIF) ([]entity.KubernetesResource, error)

	GetAllInstances(ctx context.Context) ([]entity.InstanceIF, error)
	GetInstancesByProject(ctx context.Context, project string) ([]entity.InstanceIF, error)
	GetInstancesByOwner(ctx context.Context, owner string) ([]entity.InstanceIF, error)
	GetInstanceByName(ctx context.Context, project, instance string) (entity.InstanceIF, error)

	GetInstanceResources(ctx context.Context, project *entity.Project, instance entity.InstanceIF) ([]entity.KubernetesResource, error)
	GetInstanceVolumes(ctx context.Context, instance entity.InstanceIF) []k8s.InstanceVolume
	GetInstanceSnapshots(ctx context.Context, instance entity.InstanceIF) []k8s.Snapshot
	GetInstanceSchedules(ctx context.Context, instance entity.InstanceIF) []k8s.SnapshotSchedule

	GetResourceSummary(ctx context.Context) map[string]interface{}
	GetProjectResourceSummary(ctx context.Context, project *entity.Project, extraLabels map[string]string) map[string]interface{}
	GetInstanceResourceSummary(ctx context.Context, instance entity.InstanceIF, extraLabels map[string]string) map[string]interface{}
}

type KlientIF interface {
	DeleteDynamicResource(ctx context.Context, namespace, name string, resource schema.GroupVersionResource) error
	DeleteNamespace(ctx context.Context, namespace string) error

	ApplyResource(ctx context.Context, object *unstructured.Unstructured) (*entity.KubernetesResource, error)
	ApplyResources(ctx context.Context, objects []*unstructured.Unstructured) ([]entity.KubernetesResource, error)

	GetDynamicResource(ctx context.Context, namespace, name string, resource schema.GroupVersionResource) (*unstructured.Unstructured, error)
	GetDynamicResourceList(ctx context.Context, namespace string, resource schema.GroupVersionResource) ([]unstructured.Unstructured, error)

	GetNamespace(ctx context.Context, name string) (*corev1.Namespace, error)
	GetNamespaces(ctx context.Context, labels map[string]string) ([]corev1.Namespace, error)

	GetStorageClass(ctx context.Context, name string) (*storagev1.StorageClass, error)
	GetStorageClasses(ctx context.Context) ([]storagev1.StorageClass, error)

	GetSnapshotClass(ctx context.Context, name string) (*unstructured.Unstructured, error)
	GetSnapshotClasses(ctx context.Context) ([]unstructured.Unstructured, error)

	GetDeploymentByName(ctx context.Context, namespace, name string) (*appsv1.Deployment, error)
	GetDeployments(ctx context.Context, namespace string, labels map[string]string) []appsv1.Deployment

	GetPodByName(ctx context.Context, namespace, name string) (*corev1.Pod, error)
	GetPods(ctx context.Context, namespace string, labels map[string]string) ([]corev1.Pod, error)

	GetServiceByName(ctx context.Context, namespace, name string) (*corev1.Service, error)
	GetServices(ctx context.Context, namespace string, labels map[string]string) ([]corev1.Service, error)

	GetSecretByName(ctx context.Context, namespace, name string) (*corev1.Secret, error)
	GetSecrets(ctx context.Context, namespace string, labels map[string]string) ([]corev1.Secret, error)

	GetConfigMapByName(ctx context.Context, namespace, name string) (*corev1.ConfigMap, error)
	GetConfigMaps(ctx context.Context, namespace string, labels map[string]string) ([]corev1.ConfigMap, error)

	GetPersistentVolumeByName(ctx context.Context, name string) (*corev1.PersistentVolume, error)
	GetPersistentVolumes(ctx context.Context) ([]corev1.PersistentVolume, error)

	GetPersistentVolumeClaimByName(ctx context.Context, namespace, name string) (*corev1.PersistentVolumeClaim, error)
	GetPersistentVolumeClaims(ctx context.Context, namespace string, labels map[string]string) ([]corev1.PersistentVolumeClaim, error)

	GetVolumeSnapshot(ctx context.Context, namespace, name string) (*k8s.VolumeSnapshot, error)
	GetVolumeSnapshots(ctx context.Context, namespace string, labels map[string]string) []k8s.VolumeSnapshot

	GetSnapshotSchedule(ctx context.Context, namespace, name string) (*unstructured.Unstructured, error)
	GetSnapshotSchedules(ctx context.Context, namespace string, labels map[string]string) []unstructured.Unstructured

	GetIngress(ctx context.Context, namespace, name string) (*unstructured.Unstructured, error)
}
