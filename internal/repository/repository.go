package repository

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"vega-server/pkg/config"
	"vega-server/pkg/log"

	"context"
	"fmt"
	"os"
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
	db, err := gorm.Open(sqlite.Open("storage/vega-db.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func NewRedis(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		//Addr:     config.GetString("repo.rdb.addr"),
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		//Password: config.GetString("repo.rdb.password"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       config.GetInt("repo.rdb.db"),
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return rdb
}
