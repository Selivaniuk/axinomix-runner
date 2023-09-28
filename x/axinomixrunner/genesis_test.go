package axinomixrunner_test

import (
	"testing"

	keepertest "axinomix-runner/testutil/keeper"
	"axinomix-runner/testutil/nullify"
	"axinomix-runner/x/axinomixrunner"
	"axinomix-runner/x/axinomixrunner/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		RaceList: []types.Race{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		RaceCount: 2,
		LeaderboardList: []types.Leaderboard{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		LeaderboardCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AxinomixrunnerKeeper(t)
	axinomixrunner.InitGenesis(ctx, *k, genesisState)
	got := axinomixrunner.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.RaceList, got.RaceList)
	require.Equal(t, genesisState.RaceCount, got.RaceCount)
	require.ElementsMatch(t, genesisState.LeaderboardList, got.LeaderboardList)
	require.Equal(t, genesisState.LeaderboardCount, got.LeaderboardCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
