package main

import (
	"context"
	"errors"
	"strings"

	"github.com/google/go-github/v28/github"
	"github.com/urfave/cli/v2"
)

var (
	repoFlag        = "repo"
	shaFlag         = "sha"
	stateFlag       = "state"
	targetURLFlag   = "target-url"
	descriptionFlag = "description"
	contextFlag     = "context"

	createStatusFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     repoFlag,
			Value:    "",
			Usage:    "The fullname for the repository e.g. bigkevmcd/go-github-status",
			EnvVars:  []string{"STATUS_REPO"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     shaFlag,
			Value:    "",
			Usage:    "The sha to create a state for",
			EnvVars:  []string{"STATUS_SHA"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     stateFlag,
			Value:    "",
			Usage:    "The state of the status. Can be one of error, failure, pending, or success.",
			EnvVars:  []string{"STATUS_STATE"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     targetURLFlag,
			Value:    "",
			Usage:    "The target URL to associate with this status. This URL will be linked from the GitHub UI to allow users to easily see the source of the status.",
			EnvVars:  []string{"STATUS_TARGET_URL"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     descriptionFlag,
			Value:    "",
			Usage:    "A short description of the status.",
			EnvVars:  []string{"STATUS_DESCRIPTION"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     contextFlag,
			Value:    "default",
			Usage:    "A string label to differentiate this status from the status of other systems",
			EnvVars:  []string{"STATUS_CONTEXT"},
			Required: false,
		},
	}
)

func createStatus(c *cli.Context) error {
	sha := c.String(shaFlag)
	status := createRepoStatus(c)
	_, _, err := createClient(c.String(accessTokenFlag)).Repositories.CreateStatus(context.Background(), "bigkevmcd", "taxi-stage-config", sha, status)

	return err
}

func createRepoStatus(c *cli.Context) *github.RepoStatus {
	return &github.RepoStatus{
		TargetURL:   github.String(c.String(targetURLFlag)),
		State:       github.String(c.String(stateFlag)),
		Context:     github.String(c.String(contextFlag)),
		Description: github.String(c.String(descriptionFlag)),
	}
}

func splitFullname(s string) (string, string, error) {
	seg := strings.Split(s, "/")
	if len(seg) != 2 {
		return "", "", errors.New("fullname must be owner/repo")
	}
	return seg[0], seg[1], nil
}
