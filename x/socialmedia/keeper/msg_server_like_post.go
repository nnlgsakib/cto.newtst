package keeper

import (
	"context"

	"nlg/x/socialmedia/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) LikePost(ctx context.Context, msg *types.MsgLikePost) (*types.MsgLikePostResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgLikePostResponse{}, nil
}
