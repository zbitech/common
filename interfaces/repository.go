package interfaces

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/zbitech/common/pkg/model/ztypes"
	"time"

	"github.com/zbitech/common/pkg/model/entity"
)

type RepositoryFactoryIF interface {
	Init(ctx context.Context) error
	OpenConnection(ctx context.Context) error
	GetProjectRepository() ProjectRepositoryIF
	GetAdminRepository() AdminRepositoryIF
	CloseConnection(ctx context.Context) error
}

type ProjectStatusServiceIF interface {
	UpdateProjectStatus(ctx context.Context, project string, status string, timestamp time.Time) error
	CreateProjectEvent(ctx context.Context, project, userId string, action ztypes.EventAction, evtErr error) error
	UpdateInstanceStatus(ctx context.Context, project, name, status string, timestamp time.Time) error
	CreateInstanceEvent(ctx context.Context, project, instance, userId string, action ztypes.EventAction, evtErr error) error
}

type ProjectRepositoryIF interface {
	GetProjects(ctx context.Context) ([]entity.Project, error)
	CreateProject(ctx context.Context, project *entity.Project) error
	UpdateProject(ctx context.Context, project *entity.Project) error

	GetProject(ctx context.Context, name string) (*entity.Project, error)
	GetProjectsByOwner(ctx context.Context, owner string) ([]entity.Project, error)
	GetProjectsByTeam(ctx context.Context, team string) ([]entity.Project, error)
	UpdateProjectStatus(ctx context.Context, project string, status ztypes.StatusType, timestamp time.Time) error
	CreateProjectEvent(ctx context.Context, project, actor string, action ztypes.EventAction, evtErr error) error
	GetProjectEvents(ctx context.Context, project string) ([]entity.ProjectEvent, error)

	GetInstances(ctx context.Context) ([]entity.InstanceIF, error)
	CreateInstance(ctx context.Context, instance entity.InstanceIF) error
	UpdateInstance(ctx context.Context, instance entity.InstanceIF) error
	GetInstance(ctx context.Context, project, name string) (entity.InstanceIF, error)
	GetInstancesByProject(ctx context.Context, project string) ([]entity.InstanceIF, error)
	GetInstancesByOwner(ctx context.Context, owner string) ([]entity.InstanceIF, error)
	UpdateInstanceStatus(ctx context.Context, project, name string, status ztypes.StatusType, timestamp time.Time) error
	CreateInstanceEvent(ctx context.Context, project, instance, actor string, action ztypes.EventAction, evtErr error) error
	GetInstanceEvents(ctx context.Context, project, instance string) ([]entity.InstanceEvent, error)

	GetProjectStats(ctx context.Context, project string) *entity.ProjectSummary
	GetOwnerStats(ctx context.Context, owner string) *entity.ProjectSummary

	//	GetUserSummary(ctx context.Context, userId string) (*entity.ResourceSummary, error)
}

