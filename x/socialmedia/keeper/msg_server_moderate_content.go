package keeper

import (
	"context"

	"nlg/x/socialmedia/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) ModerateContent(ctx context.Context, msg *types.MsgModerateContent) (*types.MsgModerateContentResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgModerateContentResponse{}, nil
}
