package cmd

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
	"net/http"
	"strings"
)

var Api = &cli.Command{
	Name:  "api",
	Usage: "run api",
	Subcommands: []*cli.Command{
		{
			Name:   "start",
			Usage:  "start api",
			Action: api,
		},
	},
}

func api(c *cli.Context) error {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.Use(gzip.Gzip(gzip.DefaultCompression))

	// 根目录健康检查
	r.Any("/nacos/serverlist", func(c *gin.Context) {
		hosts := strings.Split(c.Request.Host,".")
		if len(hosts) > 2 {
			c.String(http.StatusOK,hosts[0]+".mse.aliyuncs.com")
		}else{
			c.String(http.StatusInternalServerError,"Parameter resolution failed")
		}
	})

	err := endless.ListenAndServe(":8088", r)
	if err != nil {
		return err
	}
	return nil
}
