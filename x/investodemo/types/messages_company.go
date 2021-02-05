package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateCompany{}

func NewMsgCreateCompany(creator string, mktgroi string, logo string, netflow string, industry string) *MsgCreateCompany {
	return &MsgCreateCompany{
		Creator:  creator,
		Mktgroi:  mktgroi,
		Logo:     logo,
		Netflow:  netflow,
		Industry: industry,
	}
}

func (msg *MsgCreateCompany) Route() string {
	return RouterKey
}

func (msg *MsgCreateCompany) Type() string {
	return "CreateCompany"
}

func (msg *MsgCreateCompany) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCompany) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCompany) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCompany{}

func NewMsgUpdateCompany(creator string, id string, mktgroi string, logo string, netflow string, industry string) *MsgUpdateCompany {
	return &MsgUpdateCompany{
		Id:       id,
		Creator:  creator,
		Mktgroi:  mktgroi,
		Logo:     logo,
		Netflow:  netflow,
		Industry: industry,
	}
}

func (msg *MsgUpdateCompany) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCompany) Type() string {
	return "UpdateCompany"
}

func (msg *MsgUpdateCompany) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCompany) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCompany) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateCompany{}

func NewMsgDeleteCompany(creator string, id string) *MsgDeleteCompany {
	return &MsgDeleteCompany{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteCompany) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCompany) Type() string {
	return "DeleteCompany"
}

func (msg *MsgDeleteCompany) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteCompany) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCompany) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
