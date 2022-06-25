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
	InstanceTypeZCASH InstanceType = "zcash"
	InstanceTypeLWD   InstanceType = "lwd"
)

type ResourceLevelType string

const (
	NO_LEVEL       ResourceLevelType = "none"
	PROJECT_LEVEL  ResourceLevelType = "project"
	INSTANCE_LEVEL ResourceLevelType = "instance"
)

type ResourceObjectType string

const (
	ResourceNamespace             ResourceObjectType = "Namespace"
	ResourceDeployment            ResourceObjectType = "Deployment"
	ResourceService               ResourceObjectType = "Service"
	ResourceConfigMap             ResourceObjectType = "ConfigMap"
	ResourceSecret                ResourceObjectType = "Secret"
	ResourcePod                   ResourceObjectType = "Pod"
	ResourcePersistentVolume      ResourceObjectType = "PersistentVolume"
	ResourcePersistentVolumeClaim ResourceObjectType = "PersistentVolumeClaim"
	ResourceVolumeSnapshot        ResourceObjectType = "VolumeSnapshot"
	ResourceVolumeSnapshotClass   ResourceObjectType = "VolumeSnapshotClass"
	ResourceSnapshotSchedule      ResourceObjectType = "SnapshotSchedule"
	ResourceHTTPProxy             ResourceObjectType = "HTTPProxy"
)

type NetworkType string

const (
	NetworkTypeMain NetworkType = "mainnet"
	NetworkTypeTest NetworkType = "testnet"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleOwner Role = "owner"
	RoleUser  Role = "user"
)

type SubscriptionLevel string

const (
	SubscriptionNone        SubscriptionLevel = "none"
	SubscriptionTeamMember  SubscriptionLevel = "team_member"
	SubscriptionGoldLevel   SubscriptionLevel = "gold"
	SubscriptionSilverLevel SubscriptionLevel = "silver"
	SubscriptionBronzeLevel SubscriptionLevel = "bronze"
)

type InvitationStatus string

const (
	InvitationPending InvitationStatus = "pending"
	InvitationAccept  InvitationStatus = "accepted"
	InvitationReject  InvitationStatus = "rejected"
	InvitationExpired InvitationStatus = "expired"
)

type PlatformAction string

const (
	PlatformResourceCreate PlatformAction = "create"
	PlatformResourceUpdate PlatformAction = "update"
	PlatformResourceDelete PlatformAction = "delete"
	PlatformResourceRemove PlatformAction = "remove"
	PlatformResourceAccess PlatformAction = "access"
)

type ZBIManagerAction string

const (
	CreateResourceAction ZBIManagerAction = "create"
	UpdateResourceAction ZBIManagerAction = "update"
	DeleteResourceAction ZBIManagerAction = "delete"
	StopResourceAction   ZBIManagerAction = "stop"
	StartResourceAction  ZBIManagerAction = "start"
	RotateResourceAction ZBIManagerAction = "rotate"
)

type ZBIBackupType string

const (
	VolumeSnapshot ZBIBackupType = "snapshot"
	BackupSchedule ZBIBackupType = "schedule"
)

type ZBIBackupScheduleType string

const (
	DailySnapshotSchedule   ZBIBackupScheduleType = "daily"
	WeeklySnapshotSchedule  ZBIBackupScheduleType = "weekly"
	MonthlySnapshotSchedule ZBIBackupScheduleType = "monthly"
)

type DataSourceType string

const (
	NoDataSource       DataSourceType = "none"
	VolumeDataSource   DataSourceType = "volume"
	SnapshotDataSource DataSourceType = "snapshot"
)

type EventAction string

const (
	EventActionCreate         EventAction = "create"
	EventActionDelete         EventAction = "delete"
	EventActionUpdate         EventAction = "update"
	EventActionResource       EventAction = "resource"
	EventActionDeactivate     EventAction = "deactivate"
	EventActionReactivate     EventAction = "reactivate"
	EventActionRepair         EventAction = "repair"
	EventActionSnapshot       EventAction = "snapshot"
	EventActionSchedule       EventAction = "schedule"
	EventActionPurge          EventAction = "purge"
	EventActionStopInstance   EventAction = "stop"
	EventActionStartInstance  EventAction = "start"
	EventActionRotate         EventAction = "rotate"
	EventActionUpdatePolicy   EventAction = "updatepolicy"
	EventActionAddMember      EventAction = "addmember"
	EventActionRemoveMember   EventAction = "removemember"
	EventActionUpdateMember   EventAction = "updatemember"
	EventActionRegister       EventAction = "register"
	EventActionCreateKey      EventAction = "createkey"
	EventActionDeleteKey      EventAction = "deletekey"
	EventActionChangePassword EventAction = "changepassword"
	EventActionChangeEmail    EventAction = "changeemail"
	EventActionUpdateProfile  EventAction = "updateprofile"
	EventActionAcceptInvite   EventAction = "acceptinvite"
	EventActionRejectInvite   EventAction = "rejectinvite"
	EventActionExpireInvite   EventAction = "expireinvite"
)

type StatusType string

const (
	StatusActive   StatusType = "active"
	StatusStopped  StatusType = "stopped"
	StatusInActive StatusType = "inactive"
)
