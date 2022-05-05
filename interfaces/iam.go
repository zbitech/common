package interfaces

import (
	"context"
	"github.com/zbitech/common/pkg/model/entity"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/zbitech/common/pkg/model/ztypes"
)

type AuthorizationFactoryIF interface {
	Init(ctx context.Context) error
	GetJwtServer() JwtServerIF
	GetAccessAuthorizer() AccessAuthorizerIF
	GetIAMService() IAMServiceIF
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
	RegisterUser(ctx context.Context, user *entity.User, pass string) error
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

	GetTeam(ctx context.Context, teamId string) (*entity.Team, error)
	UpdateTeam(ctx context.Context, team *entity.Team) error
	DeleteTeam(ctx context.Context, teamId string) error
	GetTeamByOwner(ctx context.Context, owner string) (*entity.Team, error)

	GetTeamMembers(ctx context.Context, teamId string) ([]entity.TeamMember, error)
	AddTeamMember(ctx context.Context, teamId string, member entity.TeamMember) error
	RemoveTeamMember(ctx context.Context, teamId string, key string) error
	UpdateTeamMember(ctx context.Context, member *entity.TeamMember) error

	GetAllMemberships(ctx context.Context) ([]entity.TeamMember, error)
	GetTeamMemberships(ctx context.Context, email string) ([]entity.TeamMember, error)
	GetTeamMembership(ctx context.Context, key string) (*entity.TeamMember, error)
	GetTeamMembershipByEmail(ctx context.Context, team, email string) (*entity.TeamMember, error)
}

type AccessAuthorizerIF interface {
	ValidateProjectAction(ctx context.Context, project string, action ztypes.ZBIAction) error
	ValidateInstanceAction(ctx context.Context, project, instance string, action ztypes.ZBIAction) error
	ValidateTeamAction(ctx context.Context, team string, action ztypes.ZBIAction) error
	ValidateUserInstanceMethodAccess(ctx context.Context, project, instance, method string) (ztypes.SubscriptionLevel, error)
	ValidateAPIKeyInstanceMethodAccess(ctx context.Context, project, instance, method, apikey string) (ztypes.SubscriptionLevel, error)
}
