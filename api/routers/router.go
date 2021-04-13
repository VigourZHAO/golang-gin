package routers

import (
	"github.com/gin-gonic/gin"
	"vigourzhao/golang-gin/tools"
)

// @Summary Test
// @Produce  json
// @Param id path int true "ID"
// @Param name query string true "ID"
// @Param state query int false "State"
// @Param modified_by query string true "ModifiedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /test [get]
func SetupRouter() *gin.Engine {
	router := gin.Default()

	tools.SwaggerWraps(router)

	router.GET("/test", func(context *gin.Context) {
		context.String(200, "wa! is running.")
	})
	return router
}
