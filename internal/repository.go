package internal

import (
	"context"

	"github.com/mchaelhuang/aquafarm/internal/entity"
)

// FarmRepo is interface for pond repository.
type FarmRepo interface {
	// Get gets multiple or single row
	Get(ctx context.Context, filter entity.FarmFilter) ([]entity.Farm, error)
	// Store returns inserted or updated row id
	Store(ctx context.Context, farm entity.Farm) (id int, err error)
	// Delete soft-delete
	Delete(ctx context.Context, id int) error
}

// PondRepo is interface for pond repository.
type PondRepo interface {
	// Get gets multiple or single row
	Get(ctx context.Context, filter entity.PondFilter) ([]entity.Pond, error)
	// Store returns inserted or updated row id
	Store(ctx context.Context, farm entity.Pond) (id int, err error)
	// Delete soft-delete
	Delete(ctx context.Context, id int) error
}

type StatsRepo interface {
	// IncrEndpoint initiate or increment endpoint counter
	IncrEndpoint(ctx context.Context, info entity.StatsRequestInfo) error
	// CollectUserAgent stores unique user agent
	CollectUserAgent(ctx context.Context, info entity.StatsRequestInfo) error
	// GetEndpointCount returns endpoint list and count for each endpoint
	GetEndpointCount(ctx context.Context) (map[string]int, error)
	// GetUniqueAgentCount returns unique agent counter
	GetUniqueAgentCount(ctx context.Context, endpointKey string) (int, error)
}
