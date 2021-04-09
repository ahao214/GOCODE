package repositories

import (
	"github.com/mlogclub/simple"
	"github.com/mlogclub/simple/date"
	"gorm.io/gorm"

	"ahao.bbs/server/model"
)

var ArticleTagRepository = newArticleTagRepository()

type articleTagRepository struct {
}

func newArticleTagRepository() *articleTagRepository {
	return &articleTagRepository{}
}

func (r *articleTagRepository) Get(db *gorm.DB, id int64) *model.ArticleTag {
	res := &model.ArticleTag{}
	if err := db.First(res, "id = ?", id).Error; err != nil {
		return nil
	}
	return res
}

func (r *articleTagRepository) Take(db *gorm.DB, where ...interface{}) *model.ArticleTag {
	res := &model.ArticleTag{}
	if err := db.Take(res, where...).Error; err != nil {
		return nil
	}
	return res
}

func (r *articleTagRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.ArticleTag) {
	cnd.Find(db, &list)
	return
}

func (r *articleTagRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.ArticleTag {
	res := &model.ArticleTag{}
	if err := cnd.FindOne(db, &res); err != nil {
		return nil
	}
	return res
}

func (r *articleTagRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.ArticleTag, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *articleTagRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.ArticleTag, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.ArticleTag{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *articleTagRepository) Create(db *gorm.DB, t *model.ArticleTag) (err error) {
	err = db.Create(t).Error
	return
}

func (r *articleTagRepository) Update(db *gorm.DB, t *model.ArticleTag) (err error) {
	err = db.Save(t).Error
	return
}

func (r *articleTagRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.ArticleTag{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *articleTagRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.ArticleTag{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *articleTagRepository) Delete(db *grom.DB, id int64) {
	db.Delete(&model.ArticleTag{}, "id = ?", id)
}

func (r *articleTagRepository) AddArticleTags(db *gorm.DB, articleId int64, tagIds []int64) {
	if articleId <= 0 || len(tagIds) == 0 {
		return
	}
	for _, tagId := range tagIds {
		_ = r.Create(db, &model.ArticleTag{
			ArticleId:  articleId,
			TagId:      tagId,
			CreateTime: date.NowTimestamp(),
		})
	}
}

func (r *articleTagRepository) DeleteArticleTags(db *gorm.DB, articleId int64) {
	if articleId <= 0 {
		return
	}
	db.Where("article_id = ?", articleId).Delete(model.ArticleTag{})
}

func (r *articleTagRepository) DeleteArticleTag(db *gorm.DB, articleId, tagId int64) {
	if articleId <= 0 {
		return
	}
	db.Where("article_id = ? and tag_id = ?", articleId, tagId).Delete(model.ArticleTag{})
}

func (r *articleTagRepository) FindByArticleId(db *gorm.DB, articleId int64) []model.ArticleTag {
	return r.Find(db, simple.NewSqlCnd().Where("article_id = ?", articleId))
}
