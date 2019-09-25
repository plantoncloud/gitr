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
	SshCloneUrl  string
	HttpCloneUrl string
	WebUrl       string
}

func (c RepoUrl) get_ssh_clone_url() string {
	return fmt.Sprintf("git@%s:%s.git", c.HostName, c.RepoPath)
}

func (c RepoUrl) get_http_clone_url() string {
	return fmt.Sprintf("%s://%s/%s.git", c.Protocol, c.HostName, c.RepoPath)
}

func get_absolute_path(pemFilePath string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if strings.HasPrefix(pemFilePath, "~/") {
		pemFilePath = filepath.Join(dir, pemFilePath[2:])
	}
	return pemFilePath
}

func is_browser_url(clone_url string) bool {
	return strings.HasSuffix(clone_url, ".git")
}

func get_scm_provider(hostname string) ScmProvider {
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

func get_repo_path(clone_url string, url_path string, scm_provider ScmProvider) string {
	org_or_team := strings.Split(url_path, "/")[1]
	repo_name := strings.Split(url_path, "/")[2]
	return fmt.Sprintf("%s/%s", org_or_team, repo_name)
}

func get_repo_name(repopath string) string {
	return string(repopath[strings.LastIndex(repopath, "/")+1 : len(repopath)])
}

func parse_url(clone_url string) RepoUrl {
	var clone_url_object = RepoUrl{}
	if !is_browser_url(clone_url) {
		clone_url_object.Protocol = strings.Split(clone_url, "://")[0]
		clone_url_object.HostName = strings.Split(strings.Split(clone_url, "://")[1], "/")[0]
		clone_url_object.UrlPath = string(clone_url[strings.Index(clone_url, clone_url_object.HostName)+len(clone_url_object.HostName) : len(clone_url)])
		clone_url_object.ScmProvider = get_scm_provider(clone_url_object.HostName)
		clone_url_object.RepoPath = get_repo_path(clone_url, clone_url_object.UrlPath, clone_url_object.ScmProvider)
		clone_url_object.RepoName = get_repo_name(clone_url_object.RepoPath)
	}
	return clone_url_object
}

func Parse(url string) RepoUrl {
	return parse_url(url)
}
