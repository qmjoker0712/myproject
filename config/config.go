package config

import (
	"fmt"
	"sync"

	"myproject/node"
	"myproject/types"

	"github.com/s3dteam/go-toolkit/config"
	"github.com/s3dteam/go-toolkit/db/mysqldao"
	"github.com/s3dteam/go-toolkit/log/logruslogger"
)

var sconfig *SConfig
var lock sync.RWMutex

type ProfileConf struct {
	Profileaddress string `toml:"profileaddress"`
	Open           bool   `toml:"open"`
}

// SConfig d4d  config
type SConfig struct {
	Node    node.Config
	Log     logruslogger.Options
	Profile ProfileConf `toml:"profile"`
	Mysql   mysqldao.MysqlConifg
	Redis   types.RedisOptions
	Common  types.CommonOptions

	SMS      types.SMSOptions   `toml:"sms"`
	Email    types.EmailOptions `toml:"email"`
	Kmq      KmqConfig          `toml:"kmq"`
	GinPort  types.GinPort      `toml:"ginPort"`
	External map[string]string
}

// KmqConfig -
type KmqConfig struct {
	KafkaAddresses   []string `toml:"KafkaAddresses"`
	FetcherTopic     string   `toml:"FetcherTopic"`
	FetcherOffsetKey string   `toml:"FetcherOffsetKey"`
	SpiderTopic      string   `toml:"SpiderTopic"`
	SpiderOffsetKey  string   `toml:"SpiderOffsetKey"`
}

// Init init the config
func Init(configFile *string) {
	fmt.Printf("start init config...\n")
	lock.RLock()
	defer lock.RUnlock()

	// load file
	var cfg SConfig
	err := config.LoadConfig(*configFile, &cfg)
	if err != nil {
		panic(err)
	}
	sconfig = &cfg

	if sconfig.External != nil {
		// initExternalConf(sconfig.External, sconfig, filepath.Dir(*configFile))
	} else {
		fmt.Println("=============== External is nil")
	}

	fmt.Printf("init config successed\n")
}

// func initExternalConf(ex map[string]string, s *SConfig, prePath string) {
// 	for k, path := range ex {
// 		switch k {
// 		case "fetcher":
// 			allPath := prePath + "/" + path
// 			fmt.Println("load external[fetcher] config:", allPath)
// 			err := config.LoadConfig(allPath, &s.Fetcher)
// 			if err != nil {
// 				panic(err)
// 			}
// 			fmt.Println("=============== External[fetcher.toml] OK!")
// 			break
// 		case "sms":
// 			allPath := prePath + "/" + path
// 			fmt.Println("load external[sms] config:", allPath)
// 			err_sms := config.LoadConfig(allPath, &s.SMS)
// 			if err_sms != nil {
// 				panic(err_sms)
// 			}
// 			fmt.Println("=============== External[sms.toml] OK!")
// 			break
// 		case "email":
// 			allPath := prePath + "/" + path
// 			fmt.Println("load external[email] config:", allPath)
// 			err_email := config.LoadConfig(allPath, &s.Email)
// 			if err_email != nil {
// 				panic(err_email)
// 			}
// 			fmt.Println("=============== External[email.toml] OK!")
// 			break
// 		case "alert":
// 			allPath := prePath + "/" + path
// 			fmt.Println("load external[alert] config:", allPath)
// 			err_alert := config.LoadConfig(allPath, &s.Alert)
// 			if err_alert != nil {
// 				panic(err_alert)
// 			}
// 			fmt.Println("=============== External[alert.toml] OK!")
// 			break
// 		}
// 	}
// }

// GetConfig get global config
func GetConfig() *SConfig {
	lock.RLock()
	defer lock.RUnlock()
	return sconfig
}
