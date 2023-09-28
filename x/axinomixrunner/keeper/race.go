package keeper

import (
	"encoding/binary"

	"axinomix-runner/x/axinomixrunner/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetRaceCount get the total number of race
func (k Keeper) GetRaceCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RaceCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetRaceCount set the total number of race
func (k Keeper) SetRaceCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.RaceCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendRace appends a race in the store with a new id and update the count
func (k Keeper) AppendRace(
	ctx sdk.Context,
	race types.Race,
) uint64 {
	// Create the race
	count := k.GetRaceCount(ctx)

	// Set the ID of the appended value
	race.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RaceKey))
	appendedValue := k.cdc.MustMarshal(&race)
	store.Set(GetRaceIDBytes(race.Id), appendedValue)

	// Update race count
	k.SetRaceCount(ctx, count+1)

	return count
}

// SetRace set a specific race in the store
func (k Keeper) SetRace(ctx sdk.Context, race types.Race) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RaceKey))
	b := k.cdc.MustMarshal(&race)
	store.Set(GetRaceIDBytes(race.Id), b)
}

// GetRace returns a race from its id
func (k Keeper) GetRace(ctx sdk.Context, id uint64) (val types.Race, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RaceKey))
	b := store.Get(GetRaceIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveRace removes a race from the store
func (k Keeper) RemoveRace(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RaceKey))
	store.Delete(GetRaceIDBytes(id))
}

// GetAllRace returns all race
func (k Keeper) GetAllRace(ctx sdk.Context) (list []types.Race) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RaceKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Race
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetRaceIDBytes returns the byte representation of the ID
func GetRaceIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRaceIDFromBytes returns ID in uint64 format from a byte array
func GetRaceIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
