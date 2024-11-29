package db

import (
	"context"
	"testing"
	"time"

	util "example.com/banking/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPsw, err := util.HashPassword(util.RandomCurrency())
	args := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPsw,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	user, err := TestQueries.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.HashedPassword, user.HashedPassword)
	require.Equal(t, args.FullName, user.FullName)
	require.Equal(t, args.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := TestQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestLoginUser(t *testing.T) {
	password := util.RandomString()
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	args := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := TestQueries.CreateUser(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	isPasswordMatch := util.IsPasswordMatch(password, user.HashedPassword)
	require.True(t, isPasswordMatch, "The password should match")

	isPasswordMatch = util.IsPasswordMatch("wrongpassword", user.HashedPassword)
	require.False(t, isPasswordMatch, "The password should not match")
}
