package config

import (
	"fmt"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/object"
	"github.com/zbitech/common/pkg/model/ztypes"
	"text/template"
)

type AppConfig struct {
	AssetsPath  string
	MaxFileSize int
	Features    struct {
		EmailServiceEnabled        bool
		RegistrationInviteEnabled  bool
		CreateAdminUser            bool
		AccessAuthorizationEnabled bool
		TeamsEnabled               bool
		ApiKeysEnabled             bool
	}
	Kubernetes struct {
		InCluster  bool
		KubeConfig string
		Informer   struct {
			RequeueLimit int
			RequeueDelay int
		}
	}
	Repository struct {
		Database struct {
			Factory string
			Url     string
			Name    string
		}
		Authentication struct {
			Type string
		}
		JwtConfig struct {
			SecretKey string
		}
	}
	Mailer struct {
		Host          string
		Port          string
		Username      string
		Password      string
		Sender        string
		TemplateFiles []string
	}
	Cors struct {
		TrustedOrigins []string
	}
	Envoy  EnvoyConfig
	Policy GlobalPolicy
}

type AdminConfig struct {
	Users     []entity.User
	Keys      []entity.APIKey
	Passwords map[string]string
	Teams     []entity.Team
	Members   []entity.TeamMember
}

type EnvoyConfig struct {
	Image                 string
	Command               []string
	Timeout               float32
	AccessAuthorization   bool
	AuthServerURL         string
	AuthServerPort        int32
	AuthenticationEnabled bool
	AuthenticationMethod  string
}

type ServiceInfo struct {
	Prefix    string
	ProxyPort int32
}

type ContainerImage struct {
	Name    string
	Version string
	URL     string
	Port    int32
}

type VersionedResourceConfig struct {
	Version   string           `json:"version,omitempty"`
	Service   *ServiceInfo     `json:"service,omitempty"`
	Images    []ContainerImage `json:"images,omitempty"`
	Templates *struct {
		ActionKeys map[string][]string
		Keys       []string
		File       string
	} `json:"templates,omitempty"`
	Methods      map[string][]string `json:"methods,omitempty"`
	Volumes      []string            `json:"volumes,omitempty"`
	fileTemplate *object.FileTemplate
}

type AppResourceConfig struct {
	Version  string
	Versions map[string]*VersionedResourceConfig
}

type InstanceResourceConfig struct {
	Name     string
	Type     string
	Ports    map[string]int32
	Versions map[string]*VersionedResourceConfig
}

type ProjectResourceConfig struct {
	Name     string
	Versions map[string]*VersionedResourceConfig
}

type ResourceConfig struct {
	App       *AppResourceConfig       `json:"app,omitempty"`
	Project   *ProjectResourceConfig   `json:"project,omitempty"`
	Instances []InstanceResourceConfig `json:"instances,omitempty"`
}

func (r *ResourceConfig) getVersions(versions map[string]*VersionedResourceConfig) []string {
	list := make([]string, 0, len(versions))
	for k, _ := range versions {
		list = append(list, k)
	}
	return list
}

func (r *ResourceConfig) GetProjectVersions() []string {
	return r.getVersions(r.Project.Versions)
}

func (r *ResourceConfig) GetInstanceVersions(iType ztypes.InstanceType) []string {
	instance, ok := r.GetInstanceResourceConfig(iType)
	if !ok {
		return []string{}
	}

	return r.getVersions(instance.Versions)
}

func (v *VersionedResourceConfig) Init(path string, fmap template.FuncMap) error {
	filePath := fmt.Sprintf("%s/%s", path, v.Templates.File)
	f, err := object.NewFileTemplate(filePath, fmap)
	if err != nil {
		return err
	}

	v.fileTemplate = f
	return nil
}

func (v *VersionedResourceConfig) GetFileTemplate() *object.FileTemplate {
	return v.fileTemplate
}

func (v *VersionedResourceConfig) GetImage(name string) *ContainerImage {
	for _, image := range v.Images {
		if image.Name == name {
			return &image
		}
	}
	return nil
}

func (v *VersionedResourceConfig) GetCategoryMethods(cat string) []string {
	for key, methods := range v.Methods {
		if key == cat {
			return methods
		}
	}

	return nil
}

func (v *VersionedResourceConfig) GetAllMethods() []string {
	allMethods := make([]string, 0)
	for _, methods := range v.Methods {
		allMethods = append(allMethods, methods...)
	}
	return allMethods
}

func (r *ResourceConfig) GetInstanceResourceConfig(iType ztypes.InstanceType) (*InstanceResourceConfig, bool) {
	typeStr := string(iType)
	for _, instance := range r.Instances {
		if instance.Type == typeStr {
			return &instance, true
		}
	}

	return nil, false
}
