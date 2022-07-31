package db

import (
	"context"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.CreatedAt)
	require.NotZero(t, transfer.ID)
	require.WithinDuration(t, time.Now(), transfer.CreatedAt, time.Second)
	return transfer
}
func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)

	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.CreatedAt, transfer2.CreatedAt)
	require.Equal(t, transfer1.ID, transfer2.ID)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}
	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}
	//创建10个，跳过5个还有5个，那就说明ok
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 5, len(transfers))
	var CSTZone = time.FixedZone("CST", 8*3600) //东八区
	// var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海 windows没有安装go会获取失败，所以使用上面这个
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)

		require.WithinDuration(t, transfer.CreatedAt.In(CSTZone), time.Now().In(CSTZone), time.Second*5)
	}
}

func TestListTransfersFromAccount(t *testing.T) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		arg := CreateTransferParams{
			FromAccountID: fromAccount.ID,
			ToAccountID:   toAccount.ID,
			Amount:        util.RandomMoney(),
		}

		transfer, err := testQueries.CreateTransfer(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, transfer)

		require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
		require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
		require.Equal(t, arg.Amount, transfer.Amount)
		require.NotZero(t, transfer.CreatedAt)
		require.NotZero(t, transfer.ID)
		require.WithinDuration(t, time.Now(), transfer.CreatedAt, time.Second)
	}

	//check from_account_id
	arg := ListTransfersFromAccountParams{
		Limit:         5,
		Offset:        5,
		FromAccountID: fromAccount.ID,
	}
	//创建10个，跳过5个还有5个，那就说明ok
	transfers, err := testQueries.ListTransfersFromAccount(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 5, len(transfers))
	var CSTZone = time.FixedZone("CST", 8*3600) //东八区
	// var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海 windows没有安装go会获取失败，所以使用上面这个
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
		require.Equal(t, toAccount.ID, transfer.ToAccountID)
		require.WithinDuration(t, transfer.CreatedAt.In(CSTZone), time.Now().In(CSTZone), time.Second*5)
	}

	//check to_account_id
	arg2 := ListTransfersToAccountParams{
		Limit:       5,
		Offset:      5,
		ToAccountID: toAccount.ID,
	}
	//创建10个，跳过5个还有5个，那就说明ok
	transfers, err = testQueries.ListTransfersToAccount(context.Background(), arg2)
	require.NoError(t, err)
	require.Equal(t, 5, len(transfers))
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
		require.Equal(t, toAccount.ID, transfer.ToAccountID)
		require.WithinDuration(t, transfer.CreatedAt.In(CSTZone), time.Now().In(CSTZone), time.Second*5)
	}
}
