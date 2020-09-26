package db_mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_"github.com/go-sql-driver/mysql"
	"hl/models"
)
var DB*sql.DB
func init(){
	fmt.Println("链接数据库")
	config:=beego.AppConfig
	//appName:=config.String("appname")
	//fmt.Println("应用名称:",appName)
	//port,err:=config.Int("httpport")
	//if err!=nil{
	//	panic/*panic恐慌*/("项目配置文件解析失败，请检查配置文件")
	//}
	dbDriver:= config.String("db_driverName")
	dbUser:=config.String("db_user")
	dbPassword:=config.String("db_Password")
	dbIp:=config.String("db_Ip")
	dbName:=config.String("db_Name")
	fmt.Println(dbIp,dbName)
	connUrl:=dbUser+":"+dbPassword+ "@tcp("+dbIp+")/" + dbName + "?charset=utf8"
	//fmt.Println(connUrl)
	//fmt.Println(port)
	db,err :=sql.Open(dbDriver,connUrl)
	if err!=nil{
		panic("数据库连接错误，请检查配置")
}
	DB=db
}
func InserUser(user models.User)(int64,error){
	hashMd5:= md5.New()
	hashMd5.Write([]byte(user.Password))
    bytes:=hashMd5.Sum(nil)
    user.Password = hex.EncodeToString(bytes)
	fmt.Println("将要保存的用户名:",user.Nick,"密码:",user.Password)
	result,err:=DB.Exec("insert into user (user_name,birthday,address,nick,password) values(?,?,?,?,?)",user.Name,user.Birthday,user.Address,user.Nick, user.Password,)
	if err != nil {
		return 0,err
	}
	id,err:=result.RowsAffected()
	if err != nil {
		return 0,err
	}
	return id ,nil
}
