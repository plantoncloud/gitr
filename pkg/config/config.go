package config

import (
	"fmt"
	"github.com/leftbin/go-util/pkg/file"
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

func EnsureInitialConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get user home dir")
	}
	gitrConfigFile := fmt.Sprintf("%s/.gitr.yaml", homeDir)
	cfg := NewDefaultConfig()
	d, err := yaml.Marshal(&cfg)
	if err != nil {
		return errors.Wrap(err, "failed to marshal config")
	}
	if !file.IsFileExists(gitrConfigFile) {
		if err = os.WriteFile(gitrConfigFile, d, 0644); err != nil {
			return errors.Wrapf(err, "failed to write file %s", gitrConfigFile)
		}
	}
	return nil
}

func GetScmHost(cfg *GitrConfig, hostname string) (*ScmHost, error) {
	//return the scm system from config file
	for _, scmSystem := range cfg.Scm.Hosts {
		if scmSystem.Hostname == hostname {
			return scmSystem, nil
		}
	}
	return nil, &UnknownScmHostErr{ScmHost: hostname}
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

func NewDefaultConfig() *GitrConfig {
	return &GitrConfig{
		CopyRepoPathCdCmdToClipboard: false,
		Scm: &Scm{
			HomeDir: "",
			Hosts:   defaultScmSystems(),
		},
	}
}

func defaultScmSystems() []*ScmHost {
	defaultCloneConfig := &CloneConfig{
		HomeDir:              "",
		AlwaysCreDir:         true,
		IncludeHostForCreDir: true,
	}
	return []*ScmHost{
		{Scheme: Https, Hostname: "github.com", Provider: GitHub, DefaultBranch: "master", Clone: defaultCloneConfig},
		{Scheme: Https, Hostname: "gitlab.com", Provider: GitLab, DefaultBranch: "main", Clone: defaultCloneConfig},
		{Scheme: Https, Hostname: "bitbucket.org", Provider: BitBucketCloud, DefaultBranch: "master", Clone: defaultCloneConfig},
	}
}
