package configs

import (
	"database/sql"
	"fmt"
	"log"

	"context"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	DB            *sql.DB
	Redis         *redis.Client
	JWTSecret     string
	SessionSecret string
)

func init() {
	// 读取配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// 配置 PostgreSQL 连接
	postgresConf := viper.GetStringMapString("postgres")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgresConf["host"], postgresConf["port"], postgresConf["user"], postgresConf["password"], postgresConf["dbname"])

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Database ping failed: %s", err)
	}

	// 配置 Redis 连接
	redisConf := viper.GetStringMapString("redis")
	Redis = redis.NewClient(&redis.Options{
		Addr:     redisConf["address"],
		Password: redisConf["password"],
		DB:       viper.GetInt("redis.db"),
	})

	if err = Redis.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Error connecting to Redis: %s", err)
	}
	// 加载 JWT 密钥
	JWTSecret = viper.GetString("jwt_secret")
	SessionSecret = viper.GetString("session_secret")
}
