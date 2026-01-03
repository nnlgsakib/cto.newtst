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

func (q queryServer) ListComment(ctx context.Context, req *types.QueryAllCommentRequest) (*types.QueryAllCommentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	comments, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Comment,
		req.Pagination,
		func(_ string, value types.Comment) (types.Comment, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCommentResponse{Comment: comments, Pagination: pageRes}, nil
}

func (q queryServer) GetComment(ctx context.Context, req *types.QueryGetCommentRequest) (*types.QueryGetCommentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.Comment.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetCommentResponse{Comment: val}, nil
}
