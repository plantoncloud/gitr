package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
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

func GetScmHost(cfg *GitrConfig, hostname string) (*ScmHost, error) {
	//return the scm system from config file
	for _, scmSystem := range cfg.Scm.Hosts {
		if scmSystem.Hostname == hostname {
			return &scmSystem, nil
		}
	}
	return nil, errors.New("scm provider not found for hostname " + hostname)
}

func NewGitrConfig() (*GitrConfig, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read user home directory required for reading ${HOME}/.gitr.yaml")
	}
	gitrConfigYaml := filepath.Join(homeDir, ".gitr.yaml")
	file, err := os.ReadFile(gitrConfigYaml)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s file", gitrConfigYaml)
	}
	var cfg GitrConfig
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal %s file", gitrConfigYaml)
	}
	return &cfg, nil
}
