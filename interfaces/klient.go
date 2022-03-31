package interfaces

import (
	"context"

	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/object"
	"github.com/zbitech/common/pkg/model/ztypes"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/restmapper"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
)

type KlientFactoryIF interface {
	Init(ctx context.Context) error
	GetZBIClient() ZBIClientIF
	GetKubernesClient() KlientIF
}

type ZBIClientIF interface {
	RunInformer()

	CreateProject(ctx context.Context, request *object.ProjectRequest) (*entity.Project, error)
	CreateProjectIngress(ctx context.Context, project *entity.Project, instance *entity.Instance) error

	UpdateProject(ctx context.Context, project *entity.Project, request *object.ProjectRequest) (*entity.Project, error)
	DeleteProject(ctx context.Context, name string) error
	GetAllProjects(ctx context.Context) ([]entity.Project, error)
	GetProject(ctx context.Context, name string) (*entity.Project, error)
	GetProjectsByOwner(ctx context.Context, owner string) ([]entity.Project, error)
	GetProjectsByTeam(ctx context.Context, team string) ([]entity.Project, error)
	GetProjectResources(ctx context.Context, name string) ([]entity.KubernetesResource, error)

	CreateInstance(ctx context.Context, project *entity.Project, request ztypes.InstanceRequestIF) (*entity.Instance, error)
	UpdateInstance(ctx context.Context, project *entity.Project, instance *entity.Instance, request ztypes.InstanceRequestIF) (*entity.Instance, error)
	DeleteInstance(ctx context.Context, projectName, instanceName string) error
	GetAllInstances(ctx context.Context) ([]entity.Instance, error)
	GetInstancesByProject(ctx context.Context, projectName string) ([]entity.Instance, error)
	GetInstancesByOwner(ctx context.Context, owner string) ([]entity.Instance, error)
	GetInstanceByName(ctx context.Context, projectName, instanceName string) (*entity.Instance, error)
	GetInstanceResources(ctx context.Context, project_name, instance_name string) ([]entity.KubernetesResource, error)
}

type KlientIF interface {
	GetMapper() *restmapper.DeferredDiscoveryRESTMapper
	GetGVR(gvk schema.GroupVersionKind) (*schema.GroupVersionResource, error)

	GetResource(object *unstructured.Unstructured) *entity.KubernetesResource
	GenerateKubernetesObjects(ctx context.Context, spec_arr []string) ([]*unstructured.Unstructured, error)

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

	GetDeploymentByName(ctx context.Context, namespace, name string) (*appsv1.Deployment, error)
	GetDeployments(ctx context.Context, namespace string, labels map[string]string) ([]appsv1.Deployment, error)

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
}
