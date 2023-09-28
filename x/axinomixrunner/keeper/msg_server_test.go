package keeper_test

import (
	"context"
	"testing"

	keepertest "axinomix-runner/testutil/keeper"
	"axinomix-runner/x/axinomixrunner/keeper"
	"axinomix-runner/x/axinomixrunner/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AxinomixrunnerKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
