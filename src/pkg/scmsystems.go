package pkg

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
	BitBucketDatacenter ScmProvider = "bitbucket-datacenter"
)

type ScmSystem struct {
	Hostname string
	Scm      ScmProvider
}

func getScmSystems() []ScmSystem {
	var scmSystems []ScmSystem
	err := viper.UnmarshalKey("scmSystems", &scmSystems)
	if err != nil {
		log.Fatalf("unable to decode scm systems into array of struct, %v", err)
	}
	return scmSystems
}

func getScmProvider(hostname string) (ScmProvider, error) {
	var scmSystems = getScmSystems()
	for _,scmSystem := range scmSystems {
		if scmSystem.Hostname ==  hostname {
			return scmSystem.Scm, nil
		}
	}
	return "", errors.New("SCM Provider Not Found for hostname " + hostname + " in ~/.gitr.yaml")
}