package services

import (
	"ahao.bbs/server/cache"
	"ahao.bbs/server/model"
	"ahao.bbs/server/model/constants"
	"ahao.bbs/server/repositories"
	"strings"

	"github.com/mlogclub/simple"
)

type tagService struct {
}

var TagService = newTagService()

func newTagService() *tagService {
	return &tagService{}
}

func (s *tagService) Get(id int64) *model.Tag {
	return repositories.TagRepository.Get(simple.DB(), id)
}

func (s *tagService) Take(where ...interface{}) *model.Tag {
	return repositories.TagRepository.Take(simple.DB(), where...)
}

func (s *tagService) Find(cnd *simple.SqlCnd) []model.Tag {
	return repositories.TagRepository.Find(simple.DB(), cnd)
}

func (s *tagService) FindOne(cnd *simple.SqlCnd) *model.Tag {
	return repositories.TagRepository.FindOne(simple.DB(), cnd)
}

func (s *tagService) FindPageByParams(params *simple.QueryParams) (list []model.Tag, paging *simple.Paging) {
	return repositories.TagRepository.FindPageByParams(simple.DB(), params)
}

func (s *tagService) FindPageByCnd(cnd *simple.SqlCnd) (list []model.Tag, paging *simple.Paging) {
	return repositories.TagRepository.FindPageByCnd(simple.DB(), cnd)
}

func (s *tagService) Create(t *model.Tag) error {
	return repositories.TagRepository.Create(simple.DB(), t)
}

func (s *tagService) Update(t *model.Tag) error {
	if err := repositories.TagRepository.Update(simple.DB(), t); err != nil {
		return err
	}
	cache.TagCache.Invalidate(t.Id)
	return nil
}

func (s *tagService) GetByName(name string) *model.Tag {
	return repositories.TagRepository.GetByName(name)
}

func (s *tagService) GetTags() []model.TagResponse {
	list := repositories.TagRepository.Find(simple.DB(), simple.NewSqlCnd().Where("status = ?", constants.StatusOK))
	var tags []model.TagResponse
	for _, tag := range list {
		tags = append(tags, model.TagResponse{TagId: tag.Id, TagName: tag.Name})
	}
	return tags
}

func (s *tagService) GetTagInIds(tagIds []int64) []model.Tag {
	return repositories.TagRepository.GetTagInIds(tagIds)
}

//扫描
func (s *tagService) Scan(callback func(tags []model.Tag)) {
	var cursor int64
	for {
		list := repositories.TagRepository.Find(simple.DB(), simple.NewSqlCnd().Where("id > ?", cursor).Asc("id").Limit(100))
		if list == nil || len(list) == 0 {
			break
		}
		cursor = list[len(list)-1].Id
		callback(list)

	}
}
