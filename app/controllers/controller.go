package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewIndex(db *sql.DB) *HomeController {
	return &HomeController{
		dba: db,
	}
}

type User struct {
	Id        int32
	Loginname string
	Password  string
}

type HomeController struct {
	dba *sql.DB
}

func (ctl *HomeController) Index(c *gin.Context) {
	var user User
	row := ctl.dba.QueryRow("select id, loginname, password from user where id=?", 1)
	err := row.Scan(&user.Id, &user.Loginname, &user.Password)
	if err != nil {
		panic(err.Error())
	}
	c.HTML(http.StatusOK, "index.html", user)

}
