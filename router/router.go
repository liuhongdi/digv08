package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv08/controller"
	"github.com/liuhongdi/digv08/pkg/result"
	"log"
	"runtime/debug"
)

func Router() *gin.Engine {
	router := gin.Default()
	//处理异常
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(Recover)


	// 路径映射
	articlec:=controller.NewArticleController()
	router.GET("/article/getone/:id", articlec.GetOne);
	router.GET("/article/list", articlec.GetList);
	return router
}

//404
func HandleNotFound(c *gin.Context) {
	result.NewResult(c).Error(404,"资源未找到")
	return
}

//500
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			result.NewResult(c).Error(500,"服务器内部错误")
		}
	}()
	//继续后续接口调用
	c.Next()
}