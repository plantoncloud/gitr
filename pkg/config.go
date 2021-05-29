package pkg

import (
	"errors"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"log"
)

type ScmProvider string

type HttpScheme string

const (
	GitHub              ScmProvider = "github"
	GitLab              ScmProvider = "gitlab"
	BitBucketCloud      ScmProvider = "bitbucket-cloud"
	BitBucketDatacenter ScmProvider = "bitbucket"
	Http                HttpScheme  = "http"
	Https               HttpScheme  = "https"
)

type GitrConfig struct {
	ScmSystems []ScmSystem     `yaml:"scmSystems"`
	Clone      GitrCloneConfig `yaml:"clone"`
}

type GitrCloneConfig struct {
	ScmHome              string `yaml:"scmHome"`
	AlwaysCreDir         bool   `yaml:"alwaysCreDir"`
	IncludeHostForCreDir bool   `yaml:"includeHostForCreDir"`
}

type ScmSystem struct {
	Scheme        HttpScheme  `yaml:"scheme"`
	Hostname      string      `yaml:"hostname"`
	Provider      ScmProvider `yaml:"provider"`
	DefaultBranch string      `yaml:"defaultBranch"`
}

func defaultScmSystems() []ScmSystem {
	return []ScmSystem{
		{Https, "github.com", GitHub, "master"},
		{Https, "gitlab.com", GitLab, "main"},
		{Https, "bitbucket.org", BitBucketCloud, "master"},
	}
}

func GetScmSystem(hostname string) (*ScmSystem, error) {
	cfg, err := ReadGitrConfig()
	if err != nil {
		return nil, err
	}
	//return the scm system from config file
	for _, scmSystem := range cfg.ScmSystems {
		if scmSystem.Hostname == hostname {
			return &scmSystem, nil
		}
	}
	//return the scm system from default scm systems if one is not found in config file
	for _, scmSystem := range defaultScmSystems() {
		if scmSystem.Hostname == hostname {
			return &scmSystem, nil
		}
	}
	return nil, errors.New("scm provider not found for hostname " + hostname)
}

func LoadViperConfig() {
	home, _ := homedir.Dir()
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".gitr")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("failed to read config file")
	}
}

func ReadGitrConfig() (*GitrConfig, error) {
	LoadViperConfig()
	cfg := &GitrConfig{}
	err := viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
