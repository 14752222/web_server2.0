package main

import (
	"fmt"
	"web_server_2.0/api/router"
	"web_server_2.0/bootstrap"
	data "web_server_2.0/client"
)

func main() {
	app := bootstrap.NewApplication()
	app.Mysql = data.NewMysql(app.Env)
	app.Redis = data.NewRedisClient(app.Env)
	app.Minio = data.GetfileManagerInstance()
	router.Setup(app)
	//app.Web.Run(fmt.Sprintf(":%s", app.Env.Server.Port))
	app.Web.Run(fmt.Sprintf(":%s", "8080"))

}
