package repositories

import (
	"ahao.bbs/server/model"
	"github.com/mlogclub/simple"
	"gorm.io/gorm"
)

var UserScoreLogRepository = newUserScoreLogRepository()

type userScoreLogRepository struct {
}

func newUserScoreLogRepository() *userScoreLogRepository {
	return &userScoreLogRepository{}
}

func (r *userScoreLogRepository) Get(db *gorm.DB, id int64) *model.UserScoreLog {
	res := &model.UserScoreLog{}
	if err := db.First(res, "id = ?", id).Error; err != nil {
		return nil
	}
	return res
}

func (r *userScoreLogRepository) Take(db *gorm.DB, where ...interface{}) *model.UserScoreLog {
	res := &model.UserScoreLog{}
	if err := db.Take(res, where...).Error; err != nil {
		return nil
	}
	return res
}

func (r *userScoreLogRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserScoreLog) {
	cnd.Find(db, &list)
	return
}

func (r *userScoreLogRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.UserScoreLog {
	res := &model.UserScoreLog{}
	if err := cnd.FindOne(db, &res); err != nil {
		return nil
	}
	return res
}

func (r *userScoreLogRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.UserScoreLog, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.UserScoreLog{})
	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *userScoreLogRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.UserScoreLog, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *userScoreLogRepository) Create(db *gorm.DB, t *model.UserScoreLog) (err error) {
	err = db.Create(t).Error
	return
}

func (r *userScoreLogRepository) Update(db *gorm.DB, t *model.UserScoreLog) (err error) {
	err = db.Save(t).Error
	return
}

func (r *userScoreLogRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.UserScoreLog{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *userScoreLogRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.UserScoreLog{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *userScoreLogRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.UserScoreLog{}, "id = ?", id)
}
