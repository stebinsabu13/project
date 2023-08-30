package db

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/config"
	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initdb(cfg *config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.Db_host, cfg.Db_username, cfg.Db_password, cfg.Db_name, cfg.Db_port)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if dbErr != nil {
		log.Fatalln(dbErr)
	}
	db.AutoMigrate(&domain.User{})

	return db, dbErr
}

func InitRedis(cfg *config.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis_Address, // use default Addr
		Password: "",                // no password set
		DB:       0,                 // use default DB
	})
	pong, err := rdb.Ping(context.Background()).Result()
	fmt.Println(pong, err)
	return rdb, nil
}
