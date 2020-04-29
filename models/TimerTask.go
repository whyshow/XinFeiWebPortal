package models

import (
	"fmt"
	"github.com/astaxie/beego/toolbox"
)

//定时更新日访问量、日访客量
func TimerUpPUTask() {
	//注册每10秒执行一次的定时任务
	tk := toolbox.NewTask("myTask", "0/10 * * * * *  ", func() error { fmt.Println(GetPvUv().Pv); return nil })
	//注册 每60分钟执行一次的定时任务
	//tk := toolbox.NewTask("myTask", "0 */60 * * * *", func() error { fmt.Println("hello world"); return nil })
	//运行任务
	tk.Run()
	//加入全局的计划任务列表
	toolbox.AddTask("myTask", tk)
	//开始执行全局的任务
	toolbox.StartTask()
}
