package keeper

import (
	"context"

	"github.com/commercionetwork/commercionetwork/x/commerciomint/types"
	ctypes "github.com/commercionetwork/commercionetwork/x/common/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MintCCC(goCtx context.Context, msg *types.MsgMintCCC) (*types.MsgMintCCCResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var requestCoins sdk.Coins
	for _, coin := range msg.DepositAmount {
		requestCoins = append(requestCoins, *coin)
	}

	err := k.NewPosition(
		ctx,
		msg.Depositor,
		requestCoins,
		msg.ID,
	)
	if err != nil {
		return nil, errors.Wrap(errors.ErrInvalidRequest, err.Error())
	}
	ctypes.EmitCommonEvents(ctx, msg.Depositor)
	return &types.MsgMintCCCResponse{
		ID: msg.ID,
	}, nil
}

func (k msgServer) BurnCCC(goCtx context.Context, msg *types.MsgBurnCCC) (*types.MsgBurnCCCResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, err
	}
	residualAmount, err := k.RemoveCCC(
		ctx,
		signer,
		msg.ID,
		*msg.Amount,
	)
	if err != nil {
		return nil, err
	}
	ctypes.EmitCommonEvents(ctx, msg.Signer)
	residualCredits := sdk.NewCoin(types.CreditsDenom, residualAmount)

	return &types.MsgBurnCCCResponse{
		ID:       msg.ID,
		Residual: &residualCredits,
	}, nil
}
