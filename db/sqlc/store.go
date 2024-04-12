package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store defines all functions to execute db queries and transactions
type Store struct {
	*Queries 
	db *sql.DB
// 	Querier
// 	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
// 	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
// 	VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error)
// 
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
// 	connPool *pgxpool.Pool
// 	*Queries
 }


 func NewStore (db *sql.DB) *Store{
	return &Store {
		db : db,
		Queries: New(db),
	}
 }

// // NewStore creates a new store
// func NewStore(connPool *pgxpool.Pool) Store {
// 	// return &SQLStore{
// 	// 	connPool: connPool,
// 	// 	Queries:  New(connPool),
// 	// }
// }

func (store *Store) execTX (ctx context.Context , fn func (*Queries) error) error {
	tx , err := store.db.BeginTx(ctx,nil)
	if err !=nil {
		return nil
	}
	q:=New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback();  rbErr != nil{
			return fmt.Errorf("tx err : %v , rb err :%v", rbErr,err)
		}

        return err

	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID  int64 `json:"from_account"`	
	ToAccountID  int64 `json:"to_account"`
	Amount int64 `json:"amount"`


}

//transfer tx to perform the xfer function 

type TransferTxResult struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount Account `json:"to_account"`
	FromEntry Entry `json:"from_entry"`
	ToEntry Entry `json:"to_entry"`
}


func (store *Store ) transferTx ( ctx context.Context,arg TransferTxParams)(TransferTxResult,error){

	var result TransferTxResult

	err := store.execTX(ctx,func(q *Queries)error {

		var err error 

		result.Transfer , err = q.CreateTransfer(ctx , CreateTransferParams{
			FromAccount: int32(arg.FromAccountID),
			ToAccount: int32(arg.ToAccountID),
			Amount: arg.Amount,
		})

		if err != nil {
			return err
		}

		result.FromEntry , err = q.CreateEntry(ctx , CreateEntryParams{
			AccountID: int32(arg.FromAccountID),
			Amount: -arg.Amount,
		})

		if err != nil {
			return err
		}

		result.ToEntry , err = q.CreateEntry(ctx , CreateEntryParams{
			AccountID: int32(arg.ToAccountID),
			Amount: arg.Amount,
		})

		if err != nil {
			return err
		}


		return nil

	})

	return result ,err


}