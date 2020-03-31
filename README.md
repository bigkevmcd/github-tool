# github-tool

The start of a small tool for making changes in GitHub from the command-line, suitable for running from Tekton pipeline task.

All options can be configured from the environment.

## Creating a GitHub status notification

 ```shell
 ./github-tool --github-token <github access token> create-status --repo bigkevmcd/github-tool --sha c1f804c9a19c387f2d2febb8b30c984846233147 --state success --target-url https://example.com/testing --description "This is a test" --context "testing"
  ```

## Creating a GitHub Webhook

 ```shell
 ./github-tool --github-token <github access token> create-webhook --repo bigkevmcd/github-tool --name "My Test Hook" --target-url https://example.com/testing --secret "This is a test"
  ```


## docker image

  There is a Docker image at `quay.io/kmcdermo/github-tool`

## Testing documentation
