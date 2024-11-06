package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/MarskTM/financial_report_server/env"
	"github.com/MarskTM/financial_report_server/infrastructure/model"
	redis "github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

// InitRedis open connection to redis server
func ConnectRedis(config env.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// --------------------------------Redis Client DAO--------------------------------
func FetchAuth(accessUUID string) (*model.AccessCache, error) {
	val, err := redisClient.Get(context.Background(), accessUUID).Result()
	if err != nil {
		log.Fatalf("Could not get user from Redis: %v", err)
		return nil, err
	}

	// Deserialize JSON string back to struct
	var accessData model.AccessCache
	err = json.Unmarshal([]byte(val), &accessData)
	if err != nil {
		log.Fatalf("Could not unmarshal JSON: %v", err)
		return nil, err
	}

	return &accessData, nil
}

func StoreAuth(accessUUID string, data model.AccessCache) error {
	// Serialize the struct to JSON
	accessJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Could not marshal JSON: %v", err)
	}

	// Store the JSON string in Redis
	err = redisClient.Set(context.Background(), accessUUID, accessJSON, env.CacheExpiresAt).Err()
	if err != nil {
		log.Fatalf("Could not set user in Redis: %v", err)
	}
	fmt.Println("User stored in Redis")

	return nil
}
