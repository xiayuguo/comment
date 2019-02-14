package database

import (
	"comment/config"
	"comment/util"
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

// 定义会话和数据库连接
var (
	Session *mgo.Session
	DB      *mgo.Database
)

func init() {
	db := config.Db
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{util.Join(util.S{db.Host, db.Port}, ":")}, // 数据库地址
		Timeout:   60 * time.Second,                                   // 连接超时时间
		Database:  db.Database,                                        // 数据库
		Username:  db.User,                                            // mongodb 用户名
		Password:  db.Password,                                        // mongodb 密码
		PoolLimit: 100,                                                // 连接池数量
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		// 创建会话失败
		fmt.Printf("Create Session Fail: %s\n", err)
	}
	Session = session
	Session.SetMode(mgo.Eventual, true)
	DB = Session.DB(db.Database)
}
