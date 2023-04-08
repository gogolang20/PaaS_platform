package main

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/gogolang20/PaaS_platform/internal/apis"
	"github.com/gogolang20/PaaS_platform/internal/services/github"
	"github.com/gogolang20/PaaS_platform/internal/services/jenkins"
)

func main() {
	repo, _ := github.GetCloneURL(context.Background())
	jenkins.CreateJob(context.Background(), repo)
	return

	router := gin.Default()
	apis.Apis(router)

	router.Run(":9000")
}
