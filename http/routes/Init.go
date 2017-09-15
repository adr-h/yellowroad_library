package routes

import (
	"fmt"
	"yellowroad_library/containers"
	"yellowroad_library/http/routes/book_route"

	"github.com/gin-gonic/gin"
	"yellowroad_library/http/routes/user_route"
)

func Init(container containers.Container) {
	var ginEngine = gin.Default()
	var r = newAppRouter(ginEngine, container)
	var portString = fmt.Sprintf(":%d", container.GetConfiguration().Web.Port)

	r.Route("/api/books", book_route.Register)
	r.Route("/api/users", user_route.Register)

	ginEngine.Run(portString)
}