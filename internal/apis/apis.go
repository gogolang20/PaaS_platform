package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gogolang20/PaaS_platform/internal/services"
)

func Apis(engine *gin.Engine) {
	engine.GET("/app/v1/logic", services.Service)
}
