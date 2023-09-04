package app

// import (
// 	"encoding/json"
// 	"time"

// 	sdkmath "cosmossdk.io/math"
// 	dbm "github.com/cometbft/cometbft-db"
// 	abci "github.com/cometbft/cometbft/abci/types"
// 	"github.com/cometbft/cometbft/crypto/secp256k1"
// 	"github.com/cometbft/cometbft/libs/log"
// 	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
// 	tmtypes "github.com/cometbft/cometbft/types"
// 	"github.com/cosmos/cosmos-sdk/client/flags"
// 	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
// 	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
// 	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
// 	"github.com/cosmos/cosmos-sdk/server"
// 	"github.com/cosmos/cosmos-sdk/testutil/mock"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
// 	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
// 	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
// 	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

// 	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
// 	simcli "github.com/cosmos/cosmos-sdk/x/simulation/client/cli"
// )

// const (
// 	Bech32Prefix = "side"
// )

// var (
// 	Denoms = []string{"aside", "bside"}
// )

// // Initiate a new ElysApp object - Common function used by the following 2 functions.
// func InitiateNewSideApp() *App {
// 	db := dbm.NewMemDB()
// 	appOptions := make(simtestutil.AppOptionsMap, 0)
// 	appOptions[flags.FlagHome] = DefaultNodeHome
// 	appOptions[server.FlagInvCheckPeriod] = simcli.FlagPeriodValue
// 	nodeHome := ""
// 	appOptions[flags.FlagHome] = nodeHome // ensure unique folder
// 	appOptions[server.FlagInvCheckPeriod] = 1
// 	encodingConfig := MakeEncodingConfig()
// 	app := New(
// 		log.NewNopLogger(),
// 		db,
// 		nil,
// 		true,
// 		nil,
// 		nodeHome,
// 		0,
// 		encodingConfig,
// 		appOptions,
// 	)

// 	return app
// }

// // Initializes a new ElysApp without IBC functionality
// func InitSideTestApp(initChain bool) *App {
// 	app := InitiateNewSideApp()
// 	if initChain {
// 		genesisState, valSet, _, _ := GenesisStateWithValSet(app)
// 		stateBytes, err := json.MarshalIndent(genesisState, "", " ")
// 		if err != nil {
// 			panic(err)
// 		}

// 		app.InitChain(
// 			abci.RequestInitChain{
// 				Validators:      []abci.ValidatorUpdate{},
// 				ConsensusParams: simtestutil.DefaultConsensusParams,
// 				AppStateBytes:   stateBytes,
// 			},
// 		)

// 		// commit genesis changes
// 		app.Commit()
// 		app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{
// 			Height:             app.LastBlockHeight() + 1,
// 			AppHash:            app.LastCommitID().Hash,
// 			ValidatorsHash:     valSet.Hash(),
// 			NextValidatorsHash: valSet.Hash(),
// 		}})
// 	}

// 	return app
// }

// // Initializes a new ElysApp without IBC functionality and returns genesis account (delegator)
// func InitElysTestAppWithGenAccount() (*App, sdk.AccAddress, sdk.ValAddress) {
// 	app := InitiateNewSideApp()

// 	genesisState, valSet, genAcount, valAddress := GenesisStateWithValSet(app)
// 	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
// 	if err != nil {
// 		panic(err)
// 	}

// 	app.InitChain(
// 		abci.RequestInitChain{
// 			Validators:      []abci.ValidatorUpdate{},
// 			ConsensusParams: simtestutil.DefaultConsensusParams,
// 			AppStateBytes:   stateBytes,
// 		},
// 	)

// 	// commit genesis changes
// 	app.Commit()
// 	app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{
// 		Height:             app.LastBlockHeight() + 1,
// 		AppHash:            app.LastCommitID().Hash,
// 		ValidatorsHash:     valSet.Hash(),
// 		NextValidatorsHash: valSet.Hash(),
// 	}})

// 	return app, genAcount, valAddress
// }

// func GenesisStateWithValSet(app *App) (GenesisState, *tmtypes.ValidatorSet, sdk.AccAddress, sdk.ValAddress) {
// 	privVal := mock.NewPV()
// 	pubKey, _ := privVal.GetPubKey()
// 	validator := tmtypes.NewValidator(pubKey, 1)
// 	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})

// 	// generate genesis account
// 	senderPrivKey := secp256k1.GenPrivKey()
// 	senderPrivKey.PubKey().Address()
// 	acc := authtypes.NewBaseAccountWithAddress(senderPrivKey.PubKey().Address().Bytes())
// 	balance := banktypes.Balance{
// 		Address: acc.GetAddress().String(),
// 		Coins: sdk.NewCoins(
// 			sdk.NewCoin(Denoms[0], sdkmath.NewInt(100000000000000)),
// 			sdk.NewCoin(Denoms[1], sdkmath.NewInt(100000000000000)),
// 		),
// 	}