type AdminRepositoryIF interface {
	RegisterUser(ctx context.Context, user *entity.User, pass *entity.UserPassword) error
	GetUsers(ctx context.Context) []entity.User
	GetUser(ctx context.Context, userId string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	UpdatePassword(ctx context.Context, userid string, pass *entity.UserPassword) error

	AuthenticateUser(ctx context.Context, userId, password string) (*string, error)

	GetAPIKeys(ctx context.Context, userId string) ([]string, error)
	GetAPIKey(ctx context.Context, apiKey string) (*entity.APIKey, error)
	CreateAPIKey(ctx context.Context, user_id string) (*entity.APIKey, error)
	DeleteAPIKey(ctx context.Context, apiKey string) error

	StoreUserPolicy(ctx context.Context, p entity.UserPolicy) error
	GetUserPolicy(ctx context.Context, userId string) (*entity.UserPolicy, error)
	StoreAPIKeyPolicy(ctx context.Context, p entity.APIKeyPolicy) error
	GetAPIKeyPolicy(ctx context.Context, key string) (*entity.APIKeyPolicy, error)

	CreateProfileEvent(ctx context.Context, userId, actor string, action ztypes.EventAction, evtErr error) error
	GetProfileEvents(ctx context.Context, userId string) ([]entity.ProfileEvent, error)

	StoreInstancePolicy(ctx context.Context, p entity.InstancePolicy) error
	StoreInstancePolicies(ctx context.Context, p []entity.InstancePolicy) error
	GetInstancePolicy(ctx context.Context, project, instance string) (*entity.InstancePolicy, error)
	GetInstanceMethodPolicy(ctx context.Context, project, instance, methodName string) (*entity.MethodPolicy, error)
	GetInstanceMethodPolicies(ctx context.Context, project, instance, methodCategory string) ([]entity.MethodPolicy, error)

	CreateRegistrationInvite(ctx context.Context, invite entity.RegistrationInvite) error
	GetRegistrationInvite(ctx context.Context, key string) (*entity.RegistrationInvite, error)
	GetRegistrationInvites(ctx context.Context) ([]entity.RegistrationInvite, error)
	UpdateRegistrationInvite(ctx context.Context, invite *entity.RegistrationInvite) error

	GetExpiringInvitations(ctx context.Context, date time.Time) ([]entity.TeamMember, error)
	PurgeExpiredInvitations(ctx context.Context) (int64, error)

	CreateTeam(ctx context.Context, team entity.Team) error
	GetTeams(ctx context.Context) ([]entity.Team, error)
	GetTeamsByOwner(ctx context.Context, owner string) ([]entity.Team, error)
	GetTeamsByMembership(ctx context.Context, email string) ([]entity.Team, error)

	GetTeam(ctx context.Context, teamId string) (*entity.Team, error)
	UpdateTeam(ctx context.Context, team *entity.Team) error
	DeleteTeam(ctx context.Context, teamId string) error
	//	GetTeamByOwner(ctx context.Context, owner string) (*entity.Team, error)

	GetTeamMembers(ctx context.Context, teamId string) ([]entity.TeamMember, error)
	AddTeamMember(ctx context.Context, teamId string, member entity.TeamMember) error
	RemoveTeamMember(ctx context.Context, teamId string, key string) error
	UpdateTeamMember(ctx context.Context, member *entity.TeamMember) error

	GetAllMemberships(ctx context.Context) ([]entity.TeamMember, error)
	GetTeamMemberships(ctx context.Context, email string) ([]entity.TeamMember, error)
	GetTeamMembership(ctx context.Context, key string) (*entity.TeamMember, error)
	GetTeamMembershipByEmail(ctx context.Context, team, email string) (*entity.TeamMember, error)

	CreateTeamEvent(ctx context.Context, teamId, actor string, action ztypes.EventAction, evtErr error, targets ...string) error
	GetTeamEvents(ctx context.Context, teamId string) ([]entity.TeamEvent, error)

	GetResourceStats(ctx context.Context, owner string) *entity.ResourceSummary
}

type ServiceFactoryIF interface {
	Init(ctx context.Context) error
	GetJwtServer() JwtServerIF
	GetIAMService() IAMServiceIF
	GetProjectService() ProjectRepositoryIF
	OpenConnection(ctx context.Context) error
	CloseConnection(ctx context.Context) error
}

type AccessAuthorizerFactoryIF interface {
	Init(ctx context.Context) error
	GetAccessAuthorizer() AccessAuthorizerIF
}

type JwtServerIF interface {
	GetKey() (interface{}, error)
	ValidateToken(token *jwt.Token) error
	GetPayload() jwt.Claims
	GetUserId(claim jwt.Claims) string
	GetEmail(claim jwt.Claims) string
	GetRole(claim jwt.Claims) ztypes.Role
}

type IAMServiceIF interface {
	RegisterUser(ctx context.Context, user *entity.User, invite *entity.RegistrationInvite, pass string) error
	DeactivateUser(ctx context.Context, userid string) error
	ReactivateUser(ctx context.Context, userid string) error

	GetUsers(ctx context.Context) []entity.User
	GetUser(ctx context.Context, userId string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	ChangePassword(ctx context.Context, userid, password string) error

	AuthenticateUser(ctx context.Context, userId, password string) (*string, error)
	ValidateAuthToken(ctx context.Context, tokenString string) (jwt.Claims, *entity.User, error)

	GetAPIKeys(ctx context.Context, userId string) ([]string, error)
	GetAPIKey(ctx context.Context, apiKey string) (*entity.APIKey, error)
	CreateAPIKey(ctx context.Context, userid string) (*entity.APIKey, error)
	DeleteAPIKey(ctx context.Context, apiKey string) error

	StoreUserPolicy(ctx context.Context, p entity.UserPolicy) error
	GetUserPolicy(ctx context.Context, userId string) (*entity.UserPolicy, error)
	StoreAPIKeyPolicy(ctx context.Context, p entity.APIKeyPolicy) error
	GetAPIKeyPolicy(ctx context.Context, key string) (*entity.APIKeyPolicy, error)

	CreateProfileEvent(ctx context.Context, userId, actor string, action ztypes.EventAction, evtError error) error
	GetProfileEvents(ctx context.Context, userId string) ([]entity.ProfileEvent, error)

	StoreInstancePolicy(ctx context.Context, p entity.InstancePolicy) error
	StoreInstancePolicies(ctx context.Context, p []entity.InstancePolicy) error
	GetInstancePolicy(ctx context.Context, project, instance string) (*entity.InstancePolicy, error)
	GetInstanceMethodPolicy(ctx context.Context, project, instance, methodName string) (*entity.MethodPolicy, error)
	GetInstanceMethodPolicies(ctx context.Context, project, instance, methodCategory string) ([]entity.MethodPolicy, error)

	CreateRegistrationInvite(ctx context.Context, invite entity.RegistrationInvite) error
	GetRegistrationInvite(ctx context.Context, key string) (*entity.RegistrationInvite, error)
	GetRegistrationInvites(ctx context.Context) ([]entity.RegistrationInvite, error)
	UpdateRegistrationInvite(ctx context.Context, invite *entity.RegistrationInvite) error

	GetExpiringInvitations(ctx context.Context, date time.Time) ([]entity.TeamMember, error)
	PurgeExpiredInvitations(ctx context.Context) (int64, error)

	CreateTeam(ctx context.Context, team entity.Team) error
	GetTeams(ctx context.Context) ([]entity.Team, error)
	GetTeamsByOwner(ctx context.Context, owner string) ([]entity.Team, error)
	GetTeamsByMembership(ctx context.Context, email string) ([]entity.Team, error)

	GetTeam(ctx context.Context, teamId string) (*entity.Team, error)
	UpdateTeam(ctx context.Context, team *entity.Team) error
	DeleteTeam(ctx context.Context, teamId string) error

	GetTeamMembers(ctx context.Context, teamId string) ([]entity.TeamMember, error)
	AddTeamMember(ctx context.Context, teamId string, member entity.TeamMember) error
	RemoveTeamMember(ctx context.Context, teamId string, key string) error
	UpdateTeamMember(ctx context.Context, member *entity.TeamMember) error

	GetAllMemberships(ctx context.Context) ([]entity.TeamMember, error)
	GetTeamMemberships(ctx context.Context, email string) ([]entity.TeamMember, error)
	GetTeamMembership(ctx context.Context, key string) (*entity.TeamMember, error)
	GetTeamMembershipByEmail(ctx context.Context, team, email string) (*entity.TeamMember, error)

	CreateTeamEvent(ctx context.Context, teamId, actor string, action ztypes.EventAction, evtError error, target ...string) error
	GetTeamEvents(ctx context.Context, teamId string) ([]entity.TeamEvent, error)

	GetResourceStats(ctx context.Context, owner string) *entity.ResourceSummary
}

type AccessAuthorizerIF interface {
	ValidateProjectAction(ctx context.Context, project string, action ztypes.PlatformAction) error
	ValidateInstanceAction(ctx context.Context, project, instance string, action ztypes.PlatformAction) error
	ValidateTeamAction(ctx context.Context, team string, action ztypes.PlatformAction) error
	ValidateUserInstanceMethodAccess(ctx context.Context, project, instance, method string) (ztypes.SubscriptionLevel, error)
	ValidateAPIKeyInstanceMethodAccess(ctx context.Context, project, instance, method, apikey string) (ztypes.SubscriptionLevel, error)
}
