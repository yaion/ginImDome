package initialize

import (
	"ginIm/conf"
	"ginIm/initialize/log"
	"ginIm/initialize/rabbitmQ"
	"ginIm/initialize/redis"
	"ginIm/models/mysql"
	"github.com/gin-gonic/gin"
)

func Init() {
	// 1.配置
	conf.InitConf()
	// 2.日志
	log.InitLogger()
	// 3.数据库
	mysql.InitDb()
	// 4.redis
	redis.InitRedis()
	// 5.其他
	rabbitmQ.InitRabbitMQ()
	// 6.gin
	gin.SetMode(conf.Conf.App.Debug)
	gin.New()
}
