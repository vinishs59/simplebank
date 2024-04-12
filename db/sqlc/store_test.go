package db

import (
	"context"
	"testing"

	_ "github.com/lib/pq"

	"github.com/stretchr/testify/require"
)

func TestTransferTx (t *testing.T) {

	store := NewStore(testDb)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 5 

	amount := int64(10)

	errs := make( chan error)
	results := make( chan TransferTxResult)

	for i:=0 ; i < n ; i ++ {
		go func() {
			result,err := store.transferTx(context.Background(),TransferTxParams{
				FromAccountID: int64(account1.UserID),
				ToAccountID: int64(account2.UserID),
				Amount: amount,


			})

			errs <-err
			results <- result


		}()
	} 

	for i:=0 ; i < n ; i ++ {
		err := <-errs
		require.NoError(t,err)

		result := <- results
		require.NotEmpty(t,result)


	}
 
}