package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStartRace = "start_race"

var _ sdk.Msg = &MsgStartRace{}

func NewMsgStartRace(creator string, amount string, mode string) *MsgStartRace {
	return &MsgStartRace{
		Creator: creator,
		Amount:  amount,
		Mode:    mode,
	}
}

func (msg *MsgStartRace) Route() string {
	return RouterKey
}

func (msg *MsgStartRace) Type() string {
	return TypeMsgStartRace
}

func (msg *MsgStartRace) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStartRace) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStartRace) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
