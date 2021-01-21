package main

import (
	"message/handler"
	"message/package/dingtalk"
	"context"
	"flag"
	"github.com/xinzf/gokit/logger"
	"github.com/xinzf/gokit/registry"
	"github.com/xinzf/gokit/server"
	"github.com/xinzf/gokit/storage"
	"github.com/xinzf/gokit/utils"
)

const PROJECT_NAME = "Message"

var Host string
var Port int
var DbHost string
var DbUser string
var DbPassWord string
var DbName string
var RegisterAddr string
var RedisAddr string
var RedisPswd string
var RedisDB int
var LogPath string
var AppKey string
var AppSecret string
var debug bool

func init() {
	flag.StringVar(&Host, "host", "127.0.0.1", "host address")
	flag.IntVar(&Port, "port", 8092, "host port")
	flag.StringVar(&DbHost, "dbhost", "127.0.0.1:3306", "database ip and port")
	flag.StringVar(&DbUser, "dbuser", "root", "database user name")
	flag.StringVar(&DbPassWord, "dbpswd", "123456", "database user password")
	flag.StringVar(&DbName, "dbname", "xunray_message", "choose your database")
	flag.StringVar(&RegisterAddr, "registry", "127.0.0.1:8500", "register address")
	flag.StringVar(&LogPath, "logpath", "", "write your log address")
	flag.StringVar(&RedisAddr, "redisaddr", "127.0.0.1:6379", "input redis address")
	flag.StringVar(&RedisPswd, "redispswd", "", "input redis password")
	flag.StringVar(&AppKey, "appkey", "ding19bgmqhsevsorvyq", "input dingtalk app key")
	flag.StringVar(&AppSecret, "appsecret", "ihuzQsVaHuoMKlMSW14RNfzv93s3-Xeb-Viz4tJR_OCI_SCEH_A2WevzG1ERz1rM", "input dingtalk app secret")
	flag.IntVar(&RedisDB, "redisdb", 12, "input redis db index")
	flag.BoolVar(&debug, "debug", true, "debug mode")
	flag.Parse()
}

func main() {
	loggerLevel := logger.DebugLevel
	if debug {
		loggerLevel = logger.DebugLevel
	}
	logger.DefaultLogger.Init(
		logger.Level(loggerLevel),
		logger.ProjectName(PROJECT_NAME),
		logger.Filename(LogPath),
	)

	reg := registry.NewConsul(registry.Addrs(RegisterAddr))
	server.New(
		server.Name(PROJECT_NAME),
		server.AllowHeaders("X-USER", "X-Token"),
		server.Registry(reg),
		server.Logger(logger.DefaultLogger),
		server.Host(Host),
		server.Port(Port),
	)

	server.Register("message", new(handler.Message))
	server.Register("enums", new(handler.Enums))

	server.Run(context.TODO(), func() error {
		utils.GORM_JSON_FIELD = true
		return nil
	}, func() error {
		return storage.DB.Init(
			storage.DbConfig(DbHost, DbUser, DbPassWord, DbName),
			storage.DbLogger(logger.DefaultLogger),
		)
	}, func() error {
		return storage.Redis.Init(
			storage.RedisLogger(logger.DefaultLogger),
			storage.RedisConfig(RedisAddr, RedisPswd, RedisDB),
		)
	}, func() error {
		dingtalk.Option.AppSecret = AppSecret
		dingtalk.Option.AppKey = AppKey
		return nil
	})
}
