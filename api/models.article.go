package main

import (
	"errors"
)

type article struct {
  Id      		string
  Title   		string `sql:",notnull"`
  Content 		string `sql:",notnull"`
	Description string `sql:",notnull"`
	Author 			string `sql:",notnull"`
	Url					string `sql:",notnull"`
	UrlToImage	string `sql:",notnull"`
	PublishedAt	string `sql:",notnull"`
}

func getAllArticles() []article {
  var articles []article
  db := pgConnect()
  defer db.Close()
  err := db.Model(&articles).Select()
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