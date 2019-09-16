package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
  articles := getAllArticles()
  c.HTML(http.StatusOK, "index.html", gin.H{ "title":   "Home Page", "payload": articles })
}

func getArticle(c *gin.Context) {
	articleId := c.Param("article_id")
	if article, err := getArticleByID(articleId); err == nil {
		c.HTML(http.StatusOK, "article.html", gin.H{ "title":   article.Title, "payload": article })
	} else {
		c.AbortWithError(http.StatusNotFound, err)
	}
}