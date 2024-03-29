package repositories

import (
	"ahao.bbs/server/model"
	"github.com/mlogclub/simple"
	"gorm.io/gorm"
)

var MessageRepository = newMessageRepository()

type messageRepository struct {
}

func newMessageRepository() *messageRepository {
	return &messageRepository{}
}

func (r *messageRepository) Get(db *gorm.DB, id int64) *model.Message {
	res := &model.Message{}
	if err := db.First(res, "id = ?", id).Error; err != nil {
		return nil
	}
	return res
}

func (r *messageRepository) Take(db *gorm.DB, where ...interface{}) *model.Message {
	res := &model.Message{}
	if err := db.Take(res, where...).Error; err != nil {
		return nil
	}
	return res
}

func (r *messageRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Message) {
	cnd.Find(db, &list)
	return
}

func (r *messageRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.Message {
	res := &model.Message{}
	if err := cnd.FindOne(db, &res); err != nil {
		return nil
	}
	return res
}

func (r *messageRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.Message, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *messageRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.Message, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.Message{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *messageRepository) Create(db *gorm.DB, t *model.Message) (err error) {
	err = db.Create(t).Error
	return
}

func (r *messageRepository) Update(db *gorm.DB, t *model.Message) (err error) {
	err = db.Save(t).Error
	return
}

func (r *messageRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.Message{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *messageRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.Message{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *messageRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.Message{}, "id = ?", id)
}
