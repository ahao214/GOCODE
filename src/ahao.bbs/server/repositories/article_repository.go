package repositories

import (
	model2 "Ahao.Blogger/model"
	"ahao.bbs/server/model"
	"github.com/mlogclub/simple"
	"gorm.io/gorm"
)

var ArticleRepository = newArticleRepository()

type articleRepository struct {
}

func newArticleRepository() *articleRepository {
	return &articleRepository{}
}

func (r *articleRepository) Get(db *gorm.DB, id int64) *model.Article {
	res := &model.Article{}
	if err := db.First(res, "id=?", id).Error; err != nil {
		return nil
	}
	return res
}

func (r *articleRepository) Take(db *gorm.DB, where ...interface{}) *model.Article {
	res := &model.Article{}
	if err := db.Take(res, where...).Error; err != nil {
		return nil
	}
	return res
}

func (r *articleRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Article) {
	cnd.Find(db, &list)
	return
}

func (r *articleRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.Article {
	res := &model.Article{}
	if err := cnd.FindOne(db, &res); err != nil {
		return nil
	}
	return res
}

func (r *articleRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.Article, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *articleRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Article, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Article{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *articleRepository) Create(db *gorm.DB, t *model.Article) (err error) {
	err = db.Create(t).Error
	return
}

func (r *articleRepository) Update(db *gorm.DB, t *model.Article) (err error) {
	err = db.Save(t).Error
	return
}

func (r *articleRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Article{}).Where("id=?", id).Update(columns).Error
	return
}

func (r *articleRepository) UpdateColumn(db *grom.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Article{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *articleRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.Article{}, "id=?", id)
}
