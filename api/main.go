package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
  router = gin.Default()
  router.LoadHTMLGlob("templates/*")
	initializeRoutes()
	router.RunTLS(":443", "/etc/letsencrypt/live/newsapp.glennfmeyer.com/fullchain.pem", "/etc/letsencrypt/live/newsapp.glennfmeyer.com/privkey.pem")
}