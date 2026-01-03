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

func createNSocialConnection(keeper keeper.Keeper, ctx context.Context, n int) []types.SocialConnection {
	items := make([]types.SocialConnection, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Follower = strconv.Itoa(i)
		items[i].Following = strconv.Itoa(i)
		items[i].Timestamp = uint64(i)
		_ = keeper.SocialConnection.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestSocialConnectionQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNSocialConnection(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetSocialConnectionRequest
		response *types.QueryGetSocialConnectionResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetSocialConnectionRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetSocialConnectionResponse{SocialConnection: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetSocialConnectionRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetSocialConnectionResponse{SocialConnection: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetSocialConnectionRequest{
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
			response, err := qs.GetSocialConnection(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestSocialConnectionQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNSocialConnection(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSocialConnectionRequest {
		return &types.QueryAllSocialConnectionRequest{
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
			resp, err := qs.ListSocialConnection(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SocialConnection), step)
			require.Subset(t, msgs, resp.SocialConnection)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListSocialConnection(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SocialConnection), step)
			require.Subset(t, msgs, resp.SocialConnection)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListSocialConnection(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.SocialConnection)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListSocialConnection(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
