package util

import (
	"github.com/google/uuid"
	"github.com/nedpals/supabase-go"
	users "github.com/reward-rabieth/b2b/db/sqlc"
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func GenerateUUID() uuid.UUID {
	return uuid.New()
}

func MapSupabaseUserToParams(user supabase.UserCredentials) users.CreateUserParams {
	return users.CreateUserParams{

		Password: user.Password,
		Email:    user.Email,
	}
}
