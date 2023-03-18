package keeper_test

import (
	"testing"

	testkeeper "github.com/onchainengineer/visca/testutil/keeper"
	"github.com/onchainengineer/visca/x/visca/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ViscaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
