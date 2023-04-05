package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"todo/mypkg"
)

    func main() {
		// デフォルトのミドルウェアで新しい gin ルーターを作成する
		// logger とアプリケーションクラッシュをキャッチする recovery ミドルウェア
		router := gin.Default()
		router.LoadHTMLGlob("templates/*.html")
		mypkg.DbInit()

		//index
		router.GET("/", func(c *gin.Context){
			todos := mypkg.DbGetAll()
			c.HTML(http.StatusOK, "index.html",gin.H{
				"todos": todos,
			})
		})

		//Create
		router.POST("/new", func(c *gin.Context){
			text := c.PostForm("text")
			status := c.PostForm("status")
			mypkg.DbInsert(text, status)
			c.Redirect(302, "/")
		})
		router.Run()
	}