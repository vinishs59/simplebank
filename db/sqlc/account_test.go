package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/vinishs59/simplebank/util"
)

var testStore *Store = NewStore(testDb)

func createRandomAccount(t *testing.T) Account {
	user := "test " //createRandomUser(t)

	arg := CreateAccountParams{
		OwnerName: user, //user.Username,
		Balance:   util.RandomMoney(),
		Currency:  sql.NullString{String: "USD", Valid: true},
	}

	//testStore := NewStore(testDb)

	account, err := testStore.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.OwnerName, account.OwnerName)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.UserID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testStore.GetAccount(context.Background(), account1.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.UserID, account2.UserID)
	require.Equal(t, account1.OwnerName, account2.OwnerName)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	if account1.CreatedAt.Valid && account2.CreatedAt.Valid {
		// Convert sql.NullTime to time.Time
		time1 := account1.CreatedAt.Time
		time2 := account2.CreatedAt.Time

		// Use time1 and time2 with require.WithinDuration
		require.WithinDuration(t, time1, time2, time.Second)
	}
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		UserID:  account1.UserID,
		Balance: util.RandomMoney(),
	}

	err := testStore.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	//require.NotEmpty(t, account2)

	// require.Equal(t, account1.UserID, account2.ID)
	// require.Equal(t, account1.OwnerName, account2.Owner)
	// require.Equal(t, arg.Balance, account2.Balance)
	// require.Equal(t, account1.Currency, account2.Currency)
	// if account1.CreatedAt.Valid && account2.CreatedAt.Valid {
	// 	// Convert sql.NullTime to time.Time
	// 	time1 := account1.CreatedAt.Time
	// 	time2 := account2.CreatedAt.Time

	// 	// Use time1 and time2 with require.WithinDuration
	// 	require.WithinDuration(t, time1, time2, time.Second)
	// }
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testStore.DeleteAccount(context.Background(), account1.UserID)
	require.NoError(t, err)

	account2, err := testStore.GetAccount(context.Background(), account1.UserID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountParams{
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testStore.ListAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.OwnerName, account.OwnerName)
	}
}
