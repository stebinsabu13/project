package repository

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/domain"
	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/pb"
	"github.com/stebin13/x-tentioncrew/microservice-1/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type userRepo struct {
	DB  *gorm.DB
	RDB *redis.Client
}

func NewUserRepo(db *gorm.DB, rdb *redis.Client) interfaces.UserRepo {
	return &userRepo{
		DB:  db,
		RDB: rdb,
	}
}

func (c *userRepo) CreateUser(ctx context.Context, user domain.User) (string, error) {
	if err := c.DB.Create(&user).Error; err != nil {
		return "", err
	}
	userid := strconv.Itoa(int(user.ID))
	userjson, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	if err := c.RDB.Set(ctx, userid, userjson, 10*time.Minute).Err(); err != nil {
		return "", err
	}
	log.Println("userid", userid)
	return userid, nil
}

func (c *userRepo) GetUser(ctx context.Context, userid string) (domain.User, error) {
	var user domain.User

	userjson, err := c.RDB.Get(ctx, userid).Bytes()
	if err != nil {
		if err == redis.Nil {
			if err1 := c.DB.Where("id=?", userid).Find(&user).Error; err1 != nil {
				return user, err1
			}
			if user.ID == 0 {
				return user, errors.New("user doesn't exists")
			}
			userJson, _ := json.Marshal(user)
			if err2 := c.RDB.Set(ctx, userid, userJson, 10*time.Minute).Err(); err2 != nil {
				return user, err2
			}
			log.Println("from postgres", user)
			return user, nil
		}
		return user, err
	}
	json.Unmarshal(userjson, &user)
	log.Println("inside repo user=", user)
	return user, nil
}

func (c *userRepo) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	log.Println("inside repo update user", user)
	if err := c.DB.Where("id=?", user.ID).UpdateColumns(&user).Error; err != nil {
		return user, err
	}
	_, err := c.RDB.Get(ctx, strconv.Itoa(int(user.ID))).Bytes()
	if err != nil {
		if err == redis.Nil {
			log.Println("doesn't exists in the redis cache")
			return user, nil
		}
		return user, err
	}
	userJson, _ := json.Marshal(user)
	log.Println("exists in redis cache and updated the cache")
	if err1 := c.RDB.Set(ctx, strconv.Itoa(int(user.ID)), userJson, 10*time.Minute).Err(); err1 != nil {
		return user, err1
	}
	return user, nil
}

func (c *userRepo) DeleteUser(ctx context.Context, userid string) error {
	if err := c.DB.Where("id=?", userid).Delete(&domain.User{}).Error; err != nil {
		return err
	}
	_, err := c.RDB.Get(ctx, userid).Bytes()
	if err != nil {
		if err == redis.Nil {
			log.Println("doesn't exists in the redis cache")
			return nil
		}
		return err
	}
	if err := c.RDB.Del(ctx, userid).Err(); err != nil {
		return err
	}
	log.Println("deleted from redis")
	return nil
}

func (c *userRepo) Users() ([]*pb.User, error) {
	var users []*pb.User
	if err := c.DB.Find(&users).Error; err != nil {
		return users, err
	}
	return users, nil
}
