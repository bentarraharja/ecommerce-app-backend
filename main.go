package main

import (
	"MyEcommerce/app/config"
	"MyEcommerce/app/database"
	"MyEcommerce/app/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	dbSql := database.InitDBMysql(cfg)

	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())

	// e.Use(middleware.Logger())

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))
	
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	router.InitRouter(dbSql, e)
	//start server and port
	e.Logger.Fatal(e.Start(":8000"))
}
