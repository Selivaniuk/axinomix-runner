package keeper_test

import (
	"testing"

	keepertest "axinomix-runner/testutil/keeper"
	"axinomix-runner/testutil/nullify"
	"axinomix-runner/x/axinomixrunner/keeper"
	"axinomix-runner/x/axinomixrunner/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNLeaderboard(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Leaderboard {
	items := make([]types.Leaderboard, n)
	for i := range items {
		items[i].Id = keeper.AppendLeaderboard(ctx, items[i])
	}
	return items
}

func TestLeaderboardGet(t *testing.T) {
	keeper, ctx := keepertest.AxinomixrunnerKeeper(t)
	items := createNLeaderboard(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetLeaderboard(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestLeaderboardRemove(t *testing.T) {
	keeper, ctx := keepertest.AxinomixrunnerKeeper(t)
	items := createNLeaderboard(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveLeaderboard(ctx, item.Id)
		_, found := keeper.GetLeaderboard(ctx, item.Id)
		require.False(t, found)
	}
}

func TestLeaderboardGetAll(t *testing.T) {
	keeper, ctx := keepertest.AxinomixrunnerKeeper(t)
	items := createNLeaderboard(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllLeaderboard(ctx)),
	)
}

func TestLeaderboardCount(t *testing.T) {
	keeper, ctx := keepertest.AxinomixrunnerKeeper(t)
	items := createNLeaderboard(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetLeaderboardCount(ctx))
}
