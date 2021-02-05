package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/faddat/investo-demo/x/investodemo/types"
	"strconv"
)

// GetCompanyCount get the total number of company
func (k Keeper) GetCompanyCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompanyCountKey))
	byteKey := types.KeyPrefix(types.CompanyCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetCompanyCount set the total number of company
func (k Keeper) SetCompanyCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompanyCountKey))
	byteKey := types.KeyPrefix(types.CompanyCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// AppendCompany appends a company in the store with a new id and update the count
func (k Keeper) AppendCompany(ctx sdk.Context, msg types.MsgCreateCompany) string {
	// Create the company
	count := k.GetCompanyCount(ctx)
	id := strconv.FormatInt(count, 10)
	var company = types.Company{
		Creator:  msg.Creator,
		Id:       id,
		Mktgroi:  msg.Mktgroi,
		Logo:     msg.Logo,
		Netflow:  msg.Netflow,
		Industry: msg.Industry,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompanyKey))
	key := types.KeyPrefix(types.CompanyKey + company.Id)
	value := k.cdc.MustMarshalBinaryBare(&company)
	store.Set(key, value)

	// Update company count
	k.SetCompanyCount(ctx, count+1)

	return id
}

// SetCompany set a specific company in the store
func (k Keeper) SetCompany(ctx sdk.Context, company types.Company) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompanyKey))
	b := k.cdc.MustMarshalBinaryBare(&company)
	store.Set(types.KeyPrefix(types.CompanyKey+company.Id), b)
}

// GetCompany returns a company from its id
func (k Keeper) GetCompany(ctx sdk.Context, key string) types.Company {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompanyKey))
	var company types.Company
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.CompanyKey+key)), &company)
	return company
}

// HasCompany checks if the company exists in the store
func (k Keeper) HasCompany(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompanyKey))
	return store.Has(types.KeyPrefix(types.CompanyKey + id))
}

// GetCompanyOwner returns the creator of the company
func (k Keeper) GetCompanyOwner(ctx sdk.Context, key string) string {
	return k.GetCompany(ctx, key).Creator
}

// DeleteCompany removes a company from the store
func (k Keeper) RemoveCompany(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompanyKey))
	store.Delete(types.KeyPrefix(types.CompanyKey + key))
}

// GetAllCompany returns all company
func (k Keeper) GetAllCompany(ctx sdk.Context) (msgs []types.Company) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CompanyKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.CompanyKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Company
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
