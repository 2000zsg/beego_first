package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	//"hl/db_mysql"
	"hl/models"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//1.获取请求数据
	user:=c.Ctx.Input.Query("user")//c. getstring("user")返回整数iiiiiii'j
	psd:= c.Ctx.Input.Query("psd")
	//2.使用固定数据进行数据校验
	if user !="zhangshenggang" || psd !="123456"{
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据校验不成功"))
		return
	}else {//校验正确
		c.Ctx.ResponseWriter.Write([]byte("校验成功"))
	}

	/*c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"*/
}
/*func (c *MainController) Post()  {
	name := c.Ctx.Request.FormValue("name")
	age:= c.Ctx.Request.FormValue("age")
	sex := c.Ctx.Request.FormValue("sex")
	fmt.Println(name,sex,age)
	if name !="zhangshenggang" || age !="18"{
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据校验不成功"))
		return
	}else {
		c.Ctx.WriteString("恭喜，数据校验成功")
	}
}*/
//该方法用于处理post类型的请求
//func (c *MainController) Post()  {
//	1.解析前端提交的json格式数据
	//var person  models.Person
	//dataByes,err := ioutil.ReadAll(c.Ctx.Request.Body)
	//if err!= nil{
	//	c.Ctx.WriteString("数据接收错误，请重试")
	//	return
	//}
	//err = json.Unmarshal(dataByes ,&person)
	//if err!= nil{
	//	c.Ctx.WriteString("数据接收失败，请重试")
	//	return
	//}
	//fmt.Println("姓名:",person.Name)
	//fmt.Println("姓名:",person.Age)
	//fmt.Println("姓名:",person.Sex)
	//c.Ctx.WriteString("数据解析成功")
//}
func (c *MainController) Post() {
	var air  models.Air
	dataBytes,err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err!= nil{
		c.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	err = json.Unmarshal(dataBytes ,&air)
	if err!= nil{
		c.Ctx.WriteString("数据接收失败，请重试")
		return
	}
	//db_mysql.InserUser(air)
	//fmt.Println(air.Name)
	//fmt.Println(air.Address)
	//fmt.Println(air.Birthday)
	//fmt.Println(air.Nick)
	//c.Ctx.WriteString("数据解析成功")
}