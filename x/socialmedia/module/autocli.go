package socialmedia

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"nlg/x/socialmedia/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ListPost",
					Use:       "list-post",
					Short:     "List all post",
				},
				{
					RpcMethod:      "GetPost",
					Use:            "get-post [id]",
					Short:          "Gets a post",
					Alias:          []string{"show-post"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ListProfile",
					Use:       "list-profile",
					Short:     "List all profile",
				},
				{
					RpcMethod:      "GetProfile",
					Use:            "get-profile [id]",
					Short:          "Gets a profile",
					Alias:          []string{"show-profile"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ListComment",
					Use:       "list-comment",
					Short:     "List all comment",
				},
				{
					RpcMethod:      "GetComment",
					Use:            "get-comment [id]",
					Short:          "Gets a comment",
					Alias:          []string{"show-comment"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ListSocialConnection",
					Use:       "list-social-connection",
					Short:     "List all socialConnection",
				},
				{
					RpcMethod:      "GetSocialConnection",
					Use:            "get-social-connection [id]",
					Short:          "Gets a socialConnection",
					Alias:          []string{"show-social-connection"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreatePost",
					Use:            "create-post [index] [title] [content] [ipfs-hash] [timestamp] [likes-count]",
					Short:          "Create a new post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "title"}, {ProtoField: "content"}, {ProtoField: "ipfs_hash"}, {ProtoField: "timestamp"}, {ProtoField: "likes_count"}},
				},
				{
					RpcMethod:      "UpdatePost",
					Use:            "update-post [index] [title] [content] [ipfs-hash] [timestamp] [likes-count]",
					Short:          "Update post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "title"}, {ProtoField: "content"}, {ProtoField: "ipfs_hash"}, {ProtoField: "timestamp"}, {ProtoField: "likes_count"}},
				},
				{
					RpcMethod:      "DeletePost",
					Use:            "delete-post [index]",
					Short:          "Delete post",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateProfile",
					Use:            "create-profile [index] [username] [bio] [avatar-ipfs-hash] [reputation] [followers-count] [following-count]",
					Short:          "Create a new profile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "username"}, {ProtoField: "bio"}, {ProtoField: "avatar_ipfs_hash"}, {ProtoField: "reputation"}, {ProtoField: "followers_count"}, {ProtoField: "following_count"}},
				},
				{
					RpcMethod:      "UpdateProfile",
					Use:            "update-profile [index] [username] [bio] [avatar-ipfs-hash] [reputation] [followers-count] [following-count]",
					Short:          "Update profile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "username"}, {ProtoField: "bio"}, {ProtoField: "avatar_ipfs_hash"}, {ProtoField: "reputation"}, {ProtoField: "followers_count"}, {ProtoField: "following_count"}},
				},
				{
					RpcMethod:      "DeleteProfile",
					Use:            "delete-profile [index]",
					Short:          "Delete profile",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateComment",
					Use:            "create-comment [index] [post-id] [content] [timestamp] [author]",
					Short:          "Create a new comment",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "post_id"}, {ProtoField: "content"}, {ProtoField: "timestamp"}, {ProtoField: "author"}},
				},
				{
					RpcMethod:      "UpdateComment",
					Use:            "update-comment [index] [post-id] [content] [timestamp] [author]",
					Short:          "Update comment",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "post_id"}, {ProtoField: "content"}, {ProtoField: "timestamp"}, {ProtoField: "author"}},
				},
				{
					RpcMethod:      "DeleteComment",
					Use:            "delete-comment [index]",
					Short:          "Delete comment",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateSocialConnection",
					Use:            "create-social-connection [index] [follower] [following] [timestamp]",
					Short:          "Create a new socialConnection",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "follower"}, {ProtoField: "following"}, {ProtoField: "timestamp"}},
				},
				{
					RpcMethod:      "UpdateSocialConnection",
					Use:            "update-social-connection [index] [follower] [following] [timestamp]",
					Short:          "Update socialConnection",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "follower"}, {ProtoField: "following"}, {ProtoField: "timestamp"}},
				},
				{
					RpcMethod:      "DeleteSocialConnection",
					Use:            "delete-social-connection [index]",
					Short:          "Delete socialConnection",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "LikePost",
					Use:            "like-post [post-id]",
					Short:          "Send a likePost tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "post_id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
