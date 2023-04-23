package db

import (
	"context"
	"testing"
	"time"

	"github.com/dipu626/simple-bank/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:          util.RandomOwner(),
		FullName:          util.RandomOwner(),
		HashedPassword:    "",
		PasswordChangedAt: time.Now(),
		Email:             util.RandomEmail(),
		CreatedAt:         time.Now(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.PasswordChangedAt)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
