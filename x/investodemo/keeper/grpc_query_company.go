package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/faddat/investo-demo/x/investodemo/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CompanyAll(c context.Context, req *types.QueryAllCompanyRequest) (*types.QueryAllCompanyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var companys []*types.Company
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	companyStore := prefix.NewStore(store, types.KeyPrefix(types.CompanyKey))

	pageRes, err := query.Paginate(companyStore, req.Pagination, func(key []byte, value []byte) error {
		var company types.Company
		if err := k.cdc.UnmarshalBinaryBare(value, &company); err != nil {
			return err
		}

		companys = append(companys, &company)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCompanyResponse{Company: companys, Pagination: pageRes}, nil
}

func (k Keeper) Company(c context.Context, req *types.QueryGetCompanyRequest) (*types.QueryGetCompanyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var company types.Company
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompanyKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.CompanyKey+req.Id)), &company)

	return &types.QueryGetCompanyResponse{Company: &company}, nil
}
