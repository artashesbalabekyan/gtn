package gtn_test

import (
	"testing"

	keepertest "GTN/testutil/keeper"
	"GTN/testutil/nullify"
	"GTN/x/gtn"
	"GTN/x/gtn/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GtnKeeper(t)
	gtn.InitGenesis(ctx, *k, genesisState)
	got := gtn.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
