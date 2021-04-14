package repositories

import (
	"ahao.bbs/server/model"
	"github.com/mlogclub/simple"
	"gorm.io/gorm"
)

type userRepository struct {
}

var UserRepository = newUserRepository()

func newUserRepository() *userRepository {
	return &userRepository{}
}

func (r *userRepository) Get(db *gorm.DB, id int64) *model.User {
	res := &model.User{}
	if err := db.First(res, "id = ?", id).Error; err != nil {
		return nil
	}
	return res
}

func (r *userRepository) Take(db *gorm.DB, where ...interface{}) *model.User {
	res := &model.User{}
	if err := db.Take(res, where...).Error; err != nil {
		return nil
	}
	return res
}

func (r *userRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.User) {
	cnd.Find(db, &list)
	return
}

func (r *userRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.User {
	res := &model.User{}
	if err := cnd.FindOne(db, &res); err != nil {
		return nil
	}
	return res
}

func (r *userRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.User, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.User{})
	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *userRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.User, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *userRepository) Create(db *gorm.DB, t *model.User) (err error) {
	err = db.Create(t).Error
	return
}

func (r *userRepository) Update(db *gorm.DB, t *model.User) (err error) {
	err = db.Save(t).Error
	return
}

func (r *userRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.User{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *userRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.User{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *userRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.User{}, "id = ?", id)
}

func (r *userRepository) GetByEmail(db *gorm.DB, email string) *model.User {
	return r.Take(db, "email = ?", email)
}

func (r *userRepository) GetByUsername(db *gorm.DB, username string) *model.User {
	return r.Take(db, "username = ?", username)
}
