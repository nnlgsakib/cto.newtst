package keeper

import (
	"context"
	"errors"

	"nlg/x/socialmedia/types"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListSocialConnection(ctx context.Context, req *types.QueryAllSocialConnectionRequest) (*types.QueryAllSocialConnectionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	socialConnections, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.SocialConnection,
		req.Pagination,
		func(_ string, value types.SocialConnection) (types.SocialConnection, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSocialConnectionResponse{SocialConnection: socialConnections, Pagination: pageRes}, nil
}

func (q queryServer) GetSocialConnection(ctx context.Context, req *types.QueryGetSocialConnectionRequest) (*types.QueryGetSocialConnectionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.SocialConnection.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetSocialConnectionResponse{SocialConnection: val}, nil
}
