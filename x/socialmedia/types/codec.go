package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registrar codectypes.InterfaceRegistry) {
	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgModerateContent{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLikePost{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSocialConnection{},
		&MsgUpdateSocialConnection{},
		&MsgDeleteSocialConnection{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateComment{},
		&MsgUpdateComment{},
		&MsgDeleteComment{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProfile{},
		&MsgUpdateProfile{},
		&MsgDeleteProfile{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePost{},
		&MsgUpdatePost{},
		&MsgDeletePost{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registrar, &_Msg_serviceDesc)
}
