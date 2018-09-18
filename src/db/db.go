package db

import (
	"github.com/go-xorm/xorm"
	"github.com/wonktnodi/go-utils/log"
	"io"
	"fmt"
	"github.com/go-xorm/core"
	"time"
	"rentIntroServer/src/config"
	"github.com/gomodule/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"rentIntroServer/src/constant"
)

var usersDB *xorm.Engine = nil
var roomsDB *xorm.Engine = nil
var productsDB *xorm.Engine = nil
var ordersDB *xorm.Engine = nil
var promotionsDB *xorm.Engine = nil

var redisPool *redis.Pool = nil

func InitDatabase(out io.Writer) {
	var err error

	srvInfo, ok := config.DbSetting["users"]
	if !ok {
		log.Errorf("there's no db setting for users")
		return
	}
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&charset=utf8",
		srvInfo.User, srvInfo.Pwd, srvInfo.IP, srvInfo.Port, srvInfo.Database)
	usersDB, err = xorm.NewEngine("mysql", connectString)
	if err != nil {
		return
	}
	usersDB.SetMaxOpenConns(20)

	srvInfo, ok = config.DbSetting["rooms"]
	if !ok {
		log.Errorf("there's no db setting for rooms")
		return
	}
	connectString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&charset=utf8",
		srvInfo.User, srvInfo.Pwd, srvInfo.IP, srvInfo.Port, srvInfo.Database)
	roomsDB, err = xorm.NewEngine("mysql", connectString)
	if err != nil {
		return
	}
	roomsDB.SetMaxOpenConns(20)

	srvInfo, ok = config.DbSetting["products"]
	if !ok {
		log.Errorf("there's no db setting for rooms")
		return
	}
	connectString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&charset=utf8",
		srvInfo.User, srvInfo.Pwd, srvInfo.IP, srvInfo.Port, srvInfo.Database)
	productsDB, err = xorm.NewEngine("mysql", connectString)
	if err != nil {
		return
	}
	productsDB.SetMaxOpenConns(20)

	srvInfo, ok = config.DbSetting["orders"]
	if !ok {
		log.Errorf("there's no db setting for rooms")
		return
	}
	connectString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&charset=utf8",
		srvInfo.User, srvInfo.Pwd, srvInfo.IP, srvInfo.Port, srvInfo.Database)
	ordersDB, err = xorm.NewEngine("mysql", connectString)
	if err != nil {
		return
	}
	ordersDB.SetMaxOpenConns(20)

	srvInfo, ok = config.DbSetting["promotions"]
	if !ok {
		log.Errorf("there's no db setting for promotions")
		return
	}
	connectString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True&charset=utf8",
		srvInfo.User, srvInfo.Pwd, srvInfo.IP, srvInfo.Port, srvInfo.Database)
	promotionsDB, err = xorm.NewEngine("mysql", connectString)
	if err != nil {
		return
	}
	promotionsDB.SetMaxOpenConns(20)

	if constant.Debug == 1 {
		usersDB.ShowSQL(true)
		roomsDB.ShowSQL(true)
		productsDB.ShowSQL(true)
		ordersDB.ShowSQL(true)
		promotionsDB.ShowSQL(true)
	}

	dbLogger := xorm.NewSimpleLogger(out)
	if dbLogger != nil {
		dbLogger.SetLevel(core.LOG_INFO)
		usersDB.SetLogger(dbLogger)
		roomsDB.SetLogger(dbLogger)
		productsDB.SetLogger(dbLogger)
		ordersDB.SetLogger(dbLogger)
		promotionsDB.SetLogger(dbLogger)
	}
}

func InitRedis() {
	redisPool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%d", config.RedisSetting.IP, config.RedisSetting.Port))
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func Release() {
	if usersDB != nil {
		usersDB.Close()
		usersDB = nil
	}

	if roomsDB != nil {
		roomsDB.Close()
		roomsDB = nil
	}

	if productsDB != nil {
		productsDB.Close()
		productsDB = nil
	}

	if ordersDB != nil {
		ordersDB.Close()
		ordersDB = nil
	}

	if promotionsDB != nil {
		promotionsDB.Close()
		promotionsDB = nil
	}

	if redisPool != nil {
		redisPool.Close()
		redisPool = nil
	}
}

func GetUsersDB() *xorm.Engine {
	return usersDB
}

func GetRoomsDB() *xorm.Engine {
	return roomsDB
}

func GetProductsDB() *xorm.Engine {
	return productsDB
}

func GetOrdersDB() *xorm.Engine {
	return ordersDB
}

func GetPromotionsDB() *xorm.Engine {
	return promotionsDB
}

func GetRedisConn() (conn redis.Conn, code int) {
	if nil == redisPool {
		code = constant.ErrCodeInternal
		return
	}
	conn = redisPool.Get()
	return
}

func CheckDB() {
	err := usersDB.Ping()
	if err != nil {
		log.Fatalf("failed to ping users db, err: %v", err)
	}

	err = roomsDB.Ping()
	if err != nil {
		log.Fatalf("failed to ping rooms db, err: %v", err)
	}

	err = productsDB.Ping()
	if err != nil {
		log.Fatalf("failed to ping products db, err: %v", err)
	}

	err = ordersDB.Ping()
	if err != nil {
		log.Fatalf("failed to ping orders db, err: %v", err)
	}

	err = promotionsDB.Ping()
	if err != nil {
		log.Fatalf("failed to ping promotions db, err: %v", err)
	}

	conn, code := GetRedisConn()
	if code == constant.Success {
		_, err = conn.Do("PING")
	}

	if nil != err || code != constant.Success {
		log.Errorf("failed to check redis, code[%d], err: %s", code, err.Error())
	}
	log.Infof("db connections are healthy")
}