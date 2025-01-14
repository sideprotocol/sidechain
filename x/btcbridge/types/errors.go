package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/btcbridge module sentinel errors
var (
	ErrSenderAddressNotAuthorized = errorsmod.Register(ModuleName, 1000, "sender address not authorized")
	ErrInvalidHeader              = errorsmod.Register(ModuleName, 1100, "invalid block header")
	ErrReorgFailed                = errorsmod.Register(ModuleName, 1101, "failed to reorg chain")
	ErrForkedBlockHeader          = errorsmod.Register(ModuleName, 1102, "Invalid forked block header")

	ErrInvalidSenders = errorsmod.Register(ModuleName, 2100, "invalid allowed senders")

	ErrInvalidBtcTransaction     = errorsmod.Register(ModuleName, 3100, "invalid bitcoin transaction")
	ErrBlockNotFound             = errorsmod.Register(ModuleName, 3101, "block not found")
	ErrTransactionNotIncluded    = errorsmod.Register(ModuleName, 3102, "transaction not included in block")
	ErrNotConfirmed              = errorsmod.Register(ModuleName, 3200, "transaction not confirmed")
	ErrExceedMaxAcceptanceDepth  = errorsmod.Register(ModuleName, 3201, "exceed max acceptance block depth")
	ErrUnsupportedScriptType     = errorsmod.Register(ModuleName, 3202, "unsupported script type")
	ErrTransactionAlreadyMinted  = errorsmod.Register(ModuleName, 3203, "transaction already minted")
	ErrInvalidDepositTransaction = errorsmod.Register(ModuleName, 3204, "invalid deposit transaction")

	ErrInvalidSignatures      = errorsmod.Register(ModuleName, 4200, "invalid signatures")
	ErrInsufficientBalance    = errorsmod.Register(ModuleName, 4201, "insufficient balance")
	ErrSigningRequestNotExist = errorsmod.Register(ModuleName, 4202, "signing request does not exist")
	ErrInvalidStatus          = errorsmod.Register(ModuleName, 4203, "invalid status")

	ErrUTXODoesNotExist = errorsmod.Register(ModuleName, 5100, "utxo does not exist")
	ErrUTXOLocked       = errorsmod.Register(ModuleName, 5101, "utxo locked")
	ErrUTXOUnlocked     = errorsmod.Register(ModuleName, 5102, "utxo unlocked")

	ErrInvalidAmount       = errorsmod.Register(ModuleName, 6100, "invalid amount")
	ErrInvalidFeeRate      = errorsmod.Register(ModuleName, 6101, "invalid fee rate")
	ErrAssetNotSupported   = errorsmod.Register(ModuleName, 6102, "asset not supported")
	ErrDustOutput          = errorsmod.Register(ModuleName, 6103, "too small output amount")
	ErrInsufficientUTXOs   = errorsmod.Register(ModuleName, 6104, "insufficient utxos")
	ErrFailToSerializePsbt = errorsmod.Register(ModuleName, 6105, "failed to serialize psbt")

	ErrInvalidRunes  = errorsmod.Register(ModuleName, 7100, "invalid runes")
	ErrInvalidRuneId = errorsmod.Register(ModuleName, 7101, "invalid rune id")

	ErrInvalidParams = errorsmod.Register(ModuleName, 8100, "invalid module params")
)
