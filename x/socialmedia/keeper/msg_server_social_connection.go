package keeper

import (
	"context"
	"errors"
	"fmt"

	"nlg/x/socialmedia/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSocialConnection(ctx context.Context, msg *types.MsgCreateSocialConnection) (*types.MsgCreateSocialConnectionResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.SocialConnection.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var socialConnection = types.SocialConnection{
		Creator:   msg.Creator,
		Index:     msg.Index,
		Follower:  msg.Follower,
		Following: msg.Following,
		Timestamp: msg.Timestamp,
	}

	if err := k.SocialConnection.Set(ctx, socialConnection.Index, socialConnection); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateSocialConnectionResponse{}, nil
}

func (k msgServer) UpdateSocialConnection(ctx context.Context, msg *types.MsgUpdateSocialConnection) (*types.MsgUpdateSocialConnectionResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.SocialConnection.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var socialConnection = types.SocialConnection{
		Creator:   msg.Creator,
		Index:     msg.Index,
		Follower:  msg.Follower,
		Following: msg.Following,
		Timestamp: msg.Timestamp,
	}

	if err := k.SocialConnection.Set(ctx, socialConnection.Index, socialConnection); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update socialConnection")
	}

	return &types.MsgUpdateSocialConnectionResponse{}, nil
}

func (k msgServer) DeleteSocialConnection(ctx context.Context, msg *types.MsgDeleteSocialConnection) (*types.MsgDeleteSocialConnectionResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.SocialConnection.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.SocialConnection.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove socialConnection")
	}

	return &types.MsgDeleteSocialConnectionResponse{}, nil
}
