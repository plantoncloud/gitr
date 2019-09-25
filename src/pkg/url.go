package pkg

import (
	"fmt"
	"os/user"
	"path/filepath"
	"strings"
)

type ScmProvider string

const (
	GitHub    ScmProvider = "github.com"
	GitLab    ScmProvider = "gitlab.com"
	BitBucket ScmProvider = "bitbucket.com"
)

type RepoUrl struct {
	Protocol     string
	HostName     string
	UrlPath      string
	ScmProvider  ScmProvider
	RepoPath     string
	RepoName     string
}

func (c RepoUrl) GetSshCloneUrl() string {
	return fmt.Sprintf("git@%s:%s.git", c.HostName, c.RepoPath)
}

func (c RepoUrl) GetHttpCloneUrl() string {
	return fmt.Sprintf("%s://%s/%s.git", c.Protocol, c.HostName, c.RepoPath)
}

func (c RepoUrl) GetWebUrl() string {
	return fmt.Sprintf("%s://%s/%s", c.Protocol, c.HostName, c.RepoPath)
}


func getAbsolutePath(pemFilePath string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if strings.HasPrefix(pemFilePath, "~/") {
		pemFilePath = filepath.Join(dir, pemFilePath[2:])
	}
	return pemFilePath
}

func isGitSshUrl(repo_url string) bool {
	return strings.HasPrefix(repo_url, "git@")
}

func getScmProvider(hostname string) ScmProvider {
	switch hostname {
	case "github.com":
		return GitHub
	case "gitlab.com":
		return GitLab
	case "bitbucket.org":
		return BitBucket
	default:
		return GitHub
	}
}

func getRepoPath(url_path string) string {
	orgOrTeam := strings.Split(url_path, "/")[0]
	repoName := strings.Split(url_path, "/")[1]
	return fmt.Sprintf("%s/%s", orgOrTeam, repoName)
}

func getRepoName(repopath string) string {
	return string(repopath[strings.LastIndex(repopath, "/")+1 : len(repopath)])
}

func parseUrl(repoUrl string) RepoUrl {
	var repoUrlObject = RepoUrl{}
	if (isGitSshUrl(repoUrl)) {
		repoUrlObject.Protocol = "https"
		repoUrlObject.HostName = strings.Split(strings.Split(repoUrl, "@")[1], ":")[0]
		repoUrlObject.UrlPath = string(repoUrl[strings.Index(repoUrl, repoUrlObject.HostName)+len(repoUrlObject.HostName)+1 : strings.Index(repoUrl, ".git")])
	} else {
		repoUrlObject.Protocol = strings.Split(repoUrl, "://")[0]
		repoUrlObject.HostName = strings.Split(strings.Split(repoUrl, "://")[1], "/")[0]
		repoUrlObject.UrlPath = string(repoUrl[strings.Index(repoUrl, repoUrlObject.HostName)+1+len(repoUrlObject.HostName):])
	}
	repoUrlObject.ScmProvider = getScmProvider(repoUrlObject.HostName)
	repoUrlObject.RepoPath = getRepoPath(repoUrlObject.UrlPath)
	repoUrlObject.RepoName = getRepoName(repoUrlObject.RepoPath)
	return repoUrlObject
}

func Parse(repoUrl string) RepoUrl {
	return parseUrl(repoUrl)
}
