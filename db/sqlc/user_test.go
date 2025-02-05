package db

import (
	"context"
	"testing"

	"github.com/sonnyvictok/bank_test/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
func createRandomUser(t *testing.T) User {
	hashPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)

	return user
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)

	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Username, user2.Username)

}
