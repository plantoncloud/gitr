package repo

import (
	"github.com/pkg/errors"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	gitrgit "github.com/swarupdonepudi/gitr/v2/pkg/git"
	"github.com/swarupdonepudi/gitr/v2/pkg/url"
	"os"
)

func GetRepoPathOnHost(remoteUrl string, gitrCfg *config.GitrConfig) (string, error) {
	scmHost := url.GetHostname(remoteUrl)
	scmHostCfg, err := config.GetScmHost(gitrCfg, scmHost)
	if err != nil {
		return "", errors.Wrap(err, "failed to get scm host config")
	}
	return url.GetRepoPath(remoteUrl, scmHost, scmHostCfg.Provider), nil
}

func GetPathOnScmFromPwd() (string, error) {
	gitrCfg, err := config.NewGitrConfig()
	if err != nil {
		return "", errors.Wrap(err, "failed to read gitr config")
	}
	pwd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "failed to read current dir")
	}
	gitRepo := gitrgit.GetGitRepo(pwd)
	remoteUrl := gitrgit.GetGitRemoteUrl(gitRepo)
	projectPath, err := GetRepoPathOnHost(remoteUrl, gitrCfg)
	if err != nil {
		return "", errors.Wrap(err, "failed to get repo path on host")
	}
	return projectPath, nil
}
