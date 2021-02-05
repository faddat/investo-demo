package investodemo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/faddat/investo-demo/x/investodemo/keeper"
	"github.com/faddat/investo-demo/x/investodemo/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the company
	for _, elem := range genState.CompanyList {
		k.SetCompany(ctx, *elem)
	}

	// Set company count
	k.SetCompanyCount(ctx, int64(len(genState.CompanyList)))

}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all company
	companyList := k.GetAllCompany(ctx)
	for _, elem := range companyList {
		elem := elem
		genesis.CompanyList = append(genesis.CompanyList, &elem)
	}

	return genesis
}
