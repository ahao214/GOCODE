package repositories

import (
	"ahao.bbs/server/model"
	"github.com/mlogclub/simple"
	"gorm.io/gorm"
)

var EmailCodeRepository = newEmailCodeRepository()

type emailCodeRepository struct {
}

func newEmailCodeRepository() *emailCodeRepository {
	return &emailCodeRepository{}
}

func (r *emailCodeRepository) Get(db *gorm.DB, id int64) *model.EmailCode {
	res := &model.EmailCode{}
	if err := db.First(res, "id = ?", id).Error; err != nil {
		return nil
	}
	return res
}

func (r *emailCodeRepository) Take(db *gorm.DB, where ...interface{}) *model.EmailCode {
	res := &model.EmailCode{}
	if err := db.Take(res, where...).Error; err != nil {
		return nil
	}
	return res
}

func (r *emailCodeRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.EmailCode) {
	cnd.Find(db, &list)
	return
}

func (r *emailCodeRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.EmailCode {
	res := &model.EmailCode{}
	if err := cnd.FindOne(db, &res); err != nil {
		return nil
	}
	return res
}

func (r *emailCodeRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.EmailCode, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *emailCodeRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.EmailCode, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.EmailCode{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *emailCodeRepository) Count(db *gorm.DB, cnd *simple.SqlCnd) int64 {
	return cnd.Count(db, &model.EmailCode{})
}

func (r *emailCodeRepository) Create(db *gorm.DB, t *model.EmailCode) (err error) {
	err = db.Create(t).Error
	return
}

func (r *emailCodeRepository) Update(db *gorm.DB, t *model.EmailCode) (err error) {
	err = db.Save(t).Error
	return
}

func (r *emailCodeRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.EmailCode{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *emailCodeRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.EmailCode{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *emailCodeRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.EmailCode{}, "id = ?", id)
}