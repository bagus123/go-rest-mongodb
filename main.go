package main

import (
	"github.com/bagus123/go-rest-mongodb/daos"

	adminRouter "github.com/bagus123/go-rest-mongodb/routers/v1/admin"
	apiRouter "github.com/bagus123/go-rest-mongodb/routers/v1/api"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {

	//init DAO
	daos.InitDAO()

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.NoRoute(noRouteHandler())

	apiV1 := router.Group("/v1/api")
	adminV1 := router.Group("/v1/admin")

	apiRouter.Init(apiV1)
	adminRouter.Init(adminV1)

	// newUser := models.User{
	// 	ID:        primitive.NewObjectID(),
	// 	UserName:  "bagus",
	// 	Password:  "123",
	// 	Role:      "Admin",
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }
	//result := daos.User().InsertOne(newUser)
	//results := daos.User().GetAll()

	//log.Println(result)
	// log.Println(results)

	// for _, value := range results {
	// 	// casting from interface{} to User
	// 	user, ok := value.(models.User)
	// 	if ok {
	// 		log.Println("UserName", user.UserName, "Role", user.Role)
	// 	}

	// }

	router.Run(":8080")
}

func noRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, "Page Not Found")
	}
}
