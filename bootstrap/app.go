package bootstrap

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"web_server_2.0/client"
	"web_server_2.0/config"
)

type Application struct {
	Env   *config.Env
	Mysql *gorm.DB
	Redis *redis.Client
	//Mongo  *Mongo
	//Logger *Logger
	Web   *gin.Engine
	Minio *client.FileManager
}

func NewApplication() *Application {
	app := &Application{}
	app.Env = config.NewEnv()
	app.Run()
	return app
}

func (app *Application) CloseDBConnection() {
	sqlDB, err := app.Mysql.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDB.Close()
	if err != nil {
		return
	}
}

func (app *Application) Run() {
	r := gin.Default()
	r.Use(cors.Default())
	r.LoadHTMLGlob("views/**/*")
	r.Static("/assets", "./assets")
	app.Web = r
}
