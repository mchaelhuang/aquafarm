package repository

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"

	"github.com/mchaelhuang/aquafarm/internal/entity"
)

type StatsRepo struct {
	redis *redis.Client
}

func NewStatsRepo(
	redis *redis.Client,
) *StatsRepo {
	return &StatsRepo{
		redis: redis,
	}
}

const (
	_KeyEndpoint        = "stats:endpoint"
	_KeyPrefixUserAgent = "stats:useragent:"
)

func (s *StatsRepo) IncrEndpoint(ctx context.Context, info entity.StatsRequestInfo) error {
	return s.redis.HIncrBy(ctx, _KeyEndpoint, endpointKey(info), 1).Err()
}

func (s *StatsRepo) CollectUserAgent(ctx context.Context, info entity.StatsRequestInfo) error {
	return s.redis.HSet(ctx, _KeyPrefixUserAgent+endpointKey(info), info.UserAgent, 1).Err()
}

// endpointKey returns formatted unique key to count endpoint. eg: "GET /v1/farm"
func endpointKey(info entity.StatsRequestInfo) string {
	return info.Method + " " + info.Endpoint
}

func (s *StatsRepo) GetEndpointCount(ctx context.Context) (map[string]int, error) {
	results, err := s.redis.HGetAll(ctx, _KeyEndpoint).Result()
	if err != nil {
		return nil, err
	}

	list := map[string]int{}
	for endpoint, count := range results {
		c, _ := strconv.Atoi(count)
		list[endpoint] = c
	}
	return list, nil
}

func (s *StatsRepo) GetUniqueAgentCount(ctx context.Context, endpointKey string) (int, error) {
	count, err := s.redis.HLen(ctx, _KeyPrefixUserAgent+endpointKey).Result()
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
