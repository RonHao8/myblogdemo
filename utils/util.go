package utils

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var db *sql.DB

func InitMysql() {

	fmt.Println("InitMysql....")
	driverName := beego.AppConfig.String("driverName")

	//注册数据库驱动
	orm.RegisterDriver(driverName, orm.DRMySQL)

	//数据库连接
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/cmsproject?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	db1, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		//创建用户表
		CreateTableWithUser()
		//创建文章表
		CreateTableWithArticle()
		//创建图片表
		CreateTableWithAlbum()
	}
}

// ModifyDB 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

// CreateTableWithUser 创建用户表
func CreateTableWithUser() {
	sql := `create table if not exists users(
    id int(4) primary key auto_increment not null ,
    username varchar(64),
    password varchar(64),
    status int(4),
    createtime int(10)
);`
	ModifyDB(sql)
}

// CreateTableWithArticle 创建文章表
func CreateTableWithArticle() {
	sql := `create table if not exists article(
    id int(4) primary key auto_increment not null ,
    title varchar(30),
    author varchar(20),
    tags varchar(30),
    short varchar(255),
    content longtext,
    createtime int(10)
);`
	ModifyDB(sql)
}

//--------图片--------
func CreateTableWithAlbum() {
	sql := `create table if not exists album(
    id int(4) primary key auto_increment not null ,
    filepath varchar(255),
    filename varchar(64),
    status int(4),
    createtime int(10)
);`
	ModifyDB(sql)
}

// QueryRowDB 查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

func SwitchTimeStampToData(timeStamp int64) string {
	t := time.Unix(timeStamp, 0)
	return t.Format("2006-01-02 15:04:05")
}
