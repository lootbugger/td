package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithWriter(app.infoLog.Writer()))

	router.POST("/tasks", app.addTask)
	router.GET("/tasks", app.listTasks)

	return router
}
