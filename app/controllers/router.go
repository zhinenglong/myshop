package controllers

import (
	"database/sql"
	"myshop/core"
)
import _ "github.com/go-sql-driver/mysql"

var Dba *sql.DB

func init() {
	Dba := makeDb("mysql", "root:123456@tcp(127.0.0.1:3306)/elishop?charset=utf8mb4&parseTime=True&loc=Local")
	core.GlobService.GET("/", NewIndex(Dba).Index)
}

func makeDb(dbtype string, dsn string) *sql.DB {
	db, err := sql.Open(dbtype, dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
