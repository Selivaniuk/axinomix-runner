package keeper

import (
	"axinomix-runner/x/axinomixrunner/types"
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const R = 0.75
const CPM = 0.8
const Distance = 500
const NeedToCollect = uint64(Distance * CPM * R)

func (k msgServer) StartRace(goCtx context.Context, msg *types.MsgStartRace) (*types.MsgStartRaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	player, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	bet, err := sdk.ParseCoinNormalized(msg.Amount)
	if err != nil {
		return nil, err
	}

	minBet := sdk.NewInt(10)
	if bet.Amount.LT(minBet) {
		return nil, fmt.Errorf("min bet is %s", sdk.NewCoin(bet.Denom, minBet).String())
	}
	maxBet := sdk.NewInt(10000)
	if bet.Amount.GT(maxBet) {
		return nil, fmt.Errorf("max bet is %s", sdk.NewCoin(bet.Denom, maxBet).String())
	}

	if !k.bankKeeper.HasBalance(ctx, player, bet) {
		return nil, fmt.Errorf("not enough tokens")
	}

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, player, types.ModuleName, sdk.NewCoins(bet))
	if sdkError != nil {
		return nil, sdkError
	}

	var startTime = uint64(ctx.BlockHeader().Time.UnixMilli())
	var race = types.Race{
		PlayerAddress:      player.String(),
		Mode:               msg.Mode,
		Bet:                bet.String(),
		NeedToCollectCoins: NeedToCollect,
		StartTime:          startTime,
		EndTime:            0,
		CoinsEarned:        0,
		Distance:           Distance,
		Score:              0,
		State:              "active",
	}
	raceID := k.AppendRace(ctx, race)
	fmt.Println("Race", raceID, "successfully created")

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.StartRaceEventType,
			sdk.NewAttribute(types.StartRaceEventId, fmt.Sprint(raceID)),
			sdk.NewAttribute(types.StartRaceEventNeedToCollectCoins, fmt.Sprint(NeedToCollect)),
			sdk.NewAttribute(types.StartRaceEventDistance, fmt.Sprint(Distance)),
			sdk.NewAttribute(types.StartRaceEventBet, fmt.Sprint(bet.String())),
		),
	)

	return &types.MsgStartRaceResponse{Id: raceID}, nil
}
