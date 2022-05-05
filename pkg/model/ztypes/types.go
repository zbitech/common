package ztypes

var (
	TESTNET = []string{"testnet.z.cash"}
	MAINNET = []string{"mainnet.z.cash"}

	ZCASH_MAX_CONNECTIONS   = "6"
	ZCASH_RPCCLIENT_TIMEOUT = "30"
	ZCASH_SOLVER            = "tromp"
	ZCASH_MAINNET_SOLVER    = "default"
	ZCASH_TESTNET_SOLVER    = "tromp"
)

type InstanceType string

const (
	ZCASH_INSTANCE InstanceType = "zcash"
	LWD_INSTANCE   InstanceType = "lwd"
)

type ResourceLevelType string

const (
	NO_LEVEL       ResourceLevelType = "none"
	PROJECT_LEVEL  ResourceLevelType = "project"
	INSTANCE_LEVEL ResourceLevelType = "instance"
)

type ResourceObjectType string

const (
	NAMESPACE_RESOURCE               ResourceObjectType = "Namespace"
	DEPLOYMENT_RESOURCE              ResourceObjectType = "Deployment"
	SERVICE_RESOURCE                 ResourceObjectType = "Service"
	CONFIGMAP_RESOURCE               ResourceObjectType = "ConfigMap"
	SECRET_RESOURCE                  ResourceObjectType = "Secret"
	POD_RESOURCE                     ResourceObjectType = "Pod"
	PERSISTENT_VOLUME_RESOURCE       ResourceObjectType = "PersistentVolume"
	PERSISTENT_VOLUME_CLAIM_RESOURCE ResourceObjectType = "PersistentVolumeClaim"
	VOLUME_SNAPHOT_RESOURCE          ResourceObjectType = "VolumeSnapshot"
)

type NetworkType string

const (
	MAINNET_TYPE NetworkType = "mainnet"
	TESTNET_TYPE NetworkType = "testnet"
)

type Role string

const (
	ADMIN_ROLE Role = "admin"
	OWNER_ROLE Role = "owner"
	USER_ROLE  Role = "user"
)

// type TeamRole string

// const (
// 	TEAM_ADMIN_ROLE TeamRole = "admin"
// 	TEAM_USER_ROLE  TeamRole = "user"
// )

type SubscriptionLevel string

const (
	NO_SUB_LEVEL      SubscriptionLevel = "none"
	TEAM_MEMBER_LEVEL SubscriptionLevel = "team_member"
	GOLD_LEVEL        SubscriptionLevel = "gold"
	SILVER_LEVEL      SubscriptionLevel = "silver"
	BRONZE_LEVEL      SubscriptionLevel = "bronze"
)

type InvitationStatus string

const (
	NEW_INVITATION     InvitationStatus = "pending"
	ACCEPT_INVITATION  InvitationStatus = "accepted"
	REJECT_INVITATION  InvitationStatus = "rejected"
	EXPIRED_INVITATION InvitationStatus = "expired"
)

type ZBIAction string

const (
	ACTION_CREATE ZBIAction = "create"
	ACTION_UPDATE ZBIAction = "update"
	ACTION_DELETE ZBIAction = "delete"
	ACTION_REMOVE ZBIAction = "remove"
	ACTION_ACCESS ZBIAction = "access"
)

type InstanceRequestIF interface {
	GetName() string
	GetVersion() string
	GetInstanceType() InstanceType
	AllowMethods() bool
}

// type ObjectIF interface {
// 	GetId() string
// 	SetId(id string)
// 	GetNetwork() NetworkType
// 	GetOwner() string
// 	SetOwner(owner string)
// 	GetStatus() string
// 	SetStatus(status string)
// 	GetTimestamp() time.Time
// 	SetTimestamp(tstamp time.Time)
// }

type InstanceIF interface {
	//	ObjectIF
	GetName() string
	GetProject() string
	GetVersion() string
	GetNetwork() NetworkType
	GetOwner() string
	GetInstanceType() InstanceType

	//	SetProject(project string)
	//	GetStatus() string
	//	SetStatus(status string)
	//	GetTimestamp() time.Time
	//	SetTimestamp(tstamp time.Time)
}

type InstanceDetailIF interface {
	// json.Marshaler
	// json.Unmarshaler
	// bson.Marshaler
	// bson.Unmarshaler
}

// type ProjectSpecIF interface {
// 	ProjectIF
// 	GetProject() ProjectIF
// }

type InstanceSpecIF interface {
	// InstanceIF
	//	GetInstance() *Instance
}

type ResourceStateIF interface {
	IsActive() bool
	IsDeleted() bool
}
