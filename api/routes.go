package main

import (
)

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticle)
	v1 := router.Group("/api")
	{
		articles := v1.Group("/articles")
		{
			articles.GET("/", showIndexPage)
		}
	}
}