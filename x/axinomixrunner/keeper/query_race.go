package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types/query"

	"axinomix-runner/x/axinomixrunner/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RaceAll(goCtx context.Context, req *types.QueryAllRaceRequest) (*types.QueryAllRaceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var races []types.Race
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	raceStore := prefix.NewStore(store, types.KeyPrefix(types.RaceKey))

	var pageRes *query.PageResponse
	var err error
	switch {
	case len(req.Address) > 0:
		pageRes, err = query.Paginate(raceStore, req.Pagination, func(key []byte, value []byte) error {
			var race types.Race
			if err := k.cdc.Unmarshal(value, &race); err != nil {
				return err
			}
			has := race.PlayerAddress == req.Address
			if has {
				races = append(races, race)
			}
			return nil
		})
		pageRes.Total = uint64(len(races))
	default:
		pageRes, err = query.Paginate(raceStore, req.Pagination, func(key []byte, value []byte) error {
			var race types.Race
			if err := k.cdc.Unmarshal(value, &race); err != nil {
				return err
			}
			races = append(races, race)
			return nil
		})
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryAllRaceResponse{Race: races, Pagination: pageRes}, nil
}

func (k Keeper) Race(goCtx context.Context, req *types.QueryGetRaceRequest) (*types.QueryGetRaceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	race, found := k.GetRace(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetRaceResponse{Race: race}, nil
}
