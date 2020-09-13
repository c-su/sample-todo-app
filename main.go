package main

import (
	"strconv"
	dbaccessor "work/app/db-accessor"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("app/templates/*.html")
	dbaccessor.DbInit(dbaccessor.DbOpen())

	router.GET("/hello", func(ctx *gin.Context) {
		todos := dbaccessor.DbGetAll(dbaccessor.DbOpen())
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	router.GET("/", func(ctx *gin.Context) {
		todos := dbaccessor.DbGetAll(dbaccessor.DbOpen())
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbaccessor.DbInsert(dbaccessor.DbOpen(), text, status)
		ctx.Redirect(302, "/")
	})

	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := dbaccessor.DbGetOne(dbaccessor.DbOpen(), id)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		dbaccessor.DbUpdate(dbaccessor.DbOpen(), id, text, status)
		ctx.Redirect(302, "/")
	})

	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := dbaccessor.DbGetOne(dbaccessor.DbOpen(), id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		dbaccessor.DbDelete(dbaccessor.DbOpen(), id)
		ctx.Redirect(302, "/")
	})

	router.Run()
}
