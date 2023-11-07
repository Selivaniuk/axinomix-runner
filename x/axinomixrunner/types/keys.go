package types

const (
	// ModuleName defines the module name
	ModuleName = "axinomixrunner"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_axinomixrunner"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	RaceKey      = "Race/value/"
	RaceCountKey = "Race/count/"
)

const (
	LeaderboardKey      = "Leaderboard/value/"
	LeaderboardCountKey = "Leaderboard/count/"
)

const (
	StartRaceEventType       = "start-race"
	StartRaceEventId         = "id"
	StartRaceEventMultiplier = "multiplier "
	StartRaceEventStartTime  = "start-time"
	StartRaceEventState      = "state"
)
const (
	EndRaceEventType  = "end-race"
	EndRaceEventId    = "id"
	EndRaceEventCoins = "coins "
)
