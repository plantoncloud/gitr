package internal

import (
	"fmt"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	"github.com/swarupdonepudi/gitr/v2/pkg/file"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func EnsureInitialConfig() {
	gitrConfigFile := fmt.Sprintf("%s/.gitr.yaml", file.GetHomeDir())
	cfg := getDefaultConfig()
	d, err := yaml.Marshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	if !file.IsFileExists(gitrConfigFile) {
		err = ioutil.WriteFile(gitrConfigFile, d, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getDefaultConfig() *config.GitrConfig {
	return &config.GitrConfig{
		CopyRepoPathCdCmdToClipboard: false,
		Scm: config.Scm{
			HomeDir: "",
			Hosts:   defaultScmSystems(),
		},
	}
}

func defaultScmSystems() []config.ScmHost {
	defaultCloneConfig := config.CloneConfig{
		HomeDir:              "",
		AlwaysCreDir:         true,
		IncludeHostForCreDir: true,
	}
	return []config.ScmHost{
		{Scheme: config.Https, Hostname: "github.com", Provider: config.GitHub, DefaultBranch: "master", Clone: defaultCloneConfig},
		{Scheme: config.Https, Hostname: "gitlab.com", Provider: config.GitLab, DefaultBranch: "main", Clone: defaultCloneConfig},
		{Scheme: config.Https, Hostname: "bitbucket.org", Provider: config.BitBucketCloud, DefaultBranch: "master", Clone: defaultCloneConfig},
	}
}
