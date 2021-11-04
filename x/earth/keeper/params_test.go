package keeper_test

import (
	"testing"

	testkeeper "github.com/lubtd/earth/testutil/keeper"
	"github.com/lubtd/earth/x/earth/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.EarthKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
