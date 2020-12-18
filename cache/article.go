package cache

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/liuhongdi/digv08/global"
	"github.com/liuhongdi/digv08/model"
	"strconv"
	"time"
)
//token的过期时长
const ArticleDuration = time.Minute * 5

//cache的名字
func getArticleCacheName(articleId uint64) (string) {
	return "article_"+strconv.FormatUint(articleId,10)
}

//从cache得到一篇文章
func GetOneArticleCache(articleId uint64) (*model.Article,error) {
	key := getArticleCacheName(articleId);
	val, err := global.RedisDb.Get(key).Result()

	if (err == redis.Nil || err != nil) {
		return nil,err
	} else {
		article := model.Article{}
		if err := json.Unmarshal([]byte(val), &article); err != nil {
			//t.Error(target)
			return nil,err
		}
		return &article,nil
	}
}
//向cache保存一篇文章
func SetOneArticleCache(articleId uint64,article *model.Article) (error) {
	key := getArticleCacheName(articleId);
	content,err := json.Marshal(article)
	if (err != nil){
		fmt.Println(err)
		return err;
	}
	errSet := global.RedisDb.Set(key, content, ArticleDuration).Err()
	if (errSet != nil) {
		return errSet
	}
	return nil
}



