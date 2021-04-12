package repositories

import (
	"ahao.bbs/server/model"
	"github.com/mlogclub/simple"
	"gorm.io/gorm"
)

var CommentRepository = newCommentRepository()

type commentRepository struct {
}

func newCommentRepository() *commentRepository {
	return &commentRepository{}
}

func (r *commentRepository) Get(db *gorm.DB, id int64) *model.Comment {
	res := &model.Comment{}
	if err := db.First(res, "id = ?", id).Error; err != nil {
		return nil
	}
	return res
}

func (r *commentRepository) Take(db *gorm.DB, where ...interface{}) *model.Comment {
	res := &model.Comment{}
	if err := db.Take(res, where...).Error; err != nil {
		return nil
	}
	return res
}

func (r *commentRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Comment) {
	cnd.Find(db, &list)
	return
}

func (r *commentRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.Comment {
	res := &model.Comment{}
	if err := cnd.FindOne(db, &res); err != nil {
		return nil
	}
	return res
}

func (r *commentRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.Comment, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *commentRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Comment, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Comment{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *commentRepository) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.Comment{})
}

func (r *commentRepository) Create(db *gorm.DB, t *model.Comment) (err error) {
	err = db.Create(t).Error
	return
}

func (r *commentRepository) Update(db *gorm.DB, t *model.Comment) (err error) {
	err = db.Save(t).Error
	return
}

func (r *commentRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Comment{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *commentRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Comment{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *commentRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.Comment{}, "id = ?", id)
}
