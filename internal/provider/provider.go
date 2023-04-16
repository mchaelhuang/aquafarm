package provider

import (
	"fmt"
	"sync"

	redis "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mchaelhuang/aquafarm/internal/config"
)

var (
	logger     *zap.Logger
	loggerOnce sync.Once

	db     *gorm.DB
	dbOnce sync.Once

	redisClient     *redis.Client
	redisClientOnce sync.Once
)

func LoggerProvider() *zap.Logger {
	loggerOnce.Do(func() {
		logger, _ = zap.NewDevelopment()
		_ = zap.ReplaceGlobals(logger)
	})
	return logger
}

func DatabaseProvider(cfg *config.Cfg) *gorm.DB {
	dbOnce.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
			cfg.Storage.Postgres.Host,
			cfg.Storage.Postgres.Port,
			cfg.Storage.Postgres.User,
			cfg.Storage.Postgres.Password,
			cfg.Storage.Postgres.Name,
			"Asia/Jakarta",
		)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.Fatal("failed open database", zap.Error(err))
		}
	})
	return db
}

func RedisProvider(cfg *config.Cfg) *redis.Client {
	redisClientOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     cfg.Storage.Redis.Address,
			Password: cfg.Storage.Redis.Password,
		})
	})
	return redisClient
}
