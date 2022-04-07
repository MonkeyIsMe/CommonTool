package client

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

// MysqlProxy mysql的proxy
type MysqlProxy struct {
	Account  string
	Password string
	IP       string
	Port     string
	DBName   string
}

var _db *gorm.DB

// NewMysqlClient 创建一个mysql的client
func NewMysqlClient(proxy MysqlProxy) (*sql.DB, error) {
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", proxy.Account, proxy.Password,
		proxy.IP, proxy.Port, proxy.DBName)
	db, err := sql.Open("mysql", sourceName)
	if err != nil {
		log.Fatalf("New Mysql Client Error %+v", err)
		return nil, err
	}
	return db, nil
}

// NewClient 通过gorm 创建一个mysql的client
func NewClient(proxy MysqlProxy) (*gorm.DB, error) {
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		proxy.Account, proxy.Password, proxy.IP, proxy.Port, proxy.DBName)
	//连接MYSQL
	db, err := gorm.Open(mysql.Open(sourceName), &gorm.Config{})

	if err != nil {
		log.Fatalf("连接数据库失败, error=" + err.Error())
		return nil, err
	}

	return db, nil
}

// NewMysqlPool 创建一个gorm 的数据库连接池
func NewMysqlPool(proxy MysqlProxy) (*gorm.DB, error) {
	sourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		proxy.Account, proxy.IP, proxy.IP, proxy.Port, proxy.DBName)
	var err error
	_db, err = gorm.Open(mysql.Open(sourceName), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败, error=" + err.Error())
		return nil, err
	}

	sqlDB, _ := _db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	return _db, nil
}
