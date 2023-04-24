package db

import (
	"context"
	"testing"
	"time"

	"github.com/dipu626/simple-bank/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	password := util.RandomPassword(12)
	hashedPassword, err := util.HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	arg := CreateUserParams{
		Username:          util.RandomOwner(),
		FullName:          util.RandomOwner(),
		HashedPassword:    hashedPassword,
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

	// require.NotZero(t, user.PasswordChangedAt)
	// require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	user1, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user1)
	require.Equal(t, user.Username, user1.Username)
	require.Equal(t, user.FullName, user1.FullName)
	require.Equal(t, user.Email, user1.Email)

	// require.WithinDuration(t, user.CreatedAt, user1.CreatedAt, time.Second)
}
