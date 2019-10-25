package main

import (
	"net/http"

  "github.com/gin-gonic/gin"
)

type queryParams struct {
	PageNumber	string
	PageSize		string
}

func showIndexPage(c *gin.Context) {
	number, _ := c.GetQuery("pageNumber")
	size, _ := c.GetQuery("pageSize")
	if number == "" {
		number = "1"
	}
	if size == "" {
		size = "25"
	}
	options := queryParams{
		PageNumber: number,
		PageSize: size,
	}
	articles := getAllArticles(options)
	contentType := c.Request.Header.Get("Content-Type")
	if contentType == "application/json" {
		c.JSON(http.StatusOK, gin.H{"data": articles})
	} else {
		c.HTML(http.StatusOK, "index.html", gin.H{ "title":   "Home Page", "payload": articles })
	}
}

func getArticle(c *gin.Context) {
	articleId := c.Param("article_id")
	if article, err := getArticleByID(articleId); err == nil {
		c.HTML(http.StatusOK, "article.html", gin.H{ "title":   article.Title, "payload": article })
	} else {
		c.AbortWithError(http.StatusNotFound, err)
	}
}
