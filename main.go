package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/agentCard/api"
	"github.com/ethereum/agentCard/config"
	"github.com/ethereum/agentCard/db"
	"github.com/ethereum/agentCard/log"
	"github.com/sirupsen/logrus"
	"os"
)

const CONTRACTLEN = 42

var (
	conffile string
	env      string
)

func init() {
	flag.StringVar(&conffile, "conf", "config.yaml", "conf file")
	flag.StringVar(&env, "env", "prod", "Deploy environment: [ prod | test ]. Default value: prod")
}

func main() {
	var err error
	if config.Conf, err = config.LoadConfig("./conf"); err != nil {
		logrus.Error("🚀 Could not load environment variables")
		panic(err)
	}

	flag.Parse()

	err = log.Init("agentCard", &config.Conf)
	if err != nil {
		log.Fatal(err)
	}
	dbConnection, err := db.NewMysql(&config.Conf)
	if err != nil {
		logrus.Fatalf("connect to dbConnection error:%v", err)
	}

	apiService := api.NewApiService(dbConnection, &config.Conf)
	go apiService.Run()

	//listen kill signal
	closeCh := make(chan os.Signal, 1)

	for {
		select {
		case <-closeCh:
			fmt.Printf("receive os close sigal")
			return
		}
	}
}
