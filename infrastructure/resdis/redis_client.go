package resdis

import (
	"fmt"
	"strconv"

	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/utils"
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

// InitRedis open connection to redis server
func ConnectRedis(config env.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetRedisClient() (*redis.Client, error) {
	if redisClient == nil {
		err := fmt.Errorf("Need to create a redis client first")
		return nil, err
	}
	return redisClient, nil
}

// --------------------------------Redis Client DAO--------------------------------
// DeleteAuth delete pair key value has key same givenUUID
func DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := redisClient.Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

func ClearAuth(userID uint) {
	list, _ := redisClient.Keys(utils.PatternGet(userID)).Result()
	for _, key := range list {
		redisClient.Del(key)
	}
}

func FetchAuth(accessUUID string) (uint, error) {
	userIDStr, err := redisClient.Get(accessUUID).Result()
	if err != nil {
		return 0, err
	}

	userID, _ := strconv.ParseUint(userIDStr, 10, 64)
	return uint(userID), nil
}
