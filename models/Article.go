package models

import (
	"../utils"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

/**
 * 文章表的增删改查数据模型
 */

// 文章表的映射模型
type Xinfei_article struct {
	Article_id       string `orm:"column(article_id);pk"` //文章id
	Article_title    string //文章标题
	Article_user     string //文章发布作者
	Article_html     string // 文章内容
	Article_text     string // 文章内容
	Article_category string //文章类别
	Article_date     string //文章发布时间
	Article_display  int64  //文章是否显示
	Article_hot      int64  //文章热度
}

// 插入一篇文章
func ArticleInsertOne(artivle Xinfei_article) error {
	o := orm.NewOrm()
	artivle.Article_id = utils.Random(6) //生成文章id
	artivle.Article_hot = 1              //文章默认热度为 1
	artivle.Article_display = 0
	artivle.Article_date = time.Now().Format("2006-01-02") //添加文章的日期时间
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

// 删除一篇文章 (完成)
func ArticleDeleteOne(id string) (int64, error) {
	o := orm.NewOrm()

	if num, err := o.Delete(&Xinfei_article{Article_id: id}); err == nil {
		return num, err
	} else {
		return 0, err
	}
}

//// 上下线一篇文章 (完成)
//func ArticleActivityOne(id string) (int64, error) {
//	o := orm.NewOrm()
//	article := Xinfei_article{Article_id: id}
//	if err := o.Read(&article); err == nil {
//		if article.Article_display == 0 {
//			article.Article_display = 1
//			if num, err := o.Update(&article); err == nil {
//				return num, err
//			} else {
//				return 0, err
//			}
//		} else {
//			article.Article_display = 0
//			if num, err := o.Update(&article); err == nil {
//				return num, err
//			} else {
//				return 0, err
//			}
//		}
//	} else {
//		return 0, err
//	}
//}

// 查询一篇文章详情内容
func ArticleSeleteOne(id string) (Xinfei_article, error) {
	o := orm.NewOrm()
	// 文章热度 + 1
	article := Xinfei_article{Article_id: id}
	o.Raw("update xinfei_article set article_hot = article_hot + 1 where article_id = ?", id).Exec()
	if err := o.Read(&article); err == nil {
		return article, err
	} else {
		o.Raw("update xinfei_article set article_hot = article_hot - 1 where article_id = ?", id).Exec()
		return article, err
	}
}

/**
 * 传入参数：是否上架的（空为全部查询），第几页，一页几行，升降序关键词
 * 查询全部文章
 * return 查询的数据，受影响行数，错误
 */
func ArticleSeleteAll(p int, num int, adc string, hot bool) ([]Xinfei_article, int64, error) {
	o := orm.NewOrm()
	article := []Xinfei_article{}
	sql := ""
	if hot {
		if p == 1 {
			sql = "SELECT * FROM xinfei_article" + " ORDER BY article_hot " + adc + " limit 0" + "," + strconv.Itoa(num)
		} else if p > 1 {
			s := p*num - num
			sql = "SELECT * FROM xinfei_article" + " ORDER BY article_hot " + adc + " limit " + strconv.Itoa(s) + "," + strconv.Itoa(num)
		} else {
			sql = "SELECT * FROM xinfei_article" + " ORDER BY article_hot " + adc + " limit 0" + "," + strconv.Itoa(num)
		}
	} else {
		if p == 1 {
			sql = "SELECT * FROM xinfei_article" + " ORDER BY article_date " + adc + " limit 0" + "," + strconv.Itoa(num)
		} else if p > 1 {
			s := p*num - num
			sql = "SELECT * FROM xinfei_article" + " ORDER BY article_date " + adc + " limit " + strconv.Itoa(s) + "," + strconv.Itoa(num)
		} else {
			sql = "SELECT * FROM xinfei_article" + " ORDER BY article_date " + adc + " limit 0" + "," + strconv.Itoa(num)
		}
	}
	// 返回行数或者错误信息
	if num, err := o.Raw(sql).QueryRows(&article); err == nil {
		for key, _ := range article {
			article[key].Article_html = ""
		}
		if p == 0 {
			return article, num, err
		} else {
			cnt, _ := o.QueryTable("xinfei_article").Count()
			return article, cnt, err
		}
	} else {
		return article, 0, err
	}
}

//获取文章表中的文章数量
func GetArticleCount() int64 {
	o := orm.NewOrm()
	cnt, _ := o.QueryTable("xinfei_article").Count()
	return cnt
}
