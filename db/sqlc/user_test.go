package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vinishs59/simplebank/util"
)

//var testStore *Store = NewStore(testDb)

func createRandomUser(t *testing.T) User {
	random_user := createRandomAccount(t)

	arg := CreateUserParams{
		Username:       random_user.OwnerName, //user.Username,
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	//testStore := NewStore(testDb)

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.FullName)

	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.PasswordChangedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	user2, err := testStore.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2.CreatedAt)

	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.FullName, user2.FullName)
	require.Equal(t, user.Email, user2.Email)
	//require.Equal(t, account1.Currency, user2.Currency)

}
