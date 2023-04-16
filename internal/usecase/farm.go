package usecase

import (
	"context"

	"github.com/mchaelhuang/aquafarm/internal"
	"github.com/mchaelhuang/aquafarm/internal/entity"
)

type FarmUC struct {
	farmRepo internal.FarmRepo
	pondRepo internal.PondRepo
}

func NewFarmUC(
	farmRepo internal.FarmRepo,
	pondRepo internal.PondRepo,
) *FarmUC {
	return &FarmUC{
		farmRepo: farmRepo,
		pondRepo: pondRepo,
	}
}

func (f *FarmUC) GetList(ctx context.Context, filter entity.FarmFilter) ([]entity.Farm, error) {
	return f.farmRepo.Get(ctx, filter)
}

func (f *FarmUC) Get(ctx context.Context, filter entity.FarmFilter) (entity.Farm, error) {
	farms, err := f.farmRepo.Get(ctx, filter)
	if err != nil {
		return entity.Farm{}, err
	}
	farm := farms[0]

	pondFilter := entity.PondFilter{
		Pond: entity.Pond{
			FarmID: farm.ID,
		},
	}

	// Get relations
	farm.Ponds, err = f.pondRepo.Get(ctx, pondFilter)
	if err != nil {
		return entity.Farm{}, err
	}

	return farm, nil
}

func (f *FarmUC) Store(ctx context.Context, farm entity.Farm) (int, error) {
	return f.farmRepo.Store(ctx, farm)
}

func (f *FarmUC) Delete(ctx context.Context, id int) error {
	return f.farmRepo.Delete(ctx, id)
}
