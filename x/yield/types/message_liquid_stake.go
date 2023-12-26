package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLiquidStake = "liquid_stake"

var _ sdk.Msg = &MsgLiquidStake{}

func NewMsgLiquidStake(creator string, denom string, amount uint64) *MsgLiquidStake {
	return &MsgLiquidStake{
		Creator: creator,
		Denom:   denom,
		Amount:  amount,
	}
}

func (msg *MsgLiquidStake) Route() string {
	return RouterKey
}

func (msg *MsgLiquidStake) Type() string {
	return TypeMsgLiquidStake
}

func (msg *MsgLiquidStake) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLiquidStake) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLiquidStake) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
