package models

import "time"

/**
 * 日访问量、日访客 数据模型
 */
type PvUv struct {
	Pv int
	Uv int
}

var pv, uv int = 0, 0

func AddPv(p int) {
	pv += p
}

func AddUv(u int) {
	uv += u
}

//获取日访问量、访客量
func GetPvUv() PvUv {
	return PvUv{Pv: pv / 2, Uv: uv / 2}
}

type Xinfei_Day struct {
	Date   time.Time `orm:"column(date);pk"` //日期
	Pv_day int       //日访问量
	Uv_day int       //日访客量
}

//插入今天的日访问量、日访客量
func InsertPU(pv int, uv int) {
	//o := orm.NewOrm()

}
