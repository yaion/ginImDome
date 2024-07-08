package cmd

import (
	"ginIm/conf"
	"github.com/gin-gonic/gin"
)

func Run() {
	conf.Init()
	gin.SetMode(conf.Conf.App.Debug)
	gin.New()

}
