package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	accountF := createRandomAccount(t)
	accountT := createRandomAccount(t)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	result := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			res, err := testStore.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: accountF.ID,
				ToAccountID:   accountT.ID,
				Amount:        amount,
			})
			errs <- err
			result <- res
		}()
	}

	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		res := <-result
		require.NotEmpty(t, res)

		transfer := res.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.ToAccountID, accountT.ID)
		require.Equal(t, transfer.FromAccountID, accountF.ID)
		require.Equal(t, transfer.Amount, amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = testStore.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := res.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, accountF.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = testStore.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := res.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, accountT.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = testStore.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		//check accounts
		fromAccount := res.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, fromAccount.ID, accountF.ID)

		toAccount := res.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, toAccount.ID, accountT.ID)

		//check account balance

		fromDif := accountF.Balance - fromAccount.Balance
		toDif := toAccount.Balance - accountT.Balance
		require.Equal(t, fromDif, toDif)
		require.True(t, fromDif > 0)
		require.True(t, fromDif%amount == 0)

		k := int(fromDif / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	updatedAccount1, err := testStore.GetAccount(context.Background(), accountF.ID)
	require.NoError(t, err)

	updatedAccount2, err := testStore.GetAccount(context.Background(), accountT.ID)
	require.NoError(t, err)

	require.Equal(t, accountF.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, accountT.Balance+int64(n)*amount, updatedAccount2.Balance)
}
