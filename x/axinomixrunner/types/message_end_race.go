package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgEndRace = "end_race"

var _ sdk.Msg = &MsgEndRace{}

func NewMsgEndRace(creator string, id uint64, coins uint64, score uint64) *MsgEndRace {
	return &MsgEndRace{
		Creator: creator,
		Id:      id,
		Coins:   coins,
		Score:   score,
	}
}

func (msg *MsgEndRace) Route() string {
	return RouterKey
}

func (msg *MsgEndRace) Type() string {
	return TypeMsgEndRace
}

func (msg *MsgEndRace) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgEndRace) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEndRace) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
