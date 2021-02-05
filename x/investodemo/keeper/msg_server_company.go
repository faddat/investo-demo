package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/faddat/investo-demo/x/investodemo/types"
)

func (k msgServer) CreateCompany(goCtx context.Context, msg *types.MsgCreateCompany) (*types.MsgCreateCompanyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendCompany(ctx, *msg)

	return &types.MsgCreateCompanyResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateCompany(goCtx context.Context, msg *types.MsgUpdateCompany) (*types.MsgUpdateCompanyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var company = types.Company{
		Creator:  msg.Creator,
		Id:       msg.Id,
		Mktgroi:  msg.Mktgroi,
		Logo:     msg.Logo,
		Netflow:  msg.Netflow,
		Industry: msg.Industry,
	}

	// Checks that the element exists
	if !k.HasCompany(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetCompanyOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetCompany(ctx, company)

	return &types.MsgUpdateCompanyResponse{}, nil
}

func (k msgServer) DeleteCompany(goCtx context.Context, msg *types.MsgDeleteCompany) (*types.MsgDeleteCompanyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasCompany(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetCompanyOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveCompany(ctx, msg.Id)

	return &types.MsgDeleteCompanyResponse{}, nil
}
