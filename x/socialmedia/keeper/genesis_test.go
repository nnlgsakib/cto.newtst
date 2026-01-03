package keeper_test

import (
	"testing"

	"nlg/x/socialmedia/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params:  types.DefaultParams(),
		PostMap: []types.Post{{Index: "0"}, {Index: "1"}}, ProfileMap: []types.Profile{{Index: "0"}, {Index: "1"}}, CommentMap: []types.Comment{{Index: "0"}, {Index: "1"}}}

	f := initFixture(t)
	err := f.keeper.InitGenesis(f.ctx, genesisState)
	require.NoError(t, err)
	got, err := f.keeper.ExportGenesis(f.ctx)
	require.NoError(t, err)
	require.NotNil(t, got)

	require.EqualExportedValues(t, genesisState.Params, got.Params)
	require.EqualExportedValues(t, genesisState.PostMap, got.PostMap)
	require.EqualExportedValues(t, genesisState.ProfileMap, got.ProfileMap)
	require.EqualExportedValues(t, genesisState.CommentMap, got.CommentMap)

}
