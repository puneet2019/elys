package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m Migrator) V9Migration(ctx sdk.Context) error {
	allPools := m.keeper.GetAllAmmPools(ctx)
	for _, pool := range allPools {
		pool.TotalLiabilities = sdk.Coins{}
		m.keeper.SetAmmPool(ctx, pool)
	}
	return nil
}
