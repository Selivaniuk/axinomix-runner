package axinomixrunner

import (
	"math/rand"

	"axinomix-runner/testutil/sample"
	axinomixrunnersimulation "axinomix-runner/x/axinomixrunner/simulation"
	"axinomix-runner/x/axinomixrunner/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = axinomixrunnersimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgStartRace = "op_weight_msg_start_race"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStartRace int = 100

	opWeightMsgEndRace = "op_weight_msg_end_race"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEndRace int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	axinomixrunnerGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&axinomixrunnerGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgStartRace int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgStartRace, &weightMsgStartRace, nil,
		func(_ *rand.Rand) {
			weightMsgStartRace = defaultWeightMsgStartRace
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStartRace,
		axinomixrunnersimulation.SimulateMsgStartRace(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEndRace int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEndRace, &weightMsgEndRace, nil,
		func(_ *rand.Rand) {
			weightMsgEndRace = defaultWeightMsgEndRace
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEndRace,
		axinomixrunnersimulation.SimulateMsgEndRace(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgStartRace,
			defaultWeightMsgStartRace,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				axinomixrunnersimulation.SimulateMsgStartRace(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgEndRace,
			defaultWeightMsgEndRace,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				axinomixrunnersimulation.SimulateMsgEndRace(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
