package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	util "example.com/banking/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	args := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := TestQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2, err := TestQueries.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Owner, acc2.Owner)
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.Equal(t, acc1.Balance, acc2.Balance)
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	acc1 := createRandomAccount(t)

	args := UpdateAccountParams{
		ID:      acc1.ID,
		Balance: util.RandomMoney(),
	}

	acc2, err := TestQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Owner, acc2.Owner)
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.NotEqual(t, acc1.Balance, acc2.Balance)
}

func TestDeleteAccount(t *testing.T) {
	acc1 := createRandomAccount(t)

	err := TestQueries.DeleteAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	acc2, err := TestQueries.GetAccount(context.Background(), acc1.ID)
	require.Error(t, err)
	require.Empty(t, acc2)

	require.EqualError(t, err, "sql: no rows in result set")
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	createdAccsLen := 10
	for i := 0; i < createdAccsLen; i++ {
		lastAccount = createRandomAccount(t)
	}
	fmt.Println(lastAccount)
	args := ListAllAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  1,
		Offset: 1,
	}
	fmt.Println(args)
	accs, err := TestQueries.ListAllAccounts(context.Background(), args)
	fmt.Println(err)
	fmt.Println(accs)
	require.NoError(t, err)
	// require.NotEmpty(t, accs)
	// require.Len(t, accs, 1)

	// for _, account := range accs {
	// 	// require.NotEmpty(t, account)
	// 	// require.Equal(t, lastAccount.Owner, account.Owner)
	// }
}
