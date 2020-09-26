package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"hl/db_mysql"
	"hl/models"
	"io/ioutil"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController)Post(){
	var user models.User
	data,err:=ioutil.ReadAll(r.Ctx.Request.Body)
	if err !=nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	err = json.Unmarshal(data,&user)
	if err!=nil{
		fmt.Println(err.Error())
		r.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	id,err:=db_mysql.InserUser(user)
	if err !=nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("用户保存失败")
		return
	}
	fmt.Println(id)
	result := models.ResponseResult{
		Code: 0,
		Message:"保存成功",
		Data: nil,
	}
	r.Data["json"]=&result
	r.ServeJSON()
}
