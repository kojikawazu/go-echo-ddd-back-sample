package router

import (
	interfaces_paralell "backend/internal/interfaces/paralell"
	interfaces_sample "backend/internal/interfaces/sample"
	interfaces_user "backend/internal/interfaces/user"

	"github.com/labstack/echo/v4"
)

// ルーティングの設定
func SetUpRouter(e *echo.Echo, sampleHandler *interfaces_sample.SampleHandler, paralellHandler *interfaces_paralell.ParalellHandler, userHandler *interfaces_user.UserHandler) {
	api := e.Group("/api")
	{
		sample := api.Group("/sample")
		{
			sample.GET("", sampleHandler.ExecSample)
		}

		fetch := api.Group("/fetch")
		{
			fetch.GET("/parallel", paralellHandler.ExecParallel)
			fetch.GET("/series", paralellHandler.ExecSeries)
		}
		user := api.Group("/user")
		{
			user.GET("", userHandler.GetAllUsers)
		}
	}
}
