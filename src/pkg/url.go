package pkg

import (
	"fmt"
	"os/user"
	"path/filepath"
	"strings"
)

type ScmProvider string

const (
	GitHub    ScmProvider = "GitHub"
	GitLab    ScmProvider = "GitLab"
	BitBucket ScmProvider = "BitBucket"
)

type RepoUrl struct {
	Protocol    string
	HostName    string
	UrlPath     string
	ScmProvider ScmProvider
	RepoPath    string
	RepoName    string
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

func (c RepoUrl) GetPrsUrl() string {
	switch c.ScmProvider {
	case GitHub:
		return fmt.Sprintf("%s/pulls", c.GetWebUrl())
	case GitLab:
		return fmt.Sprintf("%s/merge_requests", c.GetWebUrl())
	case BitBucket:
		return fmt.Sprintf("%s/pull-requests", c.GetWebUrl())
	default:
		return ""
	}
}

func (c RepoUrl) GetBranchesUrl() string {
	switch c.ScmProvider {
	case GitLab:
		return fmt.Sprintf("%s/-/branches", c.GetWebUrl())
	default:
		return fmt.Sprintf("%s/branches", c.GetWebUrl())
	}
}

func (c RepoUrl) GetCommitsUrl() string {
	switch c.ScmProvider {
	case GitLab:
		return fmt.Sprintf("%s/-/commits", c.GetWebUrl())
	default:
		return fmt.Sprintf("%s/commits", c.GetWebUrl())
	}
}

func (c RepoUrl) GetIssuesUrl() string {
	switch c.ScmProvider {
	case BitBucket:
		return ""
	default:
		return fmt.Sprintf("%s/issues", c.GetWebUrl())
	}
}

func (c RepoUrl) GetReleasesUrl() string {
	switch c.ScmProvider {
	case GitHub:
		return fmt.Sprintf("%s/releases", c.GetWebUrl())
	default:
		return ""
	}
}

func (c RepoUrl) GetPipelinesUrl() string {
	switch c.ScmProvider {
	case GitLab:
		return fmt.Sprintf("%s/pipelines", c.GetWebUrl())
	case BitBucket:
		return fmt.Sprintf("%s/addon/pipelines/home", c.GetWebUrl())
	default:
		return ""
	}
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
