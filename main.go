package main

import (
	"github.com/sunmi-OS/gocore/v2/utils"
	"github.com/urfave/cli/v2"

	"nacos-endpoint/cmd"
	"nacos-endpoint/conf"

	"log"
	"os"
)

func main()  {
	utils.PrintBanner(conf.PROJECT_NAME)

	// 配置cli参数
	app := cli.NewApp()
	app.Name = conf.PROJECT_NAME
	app.Usage = conf.PROJECT_NAME
	app.Version = conf.PROJECT_VERSION
	// 指定命令运行的函数
	app.Commands = []*cli.Command{
		cmd.Api,
	}
	// 启动cli
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}




