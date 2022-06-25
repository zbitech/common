package k8s

import (
	"github.com/zbitech/common/pkg/model/ztypes"
	"time"
)

type Deployment struct {
	Name           string
	Available      int
	UpToDate       int
	Replicas       int
	ReadyReplicas  int
	Age            string
	UpdateStrategy string
	Pods           []Pod
}

type ConfigMap struct {
	Name string
	Data []string
	Age  string
}

type Pod struct {
	Name       string
	Status     string
	Containers []Container
	Age        string
}

type Container struct {
	Name         string
	Ready        bool
	RestartCount int
	Started      bool
	State        string
}

type StorageClass struct {
	Name          string
	Default       bool
	Provisioner   string
	ReclaimPolicy string
	BindingMode   string
	Expandable    bool
	Age           string
}

type SnapshotClass struct {
	Name           string
	Driver         string
	DeletionPolicy string
	Age            string
}

type InstanceVolume struct {
	Name             string
	Status           string
	VolumeType       string
	PersistentVolume string
	Capacity         string
	AccessModes      string
	StorageClass     string
	Age              string
}

type Snapshot struct {
	Name          string
	Source        string
	SnapshotClass string
	Age           string
}

type SnapshotSchedule struct {
	Name             string
	Type             ztypes.ZBIBackupScheduleType
	SnapshotClass    string
	Labels           map[string]string
	Expiration       string
	MaxCount         int
	Age              string
	LastSnapshotTime time.Time
	NextSnapshotTime time.Time
}

type Metadata struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	Labels            map[string]string `json:"labels,omitempty"`
	ResourceVersion   string            `json:"resourceVersion,omitempty"`
	Uid               string            `json:"uid,omitempty"`
	CreationTimestamp string            `json:"creationTimestamp"`
}

type VolumeSnapshot struct {
	ApiVersion string   `json:"apiVersion"`
	Metadata   Metadata `json:"metadata"`
	Spec       struct {
		VolumeSnapshotClassName string `json:"volumeSnapshotClassName"`
		Source                  struct {
			PersistentVolumeClaimName string `json:"persistentVolumeClaimName"`
		} `json:"source"`
	} `json:"spec"`
	Status struct {
		BoundVolumeSnapshotContentName string    `json:"boundVolumeSnapshotContentName"`
		CreationTime                   time.Time `json:"creationTime"`
		ReadyToUse                     bool      `json:"readyToUse"`
		RestoreSize                    string    `json:"restoreSize"`
	} `json:"status,omitempty"`
}

type VolumeSnapshotSchedule struct {
	ApiVersion string   `json:"apiVersion"`
	Metadata   Metadata `json:"metadata"`
	Spec       struct {
		ClaimSelector struct {
			MatchLabels map[string]string `json:"matchLabels,omitempty"`
		} `json:"claimSelector,omitempty"`
	} `json:"spec"`
	Disabled  bool `json:"disabled"`
	Retention struct {
		Expires  string `json:"expires"`
		MaxCount int    `json:"maxCount"`
	} `json:"retention"`
	Schedule         string `json:"schedule"`
	SnapshotTemplate struct {
		Labels            map[string]string `json:"labels"`
		SnapshotClassName string            `json:"snapshotClassName"`
	} `json:"snapshotTemplate"`
	Status struct {
		Conditions []struct {
			LastHeartBeatTime  time.Time `json:"lastHeartBeatTime"`
			LastTransitionTime time.Time `json:"lastTransitionTime"`
			Message            string    `json:"message"`
			Reason             string    `json:"reason"`
			Status             string    `json:"status"`
			Type               string    `json:"type"`
		} `json:"conditions"`
		NextSnapshotTime time.Time `json:"nextSnapshotTime"`
	} `json:"status,omitempty"`
}

type IngressService struct {
	Name string `json:"name"`
	Port int32  `json:"port"`
}

type IngressCondition struct {
	Prefix string `json:"prefix,omitempty"`
}

type IngressRoute struct {
	Conditions        []IngressCondition       `json:"conditions,omitempty"`
	Services          []IngressService         `json:"services,omitempty"`
	PathRewritePolicy IngressPathRewritePolicy `json:"pathRewritePolicy,omitempty"`
}

type IngressPathRewritePolicy struct {
	ReplacePrefix []struct {
		Replacement string `json:"replacement,omitempty"`
	} `json:"replacePrefix,omitempty"`
}

type IngressInclude struct {
	Name       string             `json:"name,omitempty"`
	Namespace  string             `json:"namespace,omitempty"`
	Conditions []IngressCondition `json:"conditions,omitempty"`
}

type Ingress struct {
	ApiVersion string   `json:"apiVersion"`
	Metadata   Metadata `json:"metadata"`
	Spec       struct {
		Includes    []IngressInclude `json:"includes,omitempty"`
		Virtualhost struct {
			Fqdn string `json:"fqdn"`
			Tls  struct {
				SecretName string `json:"secretName"`
			} `json:"tls"`
		} `json:"virtualhost,omitempty"`
		Routes []IngressRoute `json:"routes,omitempty"`
	} `json:"spec"`
	Status struct {
		Conditions []struct {
			Errors []struct {
				Message string `json:"message,omitempty"`
				Reason  string `json:"reason,omitempty"`
				Status  string `json:"status,omitempty"`
				Type    string `json:"type,omitempty"`
			} `json:"errors,omitempty"`
			LastTransitionTime string `json:"lastTransitionTime,omitempty"`
			Message            string `json:"message,omitempty"`
			ObservedGeneration int    `json:"observedGeneration,omitempty"`
			Reason             string `json:"reason,omitempty"`
			Status             string `json:"status,omitempty"`
			Type               string `json:"type,omitempty"`
		} `json:"conditions,omitempty"`
		CurrentStatus string `json:"currentStatus,omitempty"`
		Description   string `json:"description,omitempty"`
		LoadBalancer  struct {
			Ingress []struct {
				Hostname string `json:"hostname,omitempty"`
			} `json:"ingress,omitempty"`
		} `json:"loadBalancer,omitempty"`
	} `json:"status,omitempty"`
}
