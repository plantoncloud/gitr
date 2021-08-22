package config

import (
	"fmt"
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	"gopkg.in/yaml.v2"
	"os"
)

func EnsureInitialConfig() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get user home dir")
	}
	gitrConfigFile := fmt.Sprintf("%s/.gitr.yaml", homeDir)
	cfg := getDefaultConfig()
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

func getDefaultConfig() *config.GitrConfig {
	return &config.GitrConfig{
		CopyRepoPathCdCmdToClipboard: false,
		Scm: &config.Scm{
			HomeDir: "",
			Hosts:   defaultScmSystems(),
		},
	}
}

func defaultScmSystems() []*config.ScmHost {
	defaultCloneConfig := &config.CloneConfig{
		HomeDir:              "",
		AlwaysCreDir:         true,
		IncludeHostForCreDir: true,
	}
	return []*config.ScmHost{
		{Scheme: config.Https, Hostname: "github.com", Provider: config.GitHub, DefaultBranch: "master", Clone: defaultCloneConfig},
		{Scheme: config.Https, Hostname: "gitlab.com", Provider: config.GitLab, DefaultBranch: "main", Clone: defaultCloneConfig},
		{Scheme: config.Https, Hostname: "bitbucket.org", Provider: config.BitBucketCloud, DefaultBranch: "master", Clone: defaultCloneConfig},
	}
}
