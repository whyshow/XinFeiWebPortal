package models

type AppBar struct {
	Appbar []NUC `json:"appbar"`
}
type NUC struct {
	Name     string     `json:"name"`
	Url      string     `json:"url"`
	Children []Children `json:"children"`
}

type Children struct {
	Cname string `json:"cname"`
	Curl  string `json:"curl"`
}
