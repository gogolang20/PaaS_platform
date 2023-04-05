package services

import (
	"github.com/gin-gonic/gin"

	"github.com/gogolang20/PaaS_platform/dao"

	"net/http"
)

func Service(c *gin.Context) {
	dao.Set("key", "val", 12)
	dao.Create()

	c.JSON(http.StatusOK, gin.H{
		"Message": "next layer service",
	})
}
