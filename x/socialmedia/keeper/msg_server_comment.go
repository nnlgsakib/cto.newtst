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

func (k msgServer) CreateComment(ctx context.Context, msg *types.MsgCreateComment) (*types.MsgCreateCommentResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.Comment.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var comment = types.Comment{
		Creator:   msg.Creator,
		Index:     msg.Index,
		PostId:    msg.PostId,
		Content:   msg.Content,
		Timestamp: msg.Timestamp,
		Author:    msg.Author,
	}

	if err := k.Comment.Set(ctx, comment.Index, comment); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateCommentResponse{}, nil
}

func (k msgServer) UpdateComment(ctx context.Context, msg *types.MsgUpdateComment) (*types.MsgUpdateCommentResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Comment.Get(ctx, msg.Index)
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

	var comment = types.Comment{
		Creator:   msg.Creator,
		Index:     msg.Index,
		PostId:    msg.PostId,
		Content:   msg.Content,
		Timestamp: msg.Timestamp,
		Author:    msg.Author,
	}

	if err := k.Comment.Set(ctx, comment.Index, comment); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update comment")
	}

	return &types.MsgUpdateCommentResponse{}, nil
}

func (k msgServer) DeleteComment(ctx context.Context, msg *types.MsgDeleteComment) (*types.MsgDeleteCommentResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Comment.Get(ctx, msg.Index)
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

	if err := k.Comment.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove comment")
	}

	return &types.MsgDeleteCommentResponse{}, nil
}
