package main

import (
  "os"
  "time"
  
  "github.com/go-pg/pg"
)

func pgOptions() *pg.Options {
  return &pg.Options{
    User:             os.Getenv("POSTGRES_USER"),
    Password:         os.Getenv("POSTGRES_PASSWORD"),
    Database:         os.Getenv("POSTGRES_DB"),
    Addr:     "db:" + os.Getenv("POSTGRES_PORT"),

    MaxRetries:      1,
    MinRetryBackoff: -1,

    DialTimeout:  30 * time.Second,
    ReadTimeout:  10 * time.Second,
    WriteTimeout: 10 * time.Second,

    PoolSize:           10,
    MaxConnAge:         10 * time.Second,
    PoolTimeout:        30 * time.Second,
    IdleTimeout:        10 * time.Second,
    IdleCheckFrequency: 100 * time.Millisecond,
  }
}

func pgConnect() *pg.DB {
  return pg.Connect(pgOptions())
}
