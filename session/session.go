package usersession

import (
	"context"
	"errors"
	"fmt"
	"github.com/nedpals/supabase-go"
	users "github.com/reward-rabieth/b2b/db/sqlc"
	"github.com/reward-rabieth/b2b/session/models"
	"github.com/reward-rabieth/b2b/util"
)

const AuthCtxKey = "authCtx"

type Component interface {
	GetAuthContextByAccessToken(authToken string) (models.AuthContext, error)
	GetAuthContextFromCtx(ctx context.Context) *models.AuthContext
	CreateUser(ctx context.Context, email, password, username string, repos users.Store, role string) (interface{}, error)
	LoginUser(ctx context.Context, email, password string, repo users.Store) (interface{}, error)
}

type component struct {
	sbClient *supabase.Client
}

func NewComponent(sbClient *supabase.Client) Component {
	return &component{
		sbClient: sbClient,
	}
}

func (c *component) GetAuthContextByAccessToken(authToken string) (models.AuthContext, error) {
	user, err := c.sbClient.Auth.User(context.Background(), authToken)
	if err != nil {
		return models.AuthContext{}, err
	}
	userType := determineUserType(user.UserMetadata)
	return models.AuthContext{
		UserId:   user.ID,
		UserType: userType,
	}, nil
}

func (c *component) GetAuthContextFromCtx(ctx context.Context) *models.AuthContext {
	return nil

}

func (c *component) CreateUser(ctx context.Context, email, password, username string, repo users.Store, role string) (interface{}, error) {
	// Resolve the userRole
	//based on the provided role
	var userRole int
	switch role {
	case "procurer":
		userRole = 1
	case "approver":
		userRole = 2
	case "supplier":
		userRole = 3
	default:
		return nil, errors.New("invalid user role")
	}

	user, err := c.sbClient.Auth.SignUp(ctx, supabase.UserCredentials{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	hashedPassword, _ := util.HashPassword(password)
	args := users.CreateUserParams{
		Email:    user.Email,
		UserID:   user.ID,
		Password: hashedPassword,
		Username: username,
		RoleID:   int32(userRole),
	}
	createUser, err := repo.CreateUser(ctx, args)
	if err != nil {
		return nil, err
	}
	return createUser, nil
}

func (c *component) LoginUser(ctx context.Context, email, password string, repo users.Store) (interface{}, error) {
	AuthUser, err := c.sbClient.Auth.SignIn(ctx, supabase.UserCredentials{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to sign in user: %w", err)
	}

	return AuthUser, nil
}

func determineUserType(userMetaTypeData map[string]interface{}) string {
	userType, ok := userMetaTypeData["user_type"].(string)
	if !ok {
		return "unknown"
	}
	switch userType {
	case "procurer":
		return "procurer"
	case "approver":
		return "approver"

	case "supplier":
		return "supplier"
	}
	return "unknown"
}
