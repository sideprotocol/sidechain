package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/sideprotocol/side/x/gmm/keeper"
	"github.com/sideprotocol/side/x/gmm/types"
)

func SimulateMsgSwap(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSwap{
			Sender: simAccount.Address.String(),
		}

		// TODO: Handling the Swap simulation
		_ = k
		_ = bk
		_ = ak
		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Swap simulation not implemented"), nil, nil
	}
}
