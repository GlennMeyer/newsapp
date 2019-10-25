package main

import (
	"errors"
	"fmt"
	"strconv"
)

type article struct {
  Id      		string `json:"id"`
  Title   		string `json:"title"`
  Content 		string `json:"content"`
	Description string `json:"description"`
	Author 			string `json:"author"`
	Url					string `json:"url"`
	UrlToImage	string `json:"urlToImage"`
	PublishedAt	string `json:"publishedAt"`
}

func getAllArticles(options queryParams) []article {
	var articles []article
	limit, _ := strconv.Atoi(options.PageSize)
	offset, _ := strconv.Atoi(options.PageNumber)
	offset = (offset - 1) * limit

	fmt.Println("Offset: ", offset)
	fmt.Println("Limit: ", limit)

  db := pgConnect()
  defer db.Close()
  err := db.Model(&articles).Offset(offset).Limit(limit).Select()
  if err != nil {
    panic(err)
  }
  return articles
}

func getArticleByID(id string) (*article, error) {
  article := &article{Id: id}
  db := pgConnect()
  defer db.Close()
  err := db.Select(article)
  if err != nil {
    return nil, errors.New("Article not found")
  }
  return article, nil
} 