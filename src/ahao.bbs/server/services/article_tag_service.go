package services

import (
	"ahao.bbs/server/model"
	"ahao.bbs/server/model/constants"
	"ahao.bbs/server/repositories"
	"github.com/mlogclub/simple"
)

var ArticleTagService = newArticleTagService()

func newArticleTagService() *articleTagService {
	return &articleTagService{}
}

type articleTagService struct {
}

func (s *articleTagService) Get(id int64) *model.ArticleTag {
	return repositories.ArticleTagRepository.Get(simple.DB(), id)
}

func (s *articleTagService) Take(where ...interface{}) *model.ArticleTag {
	return repositories.ArticleTagRepository.Take(simple.DB(), where...)
}

func (s *articleTagService) Find(cnd *simple.SqlCnd) []model.ArticleTag {
	return repositories.ArticleTagRepository.Find(simple.DB(), cnd)
}

func (s *articleTagService) FindPageByParams(params *simple.QueryParams) (list []model.ArticleTag, paging *simple.Paging) {
	return repositories.ArticleTagRepository.FindPageByParams(simple.DB(), params)
}

func (s *articleTagService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.ArticleTag, paging *simple.Paging) {
	return repositories.ArticleTagRepository.FindPageByCnd(simple.DB(), cnd)
}

func (s *articleTagService) Create(t *model.ArticleTag) error {
	return repositories.ArticleTagRepository.Create(simple.DB(), t)
}

func (s *articleTagService) Update(t *model.ArticleTag) error {
	return repositories.ArticleTagRepository.Update(simple.DB(), t)
}

func (s *articleTagService) Updates(id int64, columns map[string]interface{}) error {
	return repositories.ArticleTagRepository.Updates(simple.DB(), id, columns)
}

func (s *articleTagService) UpdateColumn(id int64, name string, value interface{}) error {
	return repositories.ArticleTagRepository.UpdateColumn(simple.DB(), id, name, value)
}

func (s *articleTagService) DeleteByArticleId(topicId int64) {
	simple.DB().Model(model.ArticleTag{}).Where("topic_id = ?", topicId).UpdateColumn("status", constants.StatusDeleted)
}
