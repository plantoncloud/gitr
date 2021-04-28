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
	ScmSystems []ScmSystem
	Clone      GitrCloneConfig
}

type GitrCloneConfig struct {
	ScmHome              string
	AlwaysCreDir         bool
	IncludeHostForCreDir bool
}

type ScmSystem struct {
	Hostname      string
	Provider      ScmProvider
	DefaultBranch string
}

func (g *GitrConfig) loadConfig() {
	err := viper.UnmarshalKey("Clone", &g.Clone)
	if err != nil {
		log.Fatalf("unable to decode Clone config from config file, %v", err)
	}
	g.loadScmSystems()
}

func (g *GitrConfig) loadScmSystems() {
	err := viper.UnmarshalKey("ScmSystems", &g.ScmSystems)
	if err != nil {
		log.Fatalf("unable to decode scm systems into array of struct, %v", err)
	}

	g.ScmSystems = append(g.ScmSystems,
		ScmSystem{"github.com", GitHub, "master"},
		ScmSystem{"gitlab.com", GitLab, "main"},
		ScmSystem{"bitbucket.org", BitBucketCloud, "master"})
}

func (g *GitrConfig) Get() *GitrConfig {
	g.loadConfig()
	return g
}

func (g *GitrConfig) GetScmSystem(hostname string) (*ScmSystem, error) {
	g.loadConfig()
	for _, scmSystem := range g.ScmSystems {
		if scmSystem.Hostname == hostname {
			return &scmSystem, nil
		}
	}
	return nil, errors.New("scm provider not found for hostname " + hostname + " in ~/.gitr.yaml")
}
