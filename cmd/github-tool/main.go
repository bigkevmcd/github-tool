package main

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v28/github"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

const (
	accessTokenFlag = "access-token"
)

var (
	globalFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     accessTokenFlag,
			Value:    "",
			Usage:    "oauth access token to authenticate the request",
			EnvVars:  []string{"STATUS_ACCESS_TOKEN"},
			Required: true,
		},
	}
)

func main() {
	app := &cli.App{
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

func createClient(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
}
