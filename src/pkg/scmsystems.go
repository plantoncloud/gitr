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
	BitBucketDatacenter ScmProvider = "bitbucket"
)

type ScmSystem struct {
	Hostname string
	Scm      ScmProvider
}


func getScmSystems() []ScmSystem {
	var scmSystems []ScmSystem

	var github = ScmSystem{}
	github.Hostname = "github.com"
	github.Scm = GitHub

	var gitlab = ScmSystem{}
	gitlab.Hostname = "gitlab.com"
	gitlab.Scm = GitLab

	var bitbucketCloud = ScmSystem{}
	bitbucketCloud.Hostname = "bitbucket.org"
	bitbucketCloud.Scm = BitBucketCloud

	err := viper.UnmarshalKey("scmSystems", &scmSystems)
	if err != nil {
		log.Fatalf("unable to decode scm systems into array of struct, %v", err)
	}

	scmSystems = append(scmSystems, github)
	scmSystems = append(scmSystems, gitlab)
	scmSystems = append(scmSystems, bitbucketCloud)

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
