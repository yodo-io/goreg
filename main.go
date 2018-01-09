package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "Go reg"
	app.Usage = "Simple docker registry client"

	loginFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "registry-url, r",
			Usage:  "registry url to query",
			EnvVar: "REGISTRY_URL",
		},
		cli.StringFlag{
			Name:   "username, u",
			Usage:  "registry login user",
			EnvVar: "REGISTRY_USERNAME",
		},
		cli.StringFlag{
			Name:   "password, p",
			Usage:  "registry password",
			EnvVar: "REGISTRY_PASSWORD",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "list repositories",
			Action:  actionListRepos,
			Flags:   append(loginFlags, []cli.Flag{}...),
		},
		{
			Name:   "tags",
			Usage:  "list tags for a repo",
			Action: actionListTags,
			Flags: append(loginFlags, []cli.Flag{
				cli.StringFlag{
					Name:   "repository, repo",
					Usage:  "repository name",
					EnvVar: "REPOSITORY_NAME",
				},
			}...),
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func actionListRepos(c *cli.Context) error {
	login := loginOpts(c)
	return ListRepos(login)
}

func actionListTags(c *cli.Context) error {
	login := loginOpts(c)
	repo := c.String("repository")
	return ListTags(login, repo)
}

func loginOpts(c *cli.Context) RegistryLogin {
	return RegistryLogin{
		url:      addProto(c.String("registry-url")),
		username: c.String("username"),
		password: c.String("password"),
	}
}

func addProto(repositoryURL string) string {
	if len(repositoryURL) == 0 || strings.HasPrefix(repositoryURL, "http") {
		return repositoryURL
	}
	return fmt.Sprintf("https://%s", repositoryURL)
}
