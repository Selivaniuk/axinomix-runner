package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		RaceList:        []Race{},
		LeaderboardList: []Leaderboard{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in race
	raceIdMap := make(map[uint64]bool)
	raceCount := gs.GetRaceCount()
	for _, elem := range gs.RaceList {
		if _, ok := raceIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for race")
		}
		if elem.Id >= raceCount {
			return fmt.Errorf("race id should be lower or equal than the last id")
		}
		raceIdMap[elem.Id] = true
	}
	// Check for duplicated ID in leaderboard
	leaderboardIdMap := make(map[uint64]bool)
	leaderboardCount := gs.GetLeaderboardCount()
	for _, elem := range gs.LeaderboardList {
		if _, ok := leaderboardIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for leaderboard")
		}
		if elem.Id >= leaderboardCount {
			return fmt.Errorf("leaderboard id should be lower or equal than the last id")
		}
		leaderboardIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
