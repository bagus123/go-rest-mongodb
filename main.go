package main

import (
	"github.com/bagus123/go-rest-mongodb/daos"
	adminRouter "github.com/bagus123/go-rest-mongodb/routers/v1/admin"
	apiRouter "github.com/bagus123/go-rest-mongodb/routers/v1/api"
	"github.com/bagus123/go-rest-mongodb/seeds"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	// init DAO
	daos.InitDAO()

	// running user seed
	seeds.RunUserSeed()

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.NoRoute(noRouteHandler())

	apiV1 := router.Group("/v1/api")
	adminV1 := router.Group("/v1/admin")

	apiRouter.Init(apiV1)
	adminRouter.Init(adminV1)

	router.Run(":8080")
}

func noRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, "Page Not Found")
	}
}
