package main

import (
	"github.com/bze-alphateam/ct-walletconnect-server/config"
	"github.com/bze-alphateam/ct-walletconnect-server/server"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

const (
	defaultCfgPath = "./config.yml"
)

func main() {
	cfgPath := defaultCfgPath
	if len(os.Args) > 1 {
		cfgPath = os.Args[1]
	}

	cfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	logger := logrus.New()

	server.StartServer(cfg, logger)
}
