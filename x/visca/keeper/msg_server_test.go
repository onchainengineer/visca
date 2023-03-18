package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/onchainengineer/visca/testutil/keeper"
	"github.com/onchainengineer/visca/x/visca/keeper"
	"github.com/onchainengineer/visca/x/visca/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ViscaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
