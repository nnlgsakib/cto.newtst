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

func (k msgServer) CreatePost(ctx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.Post.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var post = types.Post{
		Creator:    msg.Creator,
		Index:      msg.Index,
		Title:      msg.Title,
		Content:    msg.Content,
		IpfsHash:   msg.IpfsHash,
		Timestamp:  msg.Timestamp,
		LikesCount: msg.LikesCount,
	}

	if err := k.Post.Set(ctx, post.Index, post); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreatePostResponse{}, nil
}

func (k msgServer) UpdatePost(ctx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Post.Get(ctx, msg.Index)
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

	var post = types.Post{
		Creator:    msg.Creator,
		Index:      msg.Index,
		Title:      msg.Title,
		Content:    msg.Content,
		IpfsHash:   msg.IpfsHash,
		Timestamp:  msg.Timestamp,
		LikesCount: msg.LikesCount,
	}

	if err := k.Post.Set(ctx, post.Index, post); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update post")
	}

	return &types.MsgUpdatePostResponse{}, nil
}

func (k msgServer) DeletePost(ctx context.Context, msg *types.MsgDeletePost) (*types.MsgDeletePostResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Post.Get(ctx, msg.Index)
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

	if err := k.Post.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove post")
	}

	return &types.MsgDeletePostResponse{}, nil
}
