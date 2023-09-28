package keeper

import (
	"encoding/binary"
	"sort"

	"axinomix-runner/x/axinomixrunner/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetLeaderboardCount get the total number of leaderboard
func (k Keeper) GetLeaderboardCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LeaderboardCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLeaderboardCount set the total number of leaderboard
func (k Keeper) SetLeaderboardCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.LeaderboardCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendLeaderboard appends a leaderboard in the store with a new id and update the count
func (k Keeper) AppendLeaderboard(
	ctx sdk.Context,
	leaderboard types.Leaderboard,
) uint64 {
	// Create the leaderboard
	count := k.GetLeaderboardCount(ctx)

	// Set the ID of the appended value
	leaderboard.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LeaderboardKey))
	appendedValue := k.cdc.MustMarshal(&leaderboard)
	store.Set(GetLeaderboardIDBytes(leaderboard.Id), appendedValue)

	// Update leaderboard count
	k.SetLeaderboardCount(ctx, count+1)

	return count
}

// SetLeaderboard set a specific leaderboard in the store
func (k Keeper) SetLeaderboard(ctx sdk.Context, leaderboard types.Leaderboard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LeaderboardKey))
	b := k.cdc.MustMarshal(&leaderboard)
	store.Set(GetLeaderboardIDBytes(leaderboard.Id), b)
}

func (k Keeper) SetLeaderboardList(ctx sdk.Context, leaderboardList []types.Leaderboard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LeaderboardKey))
	for _, leaderboard := range leaderboardList {
		b := k.cdc.MustMarshal(&leaderboard)
		store.Set(GetLeaderboardIDBytes(leaderboard.Id), b)
	}
}

// GetLeaderboard returns a leaderboard from its id
func (k Keeper) GetLeaderboard(ctx sdk.Context, id uint64) (val types.Leaderboard, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LeaderboardKey))
	b := store.Get(GetLeaderboardIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLeaderboard removes a leaderboard from the store
func (k Keeper) RemoveLeaderboard(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LeaderboardKey))
	store.Delete(GetLeaderboardIDBytes(id))
}

// GetAllLeaderboard returns all leaderboard
func (k Keeper) GetAllLeaderboard(ctx sdk.Context) (list []types.Leaderboard) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LeaderboardKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Leaderboard
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
func (k Keeper) GetLeaderboardByAddress(ctx sdk.Context, address string) (val types.Leaderboard, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LeaderboardKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var value types.Leaderboard
		k.cdc.MustUnmarshal(iterator.Value(), &value)
		if value.PlayerAddress == address {
			return value, true
		}
	}

	return val, false
}

// GetLeaderboardIDBytes returns the byte representation of the ID
func GetLeaderboardIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetLeaderboardIDFromBytes returns ID in uint64 format from a byte array
func GetLeaderboardIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
func (k Keeper) UpdatePlayerRank(ctx sdk.Context, address string, score uint64) {
	leaderboardList := k.GetAllLeaderboard(ctx)
	leader, found := k.GetLeaderboardByAddress(ctx, address)
	if found && leader.Score >= score {
		return
	}
	if !found {
		newLeader := types.Leaderboard{
			PlayerAddress: address,
			Score:         score,
		}
		newId := k.AppendLeaderboard(ctx, newLeader)
		newLeader.Id = newId
		leaderboardList = append(leaderboardList, newLeader)
	} else {
		var indexToUpdate = -1
		for i, leaderboard := range leaderboardList {
			if leaderboard.Id == leader.Id {
				indexToUpdate = i
				break
			}
		}
		if indexToUpdate != -1 {
			leaderboardList[indexToUpdate].Score = score
		}
	}
	sort.Slice(leaderboardList, func(i, j int) bool {
		return leaderboardList[i].Score > leaderboardList[j].Score
	})
	for i := range leaderboardList {
		leaderboardList[i].PlayerRank = uint64(i + 1)
	}
	k.SetLeaderboardList(ctx, leaderboardList)
}
