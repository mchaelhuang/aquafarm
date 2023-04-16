package usecase

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/mchaelhuang/aquafarm/internal"
	"github.com/mchaelhuang/aquafarm/internal/constant"
	"github.com/mchaelhuang/aquafarm/internal/entity"
)

type PondUC struct {
	logger   *zap.Logger
	farmRepo internal.FarmRepo
	pondRepo internal.PondRepo
}

func NewPondUC(
	logger *zap.Logger,
	farmRepo internal.FarmRepo,
	pondRepo internal.PondRepo,
) *PondUC {
	return &PondUC{
		logger:   logger,
		farmRepo: farmRepo,
		pondRepo: pondRepo,
	}
}

func (f *PondUC) GetList(ctx context.Context, filter entity.PondFilter) ([]entity.Pond, error) {
	return f.pondRepo.Get(ctx, filter)
}

func (f *PondUC) Get(ctx context.Context, filter entity.PondFilter) (entity.Pond, error) {
	pond, err := f.pondRepo.Get(ctx, filter)
	if err != nil {
		return entity.Pond{}, err
	}
	return pond[0], nil
}

func (f *PondUC) Store(ctx context.Context, pond entity.Pond) (int, error) {
	// Check farm should exist
	farmFilter := entity.FarmFilter{
		Farm: entity.Farm{ID: pond.FarmID},
	}
	_, err := f.farmRepo.Get(ctx, farmFilter)
	if errors.Is(err, constant.ErrNotFound) {
		return 0, constant.ErrIncorrectFarmID
	}
	if err != nil {
		return 0, err
	}

	return f.pondRepo.Store(ctx, pond)
}

func (f *PondUC) Delete(ctx context.Context, id int) error {
	return f.pondRepo.Delete(ctx, id)
}
