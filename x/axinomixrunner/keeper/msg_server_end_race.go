package keeper

import (
	"bytes"
	"context"
	"fmt"

	"axinomix-runner/x/axinomixrunner/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EndRace(goCtx context.Context, msg *types.MsgEndRace) (*types.MsgEndRaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	race, found := k.GetRace(ctx, msg.Id)

	if !found {
		return nil, fmt.Errorf("race does not exist")
	}
	if race.State != "active" {
		return nil, fmt.Errorf("cannot end this race")
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	player, err := sdk.AccAddressFromBech32(race.PlayerAddress)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(creator, player) {
		return nil, fmt.Errorf("can only be end by the creator")
	}

	bet, err := sdk.ParseCoinNormalized(race.Bet)
	if err != nil {
		return nil, err
	}

	var winAmount = sdk.NewCoin(bet.Denom, sdk.NewInt(0))
	var sdkError error
	if msg.Coins >= race.NeedToCollectCoins {
		winAmount = sdk.NewCoin(bet.Denom, bet.Amount.MulRaw(2))
		sdkError = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, player, sdk.NewCoins(winAmount))
	}
	if sdkError != nil {
		return nil, sdkError
	}

	race.WinAmount = winAmount.String()
	race.CoinsEarned = msg.Coins
	race.Score = msg.Score
	race.EndTime = uint64(ctx.BlockHeader().Time.UnixMilli())
	race.State = "finished"
	k.SetRace(ctx, race)
	k.UpdatePlayerRank(ctx, player.String(), msg.Score)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.EndRaceEventType,
			sdk.NewAttribute(types.EndRaceEventId, fmt.Sprint(race.Id)),
			sdk.NewAttribute(types.EndRaceEventWinAmount, fmt.Sprint(winAmount.String())),
		),
	)
	return &types.MsgEndRaceResponse{
		Id:        race.Id,
		WinAmount: winAmount.String(),
	}, nil
}
