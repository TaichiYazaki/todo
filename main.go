package main

import (
	"net/http"
	"strconv"
	"todo/mypkg"

	"github.com/gin-gonic/gin"
)

func main() {
	// デフォルトのミドルウェアで新しい gin ルーターを作成する
	// logger とアプリケーションクラッシュをキャッチする recovery ミドルウェア
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	mypkg.DbInit()

	//index
	router.GET("/", func(cxt *gin.Context) {
		todos := mypkg.DbGetAll()
		cxt.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})

	//Create
	router.POST("/new", func(cxt *gin.Context) {
		text := cxt.PostForm("text")
		status := cxt.PostForm("status")
		mypkg.DbInsert(text, status)
		cxt.Redirect(302, "/")
	})

	//Detail
	router.GET("/detail/:id", func(cxt *gin.Context) {
		n := cxt.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := mypkg.DbGetOne(id)
		cxt.HTML(http.StatusOK, "detail.html", gin.H{
			"todo": todo,
		})
	})

	//Update
	router.POST("/update/:id", func(cxt *gin.Context) {
		n := cxt.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := cxt.PostForm("text")
		status := cxt.PostForm("status")
		mypkg.DbUpdate(id, text, status)
		cxt.Redirect(302, "/")

	})

	//Delete
	router.GET("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		mypkg.DbDelete(id)
		ctx.Redirect(302, "/")
	})

	router.Run()
}
