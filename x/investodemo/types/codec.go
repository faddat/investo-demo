package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateCompany{}, "investodemo/CreateCompany", nil)
	cdc.RegisterConcrete(&MsgUpdateCompany{}, "investodemo/UpdateCompany", nil)
	cdc.RegisterConcrete(&MsgDeleteCompany{}, "investodemo/DeleteCompany", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCompany{},
		&MsgUpdateCompany{},
		&MsgDeleteCompany{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
