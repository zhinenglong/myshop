package main

import (
	"myshop/app/controllers"
	"myshop/core"
	"net/http"
)

func main() {
	defer func() {
		if controllers.Dba != nil {
			controllers.Dba.Close()
		}
	}()
	http.ListenAndServe(":8080", core.GlobService)
}
