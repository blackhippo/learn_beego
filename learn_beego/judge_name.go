package main

import "github.com/astaxie/beego"
import "math/rand"

type JudgeController struct {
	beego.Controller
}

type Person struct {
	IQ int
	LookScore int
	RealAge int
}

var WEI_NAME = "过伟"
var QI_NAME = "罗琦"

var qiIQ = 50
var weiIQ = 140
var otherIQ = 90

var qiLook = 40
var weiLook = 80
var otherLook = 60

var qiAge = 40
var weiAge = 18
var otherAge = 30

func (this *JudgeController) Judge() {
	baseIQ := otherIQ
	baseLook := otherLook
	baseAge := otherAge

	name := this.Input().Get("name");

	switch name {
	case WEI_NAME:
		baseIQ = weiIQ
		baseLook = weiLook
		baseAge = weiAge
	case QI_NAME:
		baseIQ = qiIQ
		baseLook = qiLook
		baseAge = qiAge
	default:
	}

	result := Person{baseIQ + rand.Intn(10),
		baseLook + rand.Intn(10),
		baseAge + rand.Intn(10)}

	this.Data["json"] = result
	this.ServeJSON()
}

