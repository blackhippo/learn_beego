package main

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

type User struct {
	Name string
	Age string
}

func (this *MainController) Display() {
	name := this.Input().Get("name")
	age := this.Input().Get("age")

	user := User{name, age}

	beego.Info("Created User:", user)

	this.Data["json"] = user
	this.ServeJSON()
}

func main() {
	beego.Router("/display", &MainController{}, "get:Display")
	beego.Run()
}