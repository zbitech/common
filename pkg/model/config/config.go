package config

import (
	"fmt"
	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/object"
	"github.com/zbitech/common/pkg/model/ztypes"
)

type AppConfig struct {
	AssetsPath  string
	MaxFileSize int
	Kubernetes  struct {
		InCluster  bool
		KubeConfig string
		Informer   struct {
			RequeueLimit int
			RequeueDelay int
		}
	}
	Database struct {
		Factory string
		Mongodb struct {
			Url    string
			Dbname string
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
	Policy GlobalPolicy
}

type AdminConfig struct {
	Users     []entity.User
	Keys      []entity.APIKey
	Passwords map[string]string
	Teams     []entity.Team
	Members   []entity.TeamMember
}

type ContainerImage struct {
	Name    string
	Version string
	URL     string
	Port    int32
}

type VersionedResourceConfig struct {
	Version   string
	Images    []ContainerImage
	Templates struct {
		Keys []string
		File string
	}
	Methods      map[string][]string
	fileTemplate *object.FileTemplate
}

type InstanceResourceConfig struct {
	Name     string
	Type     string
	Versions map[string]VersionedResourceConfig
}

type ProjectResourceConfig struct {
	Name     string
	Versions map[string]VersionedResourceConfig
}

type ResourceConfig struct {
	Project   ProjectResourceConfig
	Instances []InstanceResourceConfig
}

func (v *VersionedResourceConfig) GetFileTemplate(path string) (*object.FileTemplate, error) {
	if v.fileTemplate == nil {
		path := fmt.Sprintf("%s/%s", path, v.Templates.File)
		f, err := object.NewFileTemplate(path)
		if err != nil {
			return nil, err
		}

		v.fileTemplate = f
	}

	return v.fileTemplate, nil
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
