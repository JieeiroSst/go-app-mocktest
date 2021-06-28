package config

import (
	"fmt"

	"github.com/JIeeiroSst/test/repository/mysql"
	"github.com/kelseyhightower/envconfig"
)

type WebConfig struct {
	PORT string 	`envconfig:"WEB_PORT"`
	MysqlConfig mysql.Config `envconfig:"WEB_MYSQL"`
}

var Config WebConfig

func init(){
	err:=envconfig.Process("",&Config)
	if err!=nil{
		panic(err)
	}
	fmt.Println(Config)
}
