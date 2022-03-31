package interfaces

import (
	"context"
	"time"

	"github.com/zbitech/common/pkg/model/entity"
	"github.com/zbitech/common/pkg/model/ztypes"
)

type RepositoryFactoryIF interface {
	Init(ctx context.Context, create_db, load_db bool) error
	OpenConnection(ctx context.Context) error
	GetProjectRepository() ProjectRepositoryIF
	GetAdminRepository() AdminRepositoryIF
	CloseConnection(ctx context.Context) error
	CreateDatabase(ctx context.Context, purge, load bool) error
}

type ProjectRepositoryIF interface {
	GetProjects(ctx context.Context) ([]entity.Project, error)
	CreateProject(ctx context.Context, project *entity.Project) error
	UpdateProject(ctx context.Context, project *entity.Project) error

	GetProject(ctx context.Context, name string) (*entity.Project, error)
	GetProjectsByOwner(ctx context.Context, owner string) ([]entity.Project, error)
	GetProjectsByTeam(ctx context.Context, team string) ([]entity.Project, error)
	UpdateProjectStatus(ctx context.Context, project string, status string) error

	GetInstances(ctx context.Context) ([]entity.Instance, error)
	CreateInstance(ctx context.Context, instance *entity.Instance) error
	GetInstance(ctx context.Context, project, name string) (*entity.Instance, error)
	GetInstancesByProject(ctx context.Context, project string) ([]entity.Instance, error)
	GetInstancesByOwner(ctx context.Context, owner string) ([]entity.Instance, error)
	UpdateInstanceStatus(ctx context.Context, project, name, status string) error

	GetProjectResources(ctx context.Context, project string) ([]entity.KubernetesResource, error)
	SaveProjectResource(ctx context.Context, project string, resource *entity.KubernetesResource) error
	SaveProjectResources(ctx context.Context, project string, resources []entity.KubernetesResource) error

	GetInstanceResources(ctx context.Context, project, instance string) ([]entity.KubernetesResource, error)
	SaveInstanceResource(ctx context.Context, project, instance string, resource *entity.KubernetesResource) error
	SaveInstanceResources(ctx context.Context, project, instance string, resources []entity.KubernetesResource) error

	GetUserSummary(ctx context.Context, userId string) (*entity.ResourceSummary, error)
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

	StoreInstancePolicy(ctx context.Context, p entity.InstancePolicy) error
	StoreInstancePolicies(ctx context.Context, p []entity.InstancePolicy) error
	GetInstancePolicy(ctx context.Context, project, instance string) (*entity.InstancePolicy, error)
	GetInstanceMethodPolicy(ctx context.Context, project, instance, methodName string) (*entity.MethodPolicy, error)
	GetInstanceMethodPolicies(ctx context.Context, project, instance, methodCategory string) ([]entity.MethodPolicy, error)

	GetExpiringInvitations(ctx context.Context, date time.Time) ([]entity.TeamMember, error)
	PurgeExpiredInvitations(ctx context.Context) (int64, error)

	CreateTeam(ctx context.Context, team entity.Team) error
	GetTeams(ctx context.Context) ([]entity.Team, error)
	GetTeam(ctx context.Context, teamId string) (*entity.Team, error)

	UpdateTeam(ctx context.Context, team entity.Team) error
	DeleteTeam(ctx context.Context, teamId string) error
	GetTeamByOwner(ctx context.Context, owner string) (*entity.Team, error)
	GetTeamMembers(ctx context.Context, teamId string) ([]entity.TeamMember, error)
	AddTeamMember(ctx context.Context, teamId string, member entity.TeamMember) error
	RemoveTeamMembers(ctx context.Context, teamId string, key []string) error
	RemoveTeamMember(ctx context.Context, teamId string, key string) error

	UpdateTeamMemberEmail(ctx context.Context, teamId, key, email string) error
	UpdateTeamMemberRole(ctx context.Context, teamId, key string, role ztypes.Role) error
	UpdateTeamMemberStatus(ctx context.Context, teamId, key string, status ztypes.InvitationStatus) error

	GetAllMemberships(ctx context.Context) ([]entity.TeamMember, error)
	GetTeamMemberships(ctx context.Context, email string) ([]entity.TeamMember, error)
	GetTeamMembership(ctx context.Context, key string) (*entity.TeamMember, error)
	GetTeamMembershipByEmail(ctx context.Context, team, email string) (*entity.TeamMember, error)
}
