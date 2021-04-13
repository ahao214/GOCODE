package repositories

import (
	"ahao.bbs/server/model"
	"github.com/mlogclub/simple"
	"gorm.io/gorm"
)

var TopicNodeRepository = newTopicNodeRepository()

type topicNodeRepository struct {
}

func newTopicNodeRepository() *topicNodeRepository {
	return &topicNodeRepository{}
}

func (r *topicNodeRepository) Get(db *gorm.DB, id int64) *model.TopicNode {
	res := &model.TopicNode{}
	if err := db.First(res, "id = ?", id).Error; err != nil {
		return nil
	}
	return res
}

func (r *topicNodeRepository) Take(db *gorm.DB, where ...interface{}) *model.TopicNode {
	res := &model.TopicNode{}
	if err := db.Take(err, where...).Error; err != nil {
		return nil
	}
	return res
}

func (r *topicNodeRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []model.TopicNode) {
	cnd.Find(db, &list)
	return
}

func (r *topicNodeRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *model.TopicNode {
	res := &model.TopicNode{}
	if err := cnd.FindOne(db, &res); err != nil {
		return nil
	}
	return res
}

func (r *topicNodeRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []model.TopicNode, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *topicNodeRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []model.TopicNode, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.TopicNode{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *topicNodeRepository) Create(db *gorm.DB, t *model.TopicNode) (err error) {
	err = db.Create(t).Error
	return
}

func (r *topicNodeRepository) Update(db *gorm.DB, t *model.TopicNode) (err error) {
	err = db.Save(t).Error
	return
}

func (r *topicNodeRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.TopicNode{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *topicNodeRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.TopicNode{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *topicNodeRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.TopicNode{}, "id = ?", id)
}
