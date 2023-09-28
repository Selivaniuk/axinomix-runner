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

func createNRace(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Race {
	items := make([]types.Race, n)
	for i := range items {
		items[i].Id = keeper.AppendRace(ctx, items[i])
	}
	return items
}

func TestRaceGet(t *testing.T) {
	keeper, ctx := keepertest.AxinomixrunnerKeeper(t)
	items := createNRace(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetRace(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestRaceRemove(t *testing.T) {
	keeper, ctx := keepertest.AxinomixrunnerKeeper(t)
	items := createNRace(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRace(ctx, item.Id)
		_, found := keeper.GetRace(ctx, item.Id)
		require.False(t, found)
	}
}

func TestRaceGetAll(t *testing.T) {
	keeper, ctx := keepertest.AxinomixrunnerKeeper(t)
	items := createNRace(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRace(ctx)),
	)
}

func TestRaceCount(t *testing.T) {
	keeper, ctx := keepertest.AxinomixrunnerKeeper(t)
	items := createNRace(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetRaceCount(ctx))
}
