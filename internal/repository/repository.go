package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"vega-server/pkg/config"
	"vega-server/pkg/log"

	"context"
)

type Repository struct {
	db     *gorm.DB
	rdb    *redis.Client
	logger *log.Logger
}

func NewRepository(db *gorm.DB, rdb *redis.Client, logger *log.Logger) *Repository {
	return &Repository{
		db:     db,
		rdb:    rdb,
		logger: logger,
	}
}

func NewDB(config *config.Config) *gorm.DB {
	dsn := config.GetString("repo.db.dsn")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func NewRedis(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetString("repo.rdb.addr"),
		Password: config.GetString("repo.rdb.password"),
		DB:       config.GetInt("repo.rdb.db"),
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return rdb
}
