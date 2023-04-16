package usecase

import (
	"context"

	"go.uber.org/zap"

	"github.com/mchaelhuang/aquafarm/internal"
	"github.com/mchaelhuang/aquafarm/internal/entity"
)

type StatsUC struct {
	logger    *zap.Logger
	statsRepo internal.StatsRepo
}

func NewStatsUC(
	logger *zap.Logger,
	statsRepo internal.StatsRepo,
) *StatsUC {
	return &StatsUC{
		logger:    logger,
		statsRepo: statsRepo,
	}
}

func (s *StatsUC) Add(ctx context.Context, info entity.StatsRequestInfo) error {
	if err := s.statsRepo.IncrEndpoint(ctx, info); err != nil {
		return err
	}
	if err := s.statsRepo.CollectUserAgent(ctx, info); err != nil {
		return err
	}

	return nil
}

func (s *StatsUC) Get(ctx context.Context) (map[string]entity.StatsResult, error) {
	endpoints, err := s.statsRepo.GetEndpointCount(ctx)
	if err != nil {
		return nil, err
	}

	results := map[string]entity.StatsResult{}
	for endpoint, count := range endpoints {
		var unique int
		unique, err = s.statsRepo.GetUniqueAgentCount(ctx, endpoint)
		if err != nil {
			return nil, err
		}

		results[endpoint] = entity.StatsResult{
			Count:           count,
			UniqueUserAgent: unique,
		}
	}

	return results, nil
}
