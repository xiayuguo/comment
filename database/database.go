package database

import (
	. "comment/config"
	"comment/util"
	"gopkg.in/mgo.v2"
	"time"
)

var mongoDBDialInfo = &mgo.DialInfo{
	Addrs:     []string{util.Join([]interface{}{Db.Host, Db.Port}, ":")}, // 数据库地址
	Timeout:   60 * time.Second,                                          // 连接超时时间
	Database:  Db.Database,                                               // 数据库
	Username:  Db.User,                                                   // mongodb 用户名
	Password:  Db.Password,                                               // mongodb 密码
	PoolLimit: 100,                                                       // 连接池数量
}
