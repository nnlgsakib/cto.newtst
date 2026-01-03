package socialmedia

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"nlg/testutil/sample"
	socialmediasimulation "nlg/x/socialmedia/simulation"
	"nlg/x/socialmedia/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	socialmediaGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PostMap: []types.Post{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}, ProfileMap: []types.Profile{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}, CommentMap: []types.Comment{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&socialmediaGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreatePost          = "op_weight_msg_socialmedia"
		defaultWeightMsgCreatePost int = 100
	)

	var weightMsgCreatePost int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePost, &weightMsgCreatePost, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePost = defaultWeightMsgCreatePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePost,
		socialmediasimulation.SimulateMsgCreatePost(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdatePost          = "op_weight_msg_socialmedia"
		defaultWeightMsgUpdatePost int = 100
	)

	var weightMsgUpdatePost int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePost, &weightMsgUpdatePost, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePost = defaultWeightMsgUpdatePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePost,
		socialmediasimulation.SimulateMsgUpdatePost(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeletePost          = "op_weight_msg_socialmedia"
		defaultWeightMsgDeletePost int = 100
	)

	var weightMsgDeletePost int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePost, &weightMsgDeletePost, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePost = defaultWeightMsgDeletePost
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePost,
		socialmediasimulation.SimulateMsgDeletePost(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreateProfile          = "op_weight_msg_socialmedia"
		defaultWeightMsgCreateProfile int = 100
	)

	var weightMsgCreateProfile int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateProfile, &weightMsgCreateProfile, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProfile = defaultWeightMsgCreateProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProfile,
		socialmediasimulation.SimulateMsgCreateProfile(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateProfile          = "op_weight_msg_socialmedia"
		defaultWeightMsgUpdateProfile int = 100
	)

	var weightMsgUpdateProfile int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateProfile, &weightMsgUpdateProfile, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProfile = defaultWeightMsgUpdateProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateProfile,
		socialmediasimulation.SimulateMsgUpdateProfile(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteProfile          = "op_weight_msg_socialmedia"
		defaultWeightMsgDeleteProfile int = 100
	)

	var weightMsgDeleteProfile int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteProfile, &weightMsgDeleteProfile, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProfile = defaultWeightMsgDeleteProfile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteProfile,
		socialmediasimulation.SimulateMsgDeleteProfile(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreateComment          = "op_weight_msg_socialmedia"
		defaultWeightMsgCreateComment int = 100
	)

	var weightMsgCreateComment int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateComment, &weightMsgCreateComment, nil,
		func(_ *rand.Rand) {
			weightMsgCreateComment = defaultWeightMsgCreateComment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateComment,
		socialmediasimulation.SimulateMsgCreateComment(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateComment          = "op_weight_msg_socialmedia"
		defaultWeightMsgUpdateComment int = 100
	)

	var weightMsgUpdateComment int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateComment, &weightMsgUpdateComment, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateComment = defaultWeightMsgUpdateComment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateComment,
		socialmediasimulation.SimulateMsgUpdateComment(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteComment          = "op_weight_msg_socialmedia"
		defaultWeightMsgDeleteComment int = 100
	)

	var weightMsgDeleteComment int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteComment, &weightMsgDeleteComment, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteComment = defaultWeightMsgDeleteComment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteComment,
		socialmediasimulation.SimulateMsgDeleteComment(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}
