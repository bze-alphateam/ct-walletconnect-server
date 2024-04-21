package server

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type ControllerFactory struct {
	logger      logrus.FieldLogger
	redisClient *redis.Client
}

func NewControllerFactory(logger logrus.FieldLogger, redisClient *redis.Client) (*ControllerFactory, error) {
	if logger == nil || redisClient == nil {
		return nil, fmt.Errorf("invalid dependencies provided to controller factory")
	}

	return &ControllerFactory{
		logger:      logger,
		redisClient: redisClient,
	}, nil
}
