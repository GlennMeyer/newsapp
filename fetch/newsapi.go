package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
	"strings"
	"time"
)

type Source struct {
  Id          string
  Name        string
  Description string	
  Url         string
  Category    string
  Language    string
  Country     string
}
type Sources struct {
  Status  string
  Sources []Source
}
type MiniSource struct {
	id		string
	name 	string
}
type Article struct {
  Author        string          
  Title         string
  Description   string
  Url           string
  UrlToImage    string
  PublishedAt   string
	Content       string
	Source				MiniSource
}
type Articles struct {
  Status        string
  TotalResults  int
  Articles      []Article
}

const api = "https://newsapi.org/v2/"

var twentySources = []string{
  "cnn",
  "the-new-york-times",
  "the-huffington-post",
  "fox-news",
  "usa-today",
  "reuters",
  "politico",
  "breitbart-news",
  "nbc-news",
  "cbs-news",
  "abc-news",
  "newsweek",
  "the-washington-post",
  "google-news",
  "the-wall-street-journal",
  "associated-press",
}

func createTables() {
	db := pgConnect()
  defer db.Close()
  _, err := db.Exec(`CREATE TABLE IF NOT EXISTS sources (
                      id varchar PRIMARY KEY,
                      name varchar,
                      description varchar,
                      url varchar,
                      category varchar,
                      language varchar,
                      country varchar
										);`)
	if err != nil {
		fmt.Println("Error connecting to database, trying again in one minute. Error: ", err)
		time.Sleep(1 * time.Minute)
		createTables()
	}
	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
										CREATE TABLE IF NOT EXISTS articles (
                      id uuid NOT NULL DEFAULT uuid_generate_v1(),
                      source json,
                      author varchar,
                      title varchar,
                      description varchar,
                      url varchar UNIQUE,
                      url_to_image varchar,
                      published_at timestamp,
                      content text
                    );`)
  if err != nil {
    panic(err)
  }
}

func fetchSources() {
  var response Sources
	resp, err := http.Get(api + "sources?apiKey=" + os.Getenv("NEWSAPI_KEY"))
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  json.Unmarshal([]byte(body), &response)
  db := pgConnect()
  defer db.Close()
  if response.Status == "ok" {
    for _, source := range response.Sources {
      _, err := db.Model(&source).OnConflict("(id) DO UPDATE").Insert()
      if err != nil {
        panic(err)
      }
    }
  }
}

func getNews() {
  var response Articles
  everything := api + "everything?sources=" +
								strings.Join(twentySources, ",") +
								"&apiKey=" + os.Getenv("NEWSAPI_KEY") +
								"&language=en" +
								"&pageSize=100"
  fmt.Println(everything)
  resp, err := http.Get(everything)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic(err)
  }
  json.Unmarshal([]byte(body), &response)
  if response.Status =="ok" {
		fmt.Println("Number of articles: ", response.TotalResults)
    db := pgConnect()
    defer db.Close()
    for _, article := range response.Articles {
      _, err := db.Model(&article).OnConflict("(url) DO NOTHING").Insert()
      if err != nil {
        panic(err)
      }
    }
  }
}

func sleepFetch() {
	fmt.Println("fetching and sleeping")
	getNews()
	time.Sleep(3 * time.Minute)
	sleepFetch()
}
