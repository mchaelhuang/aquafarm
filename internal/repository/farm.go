package repository

import (
	"context"
	"errors"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/mchaelhuang/aquafarm/internal/config"
	"github.com/mchaelhuang/aquafarm/internal/constant"
	"github.com/mchaelhuang/aquafarm/internal/entity"
)

type FarmRepo struct {
	cfg    *config.Cfg
	logger *zap.Logger
	db     *gorm.DB
}

func NewFarmRepo(
	cfg *config.Cfg,
	logger *zap.Logger,
	db *gorm.DB,
) *FarmRepo {
	db = db.Table(constant.TableFarm).Debug()
	return &FarmRepo{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
}

func (f *FarmRepo) Get(ctx context.Context, filter entity.FarmFilter) ([]entity.Farm, error) {
	// GetList multiple records
	if filter.ID == 0 {
		var farms []entity.Farm
		tx := f.filterWhere(f.db.WithContext(ctx), filter)

		err := tx.Find(&farms).Error
		if err != nil {
			return nil, err
		}
		if len(farms) == 0 {
			return nil, constant.ErrNotFound
		}

		return farms, nil
	}

	// GetList single record
	var farm entity.Farm
	err := f.db.WithContext(ctx).Take(&farm, filter.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return []entity.Farm{farm}, nil
}

func (f *FarmRepo) filterWhere(tx *gorm.DB, filter entity.FarmFilter) *gorm.DB {
	if filter.Name != "" {
		tx = tx.Where("LOWER(name) LIKE LOWER(?)", "%"+filter.Name+"%")
	}
	if filter.Location != "" {
		tx = tx.Where("LOWER(location) LIKE LOWER(?)", "%"+filter.Location+"%")
	}

	return tx
}

func (f *FarmRepo) Store(ctx context.Context, farm entity.Farm) (int, error) {
	tx := f.db.WithContext(ctx)

	if farm.ID == 0 {
		// Create new record
		err := tx.Create(&farm).Error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return 0, constant.ErrDuplicateRecord
		}
		if err != nil {
			return 0, err
		}

		return farm.ID, nil
	}

	// Update existing record
	tx = tx.Model(&farm).Updates(&farm)
	err := tx.Error
	if err != nil {
		return 0, err
	}
	if tx.RowsAffected == 0 {
		return 0, constant.ErrNotFound
	}

	return farm.ID, nil
}

func (f *FarmRepo) Delete(ctx context.Context, id int) error {
	tx := f.db.WithContext(ctx).Delete(&entity.Farm{}, id)
	err := tx.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return constant.ErrNotFound
	}
	if tx.RowsAffected == 0 {
		return constant.ErrNotFound
	}

	return err
}
