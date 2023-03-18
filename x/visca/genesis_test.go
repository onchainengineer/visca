package visca_test

import (
	"testing"

	keepertest "github.com/onchainengineer/visca/testutil/keeper"
	"github.com/onchainengineer/visca/testutil/nullify"
	"github.com/onchainengineer/visca/x/visca"
	"github.com/onchainengineer/visca/x/visca/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ViscaKeeper(t)
	visca.InitGenesis(ctx, *k, genesisState)
	got := visca.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
