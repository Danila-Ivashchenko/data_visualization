package api

import (
	"github.com/gin-gonic/gin"
)

type CoinHandler interface {
	CreateBar(c *gin.Context)
	UploadFile(c *gin.Context)
	CreateLineChart(c *gin.Context)
	CreateScatter(c *gin.Context)
	CreateWordCloud(c *gin.Context)
}

type api struct {
	coinHandler CoinHandler
	server      *gin.Engine
	port        string
}

func New(p CoinHandler) *api {
	api := &api{
		coinHandler: p,
		port:        "8080",
	}

	api.server = gin.New()
	api.server.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/pathsNotToLog/"),
		gin.Recovery(),
	)
	api.bind()

	return api
}

func (a *api) bind() {
	a.server.POST("/coin/file", a.coinHandler.UploadFile)
	a.server.GET("/coin/bar/:key", a.coinHandler.CreateBar)
	a.server.GET("/coin/line/:key", a.coinHandler.CreateLineChart)
	a.server.GET("/coin/scatter/:key", a.coinHandler.CreateScatter)
	a.server.GET("/coin/word_cloud/:key", a.coinHandler.CreateWordCloud)
}

func (a *api) Run() error {
	return a.server.Run(":" + a.port)
}
