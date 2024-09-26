package model

import (
	"fmt"
	"ginblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	db, err = gorm.Open(mysql.New(mysql.Config{ //todo 使用格式化输出
		DSN:               utils.DbUser + ":" + utils.DbPassword + "@tcp(" + utils.DBAddress + ":" + utils.DbHost + ")/" + utils.DbName + "?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize: 256,
	}), &gorm.Config{
		SkipDefaultTransaction: false, //禁用默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //使用单数表名
		},
		//DisableForeignKeyConstraintWhenMigrating: true, //禁用外键约束
	})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}
	err = db.AutoMigrate(&User{}, &Article{}, &Category{})
	//fmt.Println(db.Migrator().HasTable(&User{}))
	if err != nil {
		fmt.Println("迁移表失败", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)                  //设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)                 //设置数据库的最大连接数量
	sqlDB.SetConnMaxLifetime(10 * time.Second) //设置连接的最大可复用时间

}
