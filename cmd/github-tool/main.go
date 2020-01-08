package main

import (
	"context"
	"log"
	"os"

	"github.com/jenkins-x/go-scm/scm"
	"github.com/jenkins-x/go-scm/scm/driver/github"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

const (
	githubTokenFlag = "github-token"
)

var (
	globalFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     githubTokenFlag,
			Value:    "",
			Usage:    "oauth access token to authenticate the request",
			EnvVars:  []string{"GITHUB_TOKEN"},
			Required: true,
		},
	}
)

func main() {
	app := &cli.App{
		Name:  "github-tool",
		Usage: "command-line access to GitHub",
		Commands: []*cli.Command{
			{
				Name:   "create-status",
				Usage:  "submit a commit-status to GitHub",
				Flags:  createStatusFlags,
				Action: createStatus,
			},
		},
		Flags: globalFlags,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func createClient(token string) *scm.Client {
	client := github.NewDefault()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	client.Client = oauth2.NewClient(context.Background(), ts)
	return client
}
