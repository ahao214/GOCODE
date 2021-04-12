package cache

import (
	"ahao.bbs/server/repositories"
	"errors"
	"github.com/goburrow/cache"
	"github.com/mlogclub/simple"
	"time"
)

type articleTagCache struct {
	cache cache.LoadingCache
}

var ArticleTagCache = newArticleTagCache()

func newArticleTagCache() *articleTagCache {
	return &articleTagCache{
		cache: cache.NewLoadingCache(
			func(key cache.Key) (value cache.Value, e error) {
				articleTags := repositories.ArticleTagRepository.FindByArticleId(simple.DB(), key2Int64(key))
				if len(articleTags) > 0 {
					var tagIds []int64
					for _, articleTag := range articleTags {
						tagIds = append(tagIds, articleTag.TagId)
					}
					value = tagIds
				} else {
					e = errors.New("文章没标签")
				}
				return
			},
			cache.WithMaximumSize(1000),
			cache.WithExpireAfterAccess(30*time.Minute),
		),
	}
}

func (c *articleCache) Get(articleId int64) []int64 {
	val, err := c.cache.Get(articleId)
	if err != nil {
		return nil
	}
	if val != nil {
		return val.([]int64)
	}
	return nil
}

func (c *articleCache) Invalidate(articleId int64) {
	c.cache.Invalidate(articleId)
}

func main() {

}
