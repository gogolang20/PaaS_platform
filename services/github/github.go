package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

const (
	token = "... your access token ..."
)

var (
	client *github.Client
)

func init() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)
}

func ListRepo() {
	ctx := context.Background()

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		logrus.Error("[][] error")
		return
	}

	fmt.Println(repos)
}
