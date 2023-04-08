package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type SoruceCode interface {
	GetCloneURL(ctx context.Context) (string, error)
}

type GithubClient struct {
	client *github.Client
}

func NewGithubClient() SoruceCode {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client = github.NewClient(tc)
	return &GithubClient{client: client}
}

func (gc *GithubClient) GetCloneURL(ctx context.Context) (string, error) {
	repos, _, err := gc.client.Repositories.List(ctx, "", nil)
	if err != nil {
		logrus.Error("[github][GetCloneURL] error.")
		return "", nil
	}

	url := ""
	for _, repo := range repos {
		if *repo.Name == "template" {
			url = *repo.CloneURL
			fmt.Println(*repo.CloneURL)
		}
	}

	return url, nil
}

const (
	token = ""
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

func GetCloneURL(ctx context.Context) (string, error) {
	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		logrus.Error("[github][GetCloneURL] error.")
		return "", nil
	}

	url := ""
	for _, repo := range repos {
		if *repo.Name == "template" {
			url = *repo.CloneURL
			fmt.Println(*repo.CloneURL)
		}
	}

	return url, nil
}
