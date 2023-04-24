package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomPassword(13)
	hashedPassword, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	err = CheckPassword(password, hashedPassword)
	require.NoError(t, err)

	wrongPassword := RandomPassword(10)
	err = CheckPassword(wrongPassword, hashedPassword)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	err = CheckPassword(password, hashedPassword2)
	require.NoError(t, err)
	require.NotEqual(t, hashedPassword, hashedPassword2)
}
