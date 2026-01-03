package types_test

import (
	"testing"

	"nlg/x/socialmedia/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc:     "valid genesis state",
			genState: &types.GenesisState{PostMap: []types.Post{{Index: "0"}, {Index: "1"}}, ProfileMap: []types.Profile{{Index: "0"}, {Index: "1"}}, CommentMap: []types.Comment{{Index: "0"}, {Index: "1"}}, SocialConnectionMap: []types.SocialConnection{{Index: "0"}, {Index: "1"}}},
			valid:    true,
		}, {
			desc: "duplicated post",
			genState: &types.GenesisState{
				PostMap: []types.Post{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
				ProfileMap: []types.Profile{{Index: "0"}, {Index: "1"}}, CommentMap: []types.Comment{{Index: "0"}, {Index: "1"}}, SocialConnectionMap: []types.SocialConnection{{Index: "0"}, {Index: "1"}}},
			valid: false,
		}, {
			desc: "duplicated profile",
			genState: &types.GenesisState{
				ProfileMap: []types.Profile{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
				CommentMap: []types.Comment{{Index: "0"}, {Index: "1"}}, SocialConnectionMap: []types.SocialConnection{{Index: "0"}, {Index: "1"}}},
			valid: false,
		}, {
			desc: "duplicated comment",
			genState: &types.GenesisState{
				CommentMap: []types.Comment{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
				SocialConnectionMap: []types.SocialConnection{{Index: "0"}, {Index: "1"}}},
			valid: false,
		}, {
			desc: "duplicated socialConnection",
			genState: &types.GenesisState{
				SocialConnectionMap: []types.SocialConnection{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
