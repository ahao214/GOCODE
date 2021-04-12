package repositories

import (
	"ahao.bbs/server/model"
	"github.com/mlogclub/simple"
	"gorm.io/gorm"
)

var LinkRepository = newLinkRepository()

type linkRepository struct {
}

func newLinkRepository() *linkRepository {
	return &linkRepository{}
}

func (r *linkRepository) Get(db *gorm.DB, id int64) *model.Link {
	res := &model.Link{}
	if err := db.First(res, "id = ?", id).Error; err != nil {
		return nil
	}
	return res
}

func (r *linkRepository) Take(db *gorm.DB, where ...interface{}) *model.Link {
	res := &model.Link{}
	if err := db.Take(res, where...).Error; err != nil {
		return nil
	}
	return res
}

func (r *linkRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Link) {
	cnd.Find(db, &list)
	return
}

func (r *linkRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.Link {
	res := &model.Link{}
	if err := cnd.FindOne(db, &res); err != nil {
		return nil
	}
	return res
}

func (r *linkRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.Link, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *linkRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Link, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Link{})
	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *linkRepository) Create(db *gorm.DB, t *model.Link) (err error) {
	err = db.Create(t).Error
	return
}
func (r *linkRepository) Update(db *gorm.DB, t *model.Link) (err error) {
	err = db.Save(t).Error
	return
}

func (r *linkRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Link{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *linkRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Link{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *linkRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.Link{}, "id = ?", id)
}
