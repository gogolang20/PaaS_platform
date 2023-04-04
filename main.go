package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gogolang20/PaaS_platform/apis"
	"github.com/gogolang20/PaaS_platform/services/github"
)

func main() {
	github.ListRepo()
	return

	router := gin.Default()
	apis.Apis(router)

	router.Run(":9000")
}
