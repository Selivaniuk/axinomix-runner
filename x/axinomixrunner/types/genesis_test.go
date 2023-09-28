package types_test

import (
	"testing"

	"axinomix-runner/x/axinomixrunner/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated race",
			genState: &types.GenesisState{
				RaceList: []types.Race{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid race count",
			genState: &types.GenesisState{
				RaceList: []types.Race{
					{
						Id: 1,
					},
				},
				RaceCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated leaderboard",
			genState: &types.GenesisState{
				LeaderboardList: []types.Leaderboard{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid leaderboard count",
			genState: &types.GenesisState{
				LeaderboardList: []types.Leaderboard{
					{
						Id: 1,
					},
				},
				LeaderboardCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
