package cmd

import (
	"fmt"
	"myproject/cmd"
	"myproject/cmd/utils"
	"myproject/config"
	"myproject/db/mysql"
	"myproject/db/redis"
	"myproject/node"
	"myproject/work"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/s3dteam/go-toolkit/log/logruslogger"
	"github.com/spf13/cobra"
)

var (
	configFile  *string
	versionFlag *bool
)

type Service struct {
	router *gin.Engine
}

var statisCmd = &cobra.Command{
	Use:   "start",
	Short: "start the statis",
	Long: `usage example:
	statis(.exe) start -c config/statis.toml
	start the statis`,
	Run: func(cmd *cobra.Command, args []string) {
		// init config for global
		config.Init(configFile)
		// init profile
		node, conf := makeFullNodeStatis()
		// init mysql
		mysql.Init(&conf.Mysql)
		// init redis
		redis.Init(&conf.Redis)
		loger := logruslogger.GetLoggerWithOptions("ginloger", &config.GetConfig().Log)

		//start gin
		go func() {
			s := &Service{}
			s.router = gin.Default()
			s.router.Use(cors.Default())
			s.registerHandlers()
			port := ":" + config.GetConfig().GinPort.GinPort
			func() {
				if err := s.router.Run(port); err != nil {
					loger.Error("could not start server: %v", err)
				}
			}()
		}()
		loger.Info(" start service ok!")
		startNode(node)
		fmt.Printf("\nconfig: %#v\n\n", conf)
		node.Wait()
	},
}

func init() {
	// add version cmd
	rootCmd.AddCommand(cmd.VersionCmd)
	versionFlag = cmd.VersionCmd.Flags().BoolP("version", "v", true, "statis config file (required)")
	cmd.VersionFlag = versionFlag

	rootCmd.AddCommand(statisCmd)
	configFile = statisCmd.Flags().StringP("config", "c", "", "statis config file (required)")
	statisCmd.MarkFlagRequired("config")
}
func startNode(stack *node.Node) {
	// Start up the node itself
	utils.StartNode(stack)
}

func (s *Service) registerHandlers() {
	routerGroup := s.router.Group("/api")
	routerGroup.GET("/test", work.Handler)
}
