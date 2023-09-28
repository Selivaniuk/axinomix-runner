package keeper

import (
	"context"

	"axinomix-runner/x/axinomixrunner/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) LeaderboardAll(goCtx context.Context, req *types.QueryAllLeaderboardRequest) (*types.QueryAllLeaderboardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var leaderboards []types.Leaderboard
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	leaderboardStore := prefix.NewStore(store, types.KeyPrefix(types.LeaderboardKey))

	pageRes, err := query.Paginate(leaderboardStore, req.Pagination, func(key []byte, value []byte) error {
		var leaderboard types.Leaderboard
		if err := k.cdc.Unmarshal(value, &leaderboard); err != nil {
			return err
		}

		leaderboards = append(leaderboards, leaderboard)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLeaderboardResponse{Leaderboard: leaderboards, Pagination: pageRes}, nil
}

func (k Keeper) Leaderboard(goCtx context.Context, req *types.QueryGetLeaderboardRequest) (*types.QueryGetLeaderboardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	leaderboard, found := k.GetLeaderboardByAddress(ctx, req.Address)
	//leaderboard, found := k.GetLeaderboard(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetLeaderboardResponse{Leaderboard: leaderboard}, nil
}
