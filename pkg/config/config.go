package config

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
	CopyCloneLocationCdCmdToClipboard bool `yaml:"copyCloneLocationCdCmdToClipboard"`
	Scm                               Scm  `yaml:"scm"`
}

type Scm struct {
	Hosts   []ScmHost `yaml:"hosts"`
	HomeDir string    `yaml:"homeDir"`
}

type ScmHost struct {
	Hostname      string      `yaml:"hostname"`
	Provider      ScmProvider `yaml:"provider"`
	DefaultBranch string      `yaml:"defaultBranch"`
	Clone         CloneConfig `yaml:"clone"`
	Scheme        HttpScheme  `yaml:"scheme"`
}

type CloneConfig struct {
	HomeDir              string `yaml:"homeDir"`
	AlwaysCreDir         bool   `yaml:"alwaysCreDir"`
	IncludeHostForCreDir bool   `yaml:"includeHostForCreDir"`
}

func GetScmHost(cfg *GitrConfig, hostname string) (*ScmHost, error) {
	//return the scm system from config file
	for _, scmSystem := range cfg.Scm.Hosts {
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

func GetGitrConfig() *GitrConfig {
	var cfg GitrConfig
	err := viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return &cfg
}
