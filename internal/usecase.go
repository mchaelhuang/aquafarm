package internal

import (
	"context"

	"github.com/mchaelhuang/aquafarm/internal/entity"
)

type FarmUC interface {
	// GetList returns list of  record in slice
	GetList(ctx context.Context, filter entity.FarmFilter) ([]entity.Farm, error)
	// Get returns single record
	Get(ctx context.Context, filter entity.FarmFilter) (entity.Farm, error)
	// Store create or update record
	Store(ctx context.Context, farm entity.Farm) (int, error)
	// Delete soft-delete the record
	Delete(ctx context.Context, id int) error
}

type PondUC interface {
	// GetList returns list of record in slice
	GetList(ctx context.Context, filter entity.PondFilter) ([]entity.Pond, error)
	// Get returns single record
	Get(ctx context.Context, filter entity.PondFilter) (entity.Pond, error)
	// Store create or update record
	Store(ctx context.Context, farm entity.Pond) (int, error)
	// Delete soft-delete the record
	Delete(ctx context.Context, id int) error
}

type StatsUC interface {
	// Add records info
	Add(ctx context.Context, info entity.StatsRequestInfo) error
	// Get returns all counted stats
	Get(ctx context.Context) (map[string]entity.StatsResult, error)
}
