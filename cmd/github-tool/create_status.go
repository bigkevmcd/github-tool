package main

import (
	"context"

	"github.com/jenkins-x/go-scm/scm"
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
			Required: false,
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
	repo := c.String(repoFlag)
	token := c.String(githubTokenFlag)
	status := createRepoStatus(c)
	_, _, err := createClient(token).Repositories.CreateStatus(context.Background(), repo, sha, status)
	return err
}

func createRepoStatus(c *cli.Context) *scm.StatusInput {
	si := &scm.StatusInput{
		State: convertState(c.String(stateFlag)),
		Label: c.String(contextFlag),
		Desc:  c.String(descriptionFlag),
	}
	if targetURL := c.String(targetURLFlag); targetURL != "" {
		si.Target = targetURL
	}
	return si
}

func convertState(s string) scm.State {
	switch s {
	case "error":
		return scm.StateError
	case "failure":
		return scm.StateFailure
	case "pending":
		return scm.StatePending
	case "success":
		return scm.StateSuccess
	default:
		return scm.StateUnknown
	}
}
