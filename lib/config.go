package lib

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

type ScmProvider string

const (
	GitHub              ScmProvider = "github"
	GitLab              ScmProvider = "gitlab"
	BitBucketCloud      ScmProvider = "bitbucket-cloud"
	BitBucketDatacenter ScmProvider = "bitbucket"
)

type GitrConfig struct {
	scmSystems []ScmSystem
}

type ScmSystem struct {
	Hostname string
	Provider ScmProvider
}

func (g *GitrConfig) loadScmSystems() {
	err := viper.UnmarshalKey("scmSystems", &g.scmSystems)
	if err != nil {
		log.Fatalf("unable to decode scm systems into array of struct, %v", err)
	}

	g.scmSystems = append(g.scmSystems,
		ScmSystem{"github.com", GitHub},
		ScmSystem{"gitlab.com", GitLab},
		ScmSystem{"bitbucket.org", BitBucketCloud})
}

func (g *GitrConfig) GetScmProvider(hostname string) (ScmProvider, error) {
	g.loadScmSystems()
	for _, scmSystem := range g.scmSystems {
		if scmSystem.Hostname == hostname {
			return scmSystem.Provider, nil
		}
	}
	return "", errors.New("SCM Provider Not Found for hostname " + hostname + " in ~/.gitr.yaml")
}
