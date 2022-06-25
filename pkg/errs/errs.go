package errs

import "errors"

var (
	ErrMarshalFailed     error = errors.New("object marshal error")
	ErrUserAlreadyExists error = errors.New("user already exists")

	ErrDBError            error = errors.New("Database error")
	ErrInvalidDBAction    error = errors.New("Request action is not valid")
	ErrDBItemNotFound     error = errors.New("Item not found in data store")
	ErrDBItemUpdateFailed error = errors.New("Item update failed")
	ErrDBItemInsertFailed error = errors.New("Item insert failed")
	ErrDBItemDeleteFailed error = errors.New("Item delete failed")
	ErrDBKeyAlreadyExists error = errors.New("repo item key already exists")

	ErrUserNotAuth         error = errors.New("User not authenticated")
	ErrAuthFailed          error = errors.New("Authentication failed")
	ErrInvalidAuthMethod   error = errors.New("Invalid authentication method")
	ErrTokenValidation     error = errors.New("Token validation error")
	ErrInvalidToken        error = errors.New("Invalid Authentication token")
	ErrExpiredToken        error = errors.New("Authentication Token expired")
	ErrUnregisteredUser    error = errors.New("user is not registered")
	ErrRegistrationFailure error = errors.New("user registration failed")

	ErrMaxProjectsCreated  error = errors.New("Max projects already created")
	ErrMaxInstancesCreated error = errors.New("Max instances already created")
	ErrMaxTeamsCreated     error = errors.New("Max teams already created")

	ErrProjectCreateNotAllowed error = errors.New("Not authorized to create project")
	ErrProjectUpdateNotAllowed error = errors.New("Not authorized to update project")
	ErrProjectDeleteNotAllowed error = errors.New("Not authorized to delete project")
	ErrProjectAccessNotAllowed error = errors.New("Not authorized to access project")
	ErrProjectAccessError      error = errors.New("Unknown project access error")

	ErrInstanceCreateNotAllowed error = errors.New("Not authorized to create instance")
	ErrInstanceUpdateNotAllowed error = errors.New("Not authorized to update instance")
	ErrInstanceDeleteNotAllowed error = errors.New("Not authorized to delete instance")
	ErrInstanceAccessNotAllowed error = errors.New("Not authorized to access instance")
	ErrInstanceAccessError      error = errors.New("Unknown instance access error")

	ErrInstancePolicyAccessNotAllowed error = errors.New("Not authorized to access instance policy")

	ErrTeamCreateNotAllowed error = errors.New("Not authorized to create team")
	ErrTeamUpdateNotAllowed error = errors.New("Not authorized to update team")
	ErrTeamDeleteNotAllowed error = errors.New("Not authorized to delete team")
	ErrTeamAccessNotAllowed error = errors.New("Not authorized to access team")
	ErrTeamAccessError      error = errors.New("Unknown team access error")

	ErrProjectResourceFailed  error = errors.New("Project Resource manager failed")
	ErrProjectDataFailed      error = errors.New("Project Data manager failed")
	ErrInstanceResourceFailed error = errors.New("Instance Resource manager failed")
	ErrInstanceDataFailed     error = errors.New("Instance Data manager failed")

	ErrIngressResourceFailed error = errors.New("App Resource manager failed")

	ErrKubernetesConnFailed     error = errors.New("Kubernetes connection failed")
	ErrKubernetesResourceFailed error = errors.New("Kubernetes resource failed")
)
