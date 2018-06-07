package main

import "github.com/astaxie/beego"

type ImageController struct {
	beego.Controller
}

var fileKey = "testFile"
var filePath = "/tmp/"

func (this *ImageController) Upload() {
	name := this.Input().Get("name");

	beego.Info("Name ", name)

	file, header, fileErr := this.GetFile(fileKey)

	if file != nil {
		filename := header.Filename
		beego.Info("Uploading file ", filename)
		beego.Info("File size ", header.Size)

		saveErr := this.SaveToFile(fileKey, filePath + filename)
		if saveErr != nil {
			beego.Error(saveErr)
		} else {
			this.Ctx.WriteString(filename + " has been uploaded!")
		}
	} else {
		beego.Error(fileErr)
	}
}
