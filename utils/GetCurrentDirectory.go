package utils

import (
	"github.com/astaxie/beego"
	"os"
	"path/filepath"
	"strings"
)

/*
获取程序运行路径
*/
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		beego.Debug(err)
	}
	return strings.Replace(dir, "\\", "/", -1) + "/"
}
