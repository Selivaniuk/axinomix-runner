package keeper

import (
	"axinomix-runner/x/axinomixrunner/types"
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

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
	if !bet.Amount.GTE(minBet) {
		return nil, fmt.Errorf("min bet is %s", sdk.NewCoin(bet.Denom, minBet).String())
	}
	maxBet := sdk.NewInt(1000)
	if !bet.Amount.LTE(maxBet) {
		return nil, fmt.Errorf("max bet is %s", sdk.NewCoin(bet.Denom, maxBet).String())
	}

	if !k.bankKeeper.HasBalance(ctx, player, bet) {
		return nil, fmt.Errorf("not enough tokens")
	}

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, player, types.ModuleName, sdk.NewCoins(bet))
	if sdkError != nil {
		return nil, sdkError
	}
	multiplier, err := sdk.NewDecFromInt(bet.Amount).Quo(sdk.NewDec(100)).Float64()
	var race = types.Race{
		PlayerAddress: player.String(),
		Bet:           bet.String(),
		Multiplier:    float32(multiplier),
		StartTime:     uint64(ctx.BlockHeader().Time.UnixMilli()),
		EndTime:       0,
		CoinsEarned:   sdk.NewCoin(bet.Denom, sdk.NewInt(0)).String(),
		Score:         0,
		State:         "active",
	}
	raceID := k.AppendRace(ctx, race)
	fmt.Println("Race", raceID, "successfully created")
	return &types.MsgStartRaceResponse{Id: raceID}, nil
}
