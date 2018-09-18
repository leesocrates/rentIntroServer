package main

import (
	"rentIntroServer/src/config"
	"rentIntroServer/src/controllers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/wonktnodi/go-utils/log"
	"rentIntroServer/src/constant"
)

var logger log.Logger

func main() {
	if constant.Debug == 1 {
		gin.SetMode(gin.DebugMode)
		logger = log.Start(log.LogFilePath("./log"), log.EveryHour, log.AlsoStdout, log.LogFlags(log.Lfunc|log.Lfile|log.Lline))
	} else {
		gin.SetMode(gin.ReleaseMode)
		logger = log.Start(log.LogFilePath("./log"), log.EveryHour, log.LogFlags(log.Lfunc|log.Lfile|log.Lline))
	}
	defer logger.Stop()

	initConfig()

	r := gin.New()
	initRouters(r)
	r.Run("127.0.0.1:8888")
}

func initRouters(r *gin.Engine){
	shares := new(controllers.ShareController)
	r.GET("/shares", shares.Get)
	r.GET("/users/:userId/shares", shares.Get)
	r.PUT("/users/:userId/shares", shares.Put)
}

var generalCfg *viper.Viper
var dbCfg *viper.Viper

func initConfig() {
	var err error

	generalCfg, err = config.InitConfig("users")
	if err != nil {
		log.Fatalf("failed to load service config, err: %s", err.Error())
		return
	}

	dbCfg, err = config.InitConfig("db")
	if err != nil {
		log.Fatalf("failed to load database config, err: %s", err.Error())
		return
	}

	if err := config.GetSettingsByKey(generalCfg, "notifications", &config.SysNotify); err != nil {
		log.Fatalf("error to read general setting, %s", err)
	}

	if err := config.GetSettingsByKey(dbCfg, "dbservers", &config.DbSetting); err != nil {
		log.Fatalf("Error reading db config file, %s", err)
	}

	if err := config.GetSettingsByKey(generalCfg, "redis", &config.RedisSetting); err != nil {
		log.Fatalf("error to read redis setting, %s", err);
	}
}