package controllers

import (
	"../utils"
	"github.com/astaxie/beego"
	"image"
	"time"
)

type UploadingController struct {
	beego.Controller
}

// 上传图片
func (c *UploadingController) UploadingImage() {

	f, h, err := c.GetFile("uploadimage")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"errno": 1, "data": ""}
		c.ServeJSON()
		return
	}
	defer f.Close()
	timeUnix := time.Now().Unix() //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02")
	//判断文件夹是否存在，如果不存在就创建
	utils.IsDir("static/upload/" + formatTimeStr + "/")
	c.SaveToFile("uploadimage", "static/upload/"+formatTimeStr+"/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建

	//压缩图片
	config, _, _ := image.DecodeConfig(f)
	// 根据图片宽高来裁剪图片
	// 默认的横图纵横比为 w 4 : h 3 即 640＊480
	// 默认的竖图纵横比为 w 0.618 : h 1 即 475 : 768
	if config.Width/config.Height > 2 {
		utils.Thumbnail("static/upload/"+formatTimeStr+"/"+h.Filename, uint(config.Width), uint(config.Height))
	} else if config.Width < 640 && config.Height < 480 && config.Width > config.Height {

		utils.Thumbnail("static/upload/"+formatTimeStr+"/"+h.Filename, uint(config.Width), uint(config.Height))
	} else if config.Width < 475 && config.Height < 768 && config.Width < config.Height {
		utils.Thumbnail("static/upload/"+formatTimeStr+"/"+h.Filename, uint(config.Width), uint(config.Height))
	} else if config.Width > config.Height {
		// 认为图片为 4 : 3
		utils.Thumbnail("static/upload/"+formatTimeStr+"/"+h.Filename, 640, 480)
	} else if config.Width < config.Height {
		// 认为图片为 0.618 : 1
		utils.Thumbnail("static/upload/"+formatTimeStr+"/"+h.Filename, 475, 768)
	} else {
		utils.Thumbnail("static/upload/"+formatTimeStr+"/"+h.Filename, uint(config.Width), uint(config.Height))
	}
	var data = []string{}
	data = append(data, "http://image.ccit.club"+"/static/upload/"+formatTimeStr+"/"+h.Filename)
	c.Data["json"] = map[string]interface{}{"errno": 0, "data": data}
	c.ServeJSON()
}

// 上传自拍照
func (c *UploadingController) UploadingSelfie() {
	//接收
	f, h, err := c.GetFile("file")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 1, "msg": "系统接收不到图片哎!", "data": nil}
		c.ServeJSON()
		return
	}
	defer f.Close()
	timeUnix := time.Now().Unix() //已知的时间戳
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02")
	fname := utils.Random(4)
	//判断文件夹是否存在，如果不存在就创建
	utils.IsDir("static/selfie/upload/" + formatTimeStr + "/")
	c.SaveToFile("file", "static/selfie/upload/"+formatTimeStr+"/"+fname+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
	//压缩图片
	utils.Thumbnail("static/selfie/upload/"+formatTimeStr+"/"+fname+h.Filename, 390, 550)
	msg := "http://image.ccit.club/static/selfie/upload/" + formatTimeStr + "/" + fname + h.Filename
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": msg, "data": nil}
	c.ServeJSON()
}
