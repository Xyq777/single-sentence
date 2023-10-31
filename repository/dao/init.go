package dao

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	c "single-sentence/config"
	"single-sentence/repository/model"
)

var RedisDao redis.Conn

var MysqlDao *gorm.DB

func init() {

	//初始化链接redis
	var err error
	RedisDao, err = redis.Dial("tcp", c.Config.RedisIP)
	if err != nil {
		log.Panicln("redis链接失败", err)
	}
	//初始化链接mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", c.Config.Username, c.Config.Password, c.Config.Host, c.Config.Port, c.Config.Dbname, c.Config.Timeout)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("连接数据库失败" + err.Error())

	}
	MysqlDao = db
	err = MysqlDao.AutoMigrate(&model.Xyq777{}, &model.User{})
	if err != nil {
		panic(err)
	}
	return
}
