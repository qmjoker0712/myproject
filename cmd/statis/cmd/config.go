package cmd

import (
	"fmt"
	"myproject/cmd/utils"
	"myproject/config"
	"myproject/db"
	"myproject/node"
	"myproject/rpc"
	"myproject/statistical"
	"myproject/work"
)

func makeConfigNodeStatis() (*node.Node, config.SConfig) {
	// get config
	cfg := config.GetConfig()

	utils.SetLogConfig(&cfg.Log)
	utils.SetMySQLConfig(&cfg.Mysql)
	utils.SetRedisConfig(&cfg.Redis)

	// set rpc logger
	utils.SetNodeConfig(&cfg.Node)
	// Warn: rpc logger name must be rpc
	fmt.Printf("init logger start...\n")
	db.SetModuleLogger("db", &cfg.Log)
	node.SetModuleLogger("node", &cfg.Log)
	rpc.SetModuleLogger("rpc", &cfg.Log)
	statistical.SetModuleLogger("statistical", &cfg.Log)
	work.SetModuleLogger("work", &cfg.Log)
	fmt.Printf("init logger end...\n")
	stack, err := node.New(&cfg.Node)
	if err != nil {
		utils.Fatalf("Failed to create the protocol stack: %v", err)
	}

	return stack, *cfg
}

func makeFullNodeStatis() (*node.Node, config.SConfig) {
	// set config
	stack, conf := makeConfigNodeStatis()
	// register service
	utils.RegisterStatisticalService(stack)

	return stack, conf
}
