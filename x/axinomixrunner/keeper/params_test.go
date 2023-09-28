package keeper_test

import (
	"testing"

	testkeeper "axinomix-runner/testutil/keeper"
	"axinomix-runner/x/axinomixrunner/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AxinomixrunnerKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
