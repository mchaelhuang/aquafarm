package repository

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/mchaelhuang/aquafarm/internal/config"
	"github.com/mchaelhuang/aquafarm/internal/constant"
	"github.com/mchaelhuang/aquafarm/internal/entity"
)

type PondRepo struct {
	cfg    *config.Cfg
	logger *zap.Logger
	db     *gorm.DB
}

func NewPondRepo(
	cfg *config.Cfg,
	logger *zap.Logger,
	db *gorm.DB,
) *PondRepo {
	db = db.Table(constant.TablePond)
	return &PondRepo{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
}

func (p *PondRepo) Get(ctx context.Context, filter entity.PondFilter) ([]entity.Pond, error) {
	// GetList multiple records
	if filter.ID == 0 {
		var ponds []entity.Pond
		tx := p.filterWhere(p.db.WithContext(ctx), filter)

		err := tx.Find(&ponds).Error
		if err != nil {
			return nil, err
		}
		if len(ponds) == 0 {
			return nil, constant.ErrNotFound
		}

		return ponds, nil
	}

	// GetList single record
	var pond entity.Pond
	err := p.db.WithContext(ctx).Take(&pond, filter.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return []entity.Pond{pond}, nil
}

func (p *PondRepo) filterWhere(tx *gorm.DB, filter entity.PondFilter) *gorm.DB {
	if filter.FarmID != 0 {
		tx = tx.Where("farm_id = ?", filter.FarmID)
	}
	if filter.Label != "" {
		tx = tx.Where("LOWER(label) LIKE LOWER(?)", "%"+filter.Label+"%")
	}
	if filter.Volume != 0 {
		tx = tx.Where("volume = ?", filter.Volume)
	}

	return tx
}

func (p *PondRepo) Store(ctx context.Context, pond entity.Pond) (int, error) {
	tx := p.db.WithContext(ctx)

	if pond.ID == 0 {
		// Create new record
		err := tx.Create(&pond).Error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return 0, constant.ErrDuplicateRecord
		}
		if err != nil {
			return 0, err
		}

		return pond.ID, nil
	}

	// Update existing record
	tx = tx.Model(&pond).Updates(&pond)
	err := tx.Error
	if err != nil {
		return 0, err
	}
	if tx.RowsAffected == 0 {
		return 0, constant.ErrNotFound
	}

	return pond.ID, nil
}

func (p *PondRepo) Delete(ctx context.Context, id int) error {
	tx := p.db.WithContext(ctx).Delete(&entity.Pond{}, id)
	err := tx.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return constant.ErrNotFound
	}
	if tx.RowsAffected == 0 {
		return constant.ErrNotFound
	}

	return err
}
