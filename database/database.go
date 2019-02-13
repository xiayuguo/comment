package database

import (
	. "comment/config"
	"comment/util"
	"fmt"
	"gopkg.in/mgo.v2"
	"time"
)

var (
	Session *mgo.Session
	DB      *mgo.Database
)

func init() {
	dialInfo := &mgo.DialInfo{
		Addrs:     []string{util.Join(util.S{Db.Host, Db.Port}, ":")}, // 数据库地址
		Timeout:   60 * time.Second,                                   // 连接超时时间
		Database:  Db.Database,                                        // 数据库
		Username:  Db.User,                                            // mongodb 用户名
		Password:  Db.Password,                                        // mongodb 密码
		PoolLimit: 100,                                                // 连接池数量
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		// 创建会话失败
		fmt.Printf("Create Session Fail: %s\n", err)
	}
	Session = session
	Session.SetMode(mgo.Eventual, true)
	DB = Session.DB(Db.Database)
}
