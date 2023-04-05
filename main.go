package main

import (
	"github.com/gin-gonic/gin"

	"github.com/gogolang20/PaaS_platform/services/github"
	"github.com/gogolang20/PaaS_platform/services/jenkins"

	"github.com/gogolang20/PaaS_platform/apis"
)

func main() {
	jenkins.CreateJob()
	github.ListRepo()
	return

	router := gin.Default()
	apis.Apis(router)

	router.Run(":9000")
}
