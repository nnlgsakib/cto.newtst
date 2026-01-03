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

func (k msgServer) CreateProfile(ctx context.Context, msg *types.MsgCreateProfile) (*types.MsgCreateProfileResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.Profile.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var profile = types.Profile{
		Creator:        msg.Creator,
		Index:          msg.Index,
		Username:       msg.Username,
		Bio:            msg.Bio,
		AvatarIpfsHash: msg.AvatarIpfsHash,
		Reputation:     msg.Reputation,
		FollowersCount: msg.FollowersCount,
		FollowingCount: msg.FollowingCount,
	}

	if err := k.Profile.Set(ctx, profile.Index, profile); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateProfileResponse{}, nil
}

func (k msgServer) UpdateProfile(ctx context.Context, msg *types.MsgUpdateProfile) (*types.MsgUpdateProfileResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Profile.Get(ctx, msg.Index)
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

	var profile = types.Profile{
		Creator:        msg.Creator,
		Index:          msg.Index,
		Username:       msg.Username,
		Bio:            msg.Bio,
		AvatarIpfsHash: msg.AvatarIpfsHash,
		Reputation:     msg.Reputation,
		FollowersCount: msg.FollowersCount,
		FollowingCount: msg.FollowingCount,
	}

	if err := k.Profile.Set(ctx, profile.Index, profile); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update profile")
	}

	return &types.MsgUpdateProfileResponse{}, nil
}

func (k msgServer) DeleteProfile(ctx context.Context, msg *types.MsgDeleteProfile) (*types.MsgDeleteProfileResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Profile.Get(ctx, msg.Index)
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

	if err := k.Profile.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove profile")
	}

	return &types.MsgDeleteProfileResponse{}, nil
}
