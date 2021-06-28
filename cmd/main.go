package main

import (
	"github.com/JIeeiroSst/test/config"
	"github.com/JIeeiroSst/test/delivery/http"
	repo "github.com/JIeeiroSst/test/repository/mysql"
	"github.com/labstack/echo/v4"
)

func main(){
	e:=echo.New()
	repo:=repo.NewMysqlConn(&config.Config.MysqlConfig)
	http.NewHandler(e,repo)
	e.Logger.Fatal(e.Start(":1234"))
}