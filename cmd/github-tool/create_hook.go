package main

import (
	"context"

	"github.com/jenkins-x/go-scm/scm"
	"github.com/urfave/cli/v2"
)

var (
	nameFlag          = "name"
	webhookSecretFlag = "secret"

	createHookFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     repoFlag,
			Value:    "",
			Usage:    "The fullname for the repository e.g. bigkevmcd/go-github-status",
			EnvVars:  []string{"CREATE_HOOK_REPO"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     targetURLFlag,
			Value:    "",
			Usage:    "The target URL to associate with this status. This URL will be linked from the GitHub UI to allow users to easily see the source of the status.",
			EnvVars:  []string{"CREATE_HOOK_TARGET_URL"},
			Required: false,
		},
		&cli.StringFlag{
			Name:     nameFlag,
			Value:    "",
			Usage:    "A name for the webhook to identify it",
			EnvVars:  []string{"CREATE_HOOK_NAME"},
			Required: true,
		},
		&cli.StringFlag{
			Name:     webhookSecretFlag,
			Value:    "",
			Usage:    "A shared secret used to validate incoming requests",
			EnvVars:  []string{"CREATE_HOOK_NAME"},
			Required: true,
		},
	}
)

func createHook(c *cli.Context) error {
	repo := c.String(repoFlag)
	token := c.String(githubTokenFlag)
	hookInput := createHookInput(c)
	_, _, err := createClient(token).Repositories.CreateHook(context.Background(), repo, hookInput)
	return err
}

func createHookInput(c *cli.Context) *scm.HookInput {
	hi := &scm.HookInput{
		Name:       c.String(nameFlag),
		Target:     c.String(targetURLFlag),
		Secret:     c.String(webhookSecretFlag),
		Events:     defaultHookEvents(),
		SkipVerify: false,
	}
	return hi
}

// TODO: provide some way to configure these from the cmd-line.
//
// By default it's setup for PullRequest and Push notifications.
func defaultHookEvents() scm.HookEvents {
	return scm.HookEvents{
		Branch:             false,
		Issue:              false,
		IssueComment:       false,
		PullRequest:        true,
		PullRequestComment: false,
		Push:               true,
		ReviewComment:      false,
		Tag:                false,
	}
}
