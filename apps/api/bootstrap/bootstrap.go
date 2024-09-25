package bootstrap

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupServer(server *gin.Engine) {

	// routers
	testGroup := server.Group("/test")
	{
		testGroup.GET(
			"/alive",
			func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "pong")
			},
		)
	}
}
