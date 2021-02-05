package keeper

import (
	"github.com/faddat/investo-demo/x/investodemo/types"
)

var _ types.QueryServer = Keeper{}
