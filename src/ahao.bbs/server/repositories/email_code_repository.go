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
