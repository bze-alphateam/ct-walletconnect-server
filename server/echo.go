package server

import (
	"context"
	"fmt"
	"github.com/bze-alphateam/ct-walletconnect-server/config"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"log"
)

func StartServer(cfg *config.App, logger logrus.FieldLogger) {
	e := echo.New()
	//TODO: implement rate limiter on both WS + HTTP endpoints
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	redisCl, err := getRedisClient(cfg.Redis)
	if err != nil {
		log.Fatal(err)
	}

	ctrlFactory, err := NewControllerFactory(logger, redisCl)
	if err != nil {
		log.Fatal(err)
	}

	registerRoutes(e, ctrlFactory)

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Server.GetHost(), cfg.Server.GetPort())))
}

func getRedisClient(r *config.Redis) (*redis.Client, error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.Address,
		Password: r.Password,
		DB:       r.Database,
	})

	// Perform a "ping" to check connectivity
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("error connecting to Redis: %s", err)
	}

	return rdb, nil
}
