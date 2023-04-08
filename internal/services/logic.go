package services

import (
	"github.com/gin-gonic/gin"
	dao2 "github.com/gogolang20/PaaS_platform/internal/dao"

	"net/http"
)

func Service(c *gin.Context) {
	dao2.Set("key", "val", 12)
	dao2.Create()

	c.JSON(http.StatusOK, gin.H{
		"Message": "next layer service",
	})
}
