package models

import (
	"../utils"
	"github.com/astaxie/beego/orm"
	"time"
)

//操作文章表的数据模型

// 文章表的映射模型
type Xinfei_article struct {
	Article_id       string `orm:"column(article_id);pk"` //文章id
	Article_title    string //文章标题
	Article_user     string //文章发布作者
	Article_text     string // 文章内容
	Article_category string //文章类别
	Article_date     string //文章发布时间
	Article_hot      int64  //文章热度
	Article_display  int64  //文章是否上线。 0或者1 -》下线/上线
}

// 插入一篇文章
func ArticleInsertOne(artivle Xinfei_article) error {
	o := orm.NewOrm()
	artivle.Article_id = utils.Random(6)
	artivle.Article_hot = 1
	artivle.Article_display = 1
	artivle.Article_date = time.Now().Format("2006-01-02")
	if _, err := o.Insert(&artivle); err == nil {
		return err
	} else {
		return err
	}
}

// 修改一篇文章
func ArticleUpdateOne(artivle Xinfei_article) (int64, error) {
	o := orm.NewOrm()
	artle := Xinfei_article{Article_id: artivle.Article_id}
	if err := o.Read(&artle); err == nil {
		if num, err := o.Update(&artivle); err == nil {
			return 0, err
		} else {
			return num, err
		}
	} else {
		return 0, err
	}
}

// 删除一篇文章
func ArticleDeleteOne(id string) (int64, error) {
	o := orm.NewOrm()
	if num, err := o.Delete(&Xinfei_article{Article_id: id}); err == nil {
		return num, err
	} else {
		return 0, err
	}
}

// 上下线一篇文章
func ArticleActivityOne(id string) (int64, error) {
	o := orm.NewOrm()
	article := Xinfei_article{Article_id: id}
	if err := o.Read(&article); err == nil {
		if article.Article_display == 0 {
			article.Article_display = 1
			if num, err := o.Update(&article); err == nil {
				return num, err
			} else {
				return 0, err
			}
		} else {
			article.Article_display = 0
			if num, err := o.Update(&article); err == nil {
				return num, err
			} else {
				return 0, err
			}
		}
	} else {
		return 0, err
	}
}

// 查询一篇文章
func ArticleSeleteOne(id string) (Xinfei_article, error) {
	o := orm.NewOrm()
	article := Xinfei_article{Article_id: id}
	if err := o.Read(&article); err == nil {
		// 文章热度 + 1
		o.Raw("update xinfei_article set article_hot = article_hot + 1 where article_id = ?", id).Exec()
		return article, err
	} else {
		return article, err
	}
}

// 查询全部文章
func ArticleSeleteAll() ([]Xinfei_article, error) {
	o := orm.NewOrm()
	article := []Xinfei_article{}
	_, err := o.Raw("SELECT * FROM xinfei_article").QueryRows(&article)
	if err == nil {
		return article, err
	} else {
		return article, err
	}
}
