package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"nlg/x/socialmedia/keeper"
	"nlg/x/socialmedia/types"
)

func createNProfile(keeper keeper.Keeper, ctx context.Context, n int) []types.Profile {
	items := make([]types.Profile, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Username = strconv.Itoa(i)
		items[i].Bio = strconv.Itoa(i)
		items[i].AvatarIpfsHash = strconv.Itoa(i)
		items[i].Reputation = uint64(i)
		items[i].FollowersCount = uint64(i)
		items[i].FollowingCount = uint64(i)
		_ = keeper.Profile.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestProfileQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNProfile(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetProfileRequest
		response *types.QueryGetProfileResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetProfileRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetProfileResponse{Profile: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetProfileRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetProfileResponse{Profile: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetProfileRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetProfile(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestProfileQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNProfile(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllProfileRequest {
		return &types.QueryAllProfileRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListProfile(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Profile), step)
			require.Subset(t, msgs, resp.Profile)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListProfile(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Profile), step)
			require.Subset(t, msgs, resp.Profile)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListProfile(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Profile)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListProfile(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
