package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/yodo-io/docker-registry-client/registry"
)

// RegistryLogin defines registry connection parameters
type RegistryLogin struct {
	url      string
	username string
	password string
}

func validateLogin(r RegistryLogin) error {
	if r.url == "" {
		return errors.New("Registry url must not be empty")
	}
	return nil
}

// ListRepos lists all repositories in the registry
func ListRepos(login RegistryLogin) error {
	registry, err := newRegistry(login)

	if err != nil {
		return err
	}

	repos, err := registry.Repositories()

	if err != nil {
		return err
	}

	for _, repo := range repos {
		fmt.Println(repo)
	}

	return nil
}

// ListTags lists all tags for given repo
func ListTags(login RegistryLogin, repo string) error {
	if repo == "" {
		return errors.New("Repo url must not be empty")
	}

	registry, err := newRegistry(login)

	if err != nil {
		return err
	}

	tags, err := registry.Tags(repo)

	for _, tag := range tags {
		fmt.Println(tag)
	}

	return nil
}

func newRegistry(login RegistryLogin) (*registry.Registry, error) {
	if err := validateLogin(login); err != nil {
		return nil, err
	}

	transport := &http.Transport{
		ResponseHeaderTimeout: time.Second * 10,
		ExpectContinueTimeout: time.Second * 10,
	}

	registry, err := registry.NewFromTransport(
		login.url,
		login.username,
		login.password,
		transport,
	)

	if err != nil {
		return nil, err
	}

	return registry, nil
}
