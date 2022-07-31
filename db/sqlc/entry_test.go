package db

import (
	"context"
	"simple_bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account1 := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.CreatedAt)
	require.NotZero(t, entry.ID)
	require.WithinDuration(t, time.Now(), entry.CreatedAt, time.Second)
	return entry
}
func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)
	require.Equal(t, entry1.ID, entry2.ID)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}
	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	//创建10个，跳过5个还有5个，那就说明ok
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, entries)
	require.Equal(t, 5, len(entries))
	var CSTZone = time.FixedZone("CST", 8*3600) //东八区
	// var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海 windows没有安装go会获取失败，所以使用上面这个
	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.WithinDuration(t, entry.CreatedAt.In(CSTZone), time.Now().In(CSTZone), time.Second*5)
	}
}

func TestListEntriesByAccount(t *testing.T) {
	account := createRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	for i := 0; i < 10; i++ {

		entry, err := testQueries.CreateEntry(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, entry)

		require.Equal(t, arg.AccountID, entry.AccountID)
		require.Equal(t, arg.Amount, entry.Amount)
		require.NotZero(t, entry.CreatedAt)
		require.NotZero(t, entry.ID)
		require.WithinDuration(t, time.Now(), entry.CreatedAt, time.Second)
	}

	arg2 := ListEntriesByAccountParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	//check account_id
	//创建10个，跳过5个还有5个，那就说明ok
	entries, err := testQueries.ListEntriesByAccount(context.Background(), arg2)
	require.NoError(t, err)
	require.Equal(t, 5, len(entries))
	var CSTZone = time.FixedZone("CST", 8*3600) //东八区
	// var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海 windows没有安装go会获取失败，所以使用上面这个
	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
		require.Equal(t, arg.Amount, entry.Amount)
		require.WithinDuration(t, entry.CreatedAt.In(CSTZone), time.Now().In(CSTZone), time.Second*5)
	}
}