// 	//////////////////////
// 	balances := []banktypes.Balance{balance}
// 	genesisState := NewDefaultGenesisState(app.AppCodec())
// 	genAccs := []authtypes.GenesisAccount{acc}
// 	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
// 	genesisState[authtypes.ModuleName] = app.AppCodec().MustMarshalJSON(authGenesis)

// 	validators := make([]stakingtypes.Validator, 0, len(valSet.Validators))
// 	delegations := make([]stakingtypes.Delegation, 0, len(valSet.Validators))

// 	bondAmt := sdk.DefaultPowerReduction

// 	for _, val := range valSet.Validators {
// 		pk, _ := cryptocodec.FromTmPubKeyInterface(val.PubKey)
// 		pkAny, _ := codectypes.NewAnyWithValue(pk)
// 		validator := stakingtypes.Validator{
// 			OperatorAddress:   sdk.ValAddress(val.Address).String(),
// 			ConsensusPubkey:   pkAny,
// 			Jailed:            false,
// 			Status:            stakingtypes.Bonded,
// 			Tokens:            bondAmt,
// 			DelegatorShares:   sdk.OneDec(),
// 			Description:       stakingtypes.Description{},
// 			UnbondingHeight:   int64(0),
// 			UnbondingTime:     time.Unix(0, 0).UTC(),
// 			Commission:        stakingtypes.NewCommission(sdk.NewDecWithPrec(5, 2), sdk.NewDecWithPrec(10, 2), sdk.NewDecWithPrec(10, 2)),
// 			MinSelfDelegation: sdkmath.OneInt(),
// 		}
// 		validators = append(validators, validator)
// 		delegations = append(delegations, stakingtypes.NewDelegation(genAccs[0].GetAddress(), val.Address.Bytes(), sdk.OneDec()))

// 	}
// 	// set validators and delegations
// 	params := stakingtypes.DefaultParams()
// 	params.BondDenom = Denoms[0]

// 	stakingGenesis := stakingtypes.NewGenesisState(params, validators, delegations)
// 	genesisState[stakingtypes.ModuleName] = app.AppCodec().MustMarshalJSON(stakingGenesis)

// 	totalSupply := sdk.NewCoins()
// 	for _, b := range balances {
// 		// add genesis acc tokens to total supply
// 		totalSupply = totalSupply.Add(b.Coins...)
// 	}

// 	for range delegations {
// 		// add delegated tokens to total supply
// 		totalSupply = totalSupply.Add(sdk.NewCoin(Denoms[0], bondAmt))
// 	}

// 	// add bonded amount to bonded pool module account
// 	balances = append(balances, banktypes.Balance{
// 		Address: authtypes.NewModuleAddress(stakingtypes.BondedPoolName).String(),
// 		Coins:   sdk.Coins{sdk.NewCoin(Denoms[0], bondAmt)},
// 	})

// 	// update total supply
// 	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{}, []banktypes.SendEnabled{})
// 	genesisState[banktypes.ModuleName] = app.AppCodec().MustMarshalJSON(bankGenesis)
// 	return genesisState, valSet, genAccs[0].GetAddress(), sdk.ValAddress(validator.Address)
// }

// type GenerateAccountStrategy func(int) []sdk.AccAddress

// // createRandomAccounts is a strategy used by addTestAddrs() in order to generated addresses in random order.
// func createRandomAccounts(accNum int) []sdk.AccAddress {
// 	testAddrs := make([]sdk.AccAddress, accNum)
// 	for i := 0; i < accNum; i++ {
// 		pk := ed25519.GenPrivKey().PubKey()
// 		testAddrs[i] = sdk.AccAddress(pk.Address())
// 	}

// 	return testAddrs
// }

// // AddTestAddrs constructs and returns accNum amount of accounts with an
// // initial balance of accAmt in random order
// func AddTestAddrs(app *App, ctx sdk.Context, accNum int, accAmt sdkmath.Int) []sdk.AccAddress {
// 	return addTestAddrs(app, ctx, accNum, accAmt, createRandomAccounts)
// }

// func addTestAddrs(app *App, ctx sdk.Context, accNum int, accAmt sdkmath.Int, strategy GenerateAccountStrategy) []sdk.AccAddress {
// 	testAddrs := strategy(accNum)

// 	bondDenom := app.StakingKeeper.BondDenom(ctx)
// 	initCoins := sdk.NewCoins(sdk.NewCoin(bondDenom, accAmt))

// 	for _, addr := range testAddrs {
// 		initAccountWithCoins(app, ctx, addr, initCoins)
// 	}

// 	return testAddrs
// }

// func initAccountWithCoins(app *App, ctx sdk.Context, addr sdk.AccAddress, coins sdk.Coins) {
// 	err := app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, coins)
// 	if err != nil {
// 		panic(err)
// 	}
// }
