package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"swisstronik/x/compliance/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) AddressInfo(goCtx context.Context, req *types.QueryAddressInfoRequest) (*types.QueryAddressInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}

	verificationData, err := k.GetAddressInfo(ctx, address)
	if err != nil {
		return nil, err
	}

	return &types.QueryAddressInfoResponse{Data: verificationData}, nil
}

func (k Keeper) IssuerDetails(goCtx context.Context, req *types.QueryIssuerDetailsRequest) (*types.QueryIssuerDetailsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	issuerAddress, err := sdk.AccAddressFromBech32(req.IssuerAddress)
	if err != nil {
		return nil, err
	}

	issuerDetails, err := k.GetIssuerDetails(ctx, issuerAddress)
	if err != nil {
		return nil, err
	}

	return &types.QueryIssuerDetailsResponse{Details: issuerDetails}, nil
}

func (k Keeper) VerificationDetails(goCtx context.Context, req *types.QueryVerificationDetailsRequest) (*types.QueryVerificationDetailsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	details, err := k.GetVerificationDetails(ctx, []byte(req.VerificationID))
	if err != nil {
		return nil, err
	}

	if details == nil {
		return &types.QueryVerificationDetailsResponse{}, nil
	}

	return &types.QueryVerificationDetailsResponse{Details: details}, nil
}
