package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/elys-network/elys/x/masterchef/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TotalRewards(goCtx context.Context, req *types.QueryTotalRewardsRequest) (*types.QueryTotalRewardsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	v := k.GetTotalRewards(ctx)
	return &types.QueryTotalRewardsResponse{
		TotalRewards: v,
	}, nil
}
