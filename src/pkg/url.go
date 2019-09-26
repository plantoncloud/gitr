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

type GitrRepo struct {
	Protocol      string
	HostName      string
	UrlPath       string
	ScmProvider   ScmProvider
	RepoPath      string
	RepoName      string
	GitRemSshUrl  string
	GitRemHttpUrl string
}

func (c GitrRepo) ToString() string {
	return fmt.Sprintf("ScmProvider\t: \t%s" +
		"\nProtocol\t: \t%s" +
		"\nHostName\t: \t%s" +
		"\nUrlPath\t\t: \t%s" +
		"\nRepoPath\t: \t%s" +
		"\nGitRemSshUrl\t: \t%s" +
		"\nGitRemHttpUrl\t: \t%s" +
		"\nWeb\t\t: \t%s" +
		"\nPRs\t\t: \t%s" +
		"\nBranches\t: \t%s" +
		"\nCommits\t\t: \t%s" +
		"\nIssues\t\t: \t%s" +
		"\nReleases\t: \t%s" +
		"\nPipelines\t: \t%s",
		c.ScmProvider, c.Protocol, c.HostName, c.UrlPath, c.RepoPath, c.GitRemSshUrl,
		c.GitRemHttpUrl, c.GetWebUrl(), c.GetPrsUrl(), c.GetBranchesUrl(), c.GetCommitsUrl(),
		c.GetIssuesUrl(), c.GetReleasesUrl(), c.GetPipelinesUrl())
}

func (c GitrRepo) GetWebUrl() string {
	return fmt.Sprintf("%s://%s/%s", c.Protocol, c.HostName, c.RepoPath)
}

func (c GitrRepo) GetPrsUrl() string {
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

func (c GitrRepo) GetBranchesUrl() string {
	switch c.ScmProvider {
	case GitLab:
		return fmt.Sprintf("%s/-/branches", c.GetWebUrl())
	default:
		return fmt.Sprintf("%s/branches", c.GetWebUrl())
	}
}

func (c GitrRepo) GetCommitsUrl() string {
	switch c.ScmProvider {
	case GitLab:
		return fmt.Sprintf("%s/-/commits", c.GetWebUrl())
	default:
		return fmt.Sprintf("%s/commits", c.GetWebUrl())
	}
}

func (c GitrRepo) GetIssuesUrl() string {
	switch c.ScmProvider {
	case BitBucket:
		return ""
	default:
		return fmt.Sprintf("%s/issues", c.GetWebUrl())
	}
}

func (c GitrRepo) GetReleasesUrl() string {
	switch c.ScmProvider {
	case GitHub:
		return fmt.Sprintf("%s/releases", c.GetWebUrl())
	default:
		return ""
	}
}

func (c GitrRepo) GetPipelinesUrl() string {
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

func isGitUrl(repoUrl string) bool {
	return strings.HasSuffix(repoUrl, ".git")
}

func isGitSshUrl(repoUrl string) bool {
	return strings.HasPrefix(repoUrl, "git@")
}

func getScmProvider(hostname string) ScmProvider {
	switch hostname {
	case "github.com":
		return GitHub
	case "gitlab.com", "gitlab.zgtools.net":
		return GitLab
	case "bitbucket.org":
		return BitBucket
	default:
		return GitHub
	}
}

func getGitRemSshUrl(gitrRepo GitrRepo) string {
	return fmt.Sprintf("git@%s:%s.git", gitrRepo.HostName, gitrRepo.RepoPath)
}

func getGitRemHttpUrl(gitrRepo GitrRepo) string {
	return fmt.Sprintf("%s://%s/%s.git", gitrRepo.Protocol, gitrRepo.HostName, gitrRepo.RepoPath)
}

func getRepoPath(urlPath string) string {
	orgOrTeam := strings.Split(urlPath, "/")[0]
	repoName := strings.Split(urlPath, "/")[1]
	return fmt.Sprintf("%s/%s", orgOrTeam, repoName)
}

func getRepoName(repoPath string) string {
	return string(repoPath[strings.LastIndex(repoPath, "/")+1 : len(repoPath)])
}

func ParseGitRemoteUrl(repoUrl string) GitrRepo {
	var gitrRepo = GitrRepo{}
	if isGitSshUrl(repoUrl) {
		gitrRepo.Protocol = "https"
		gitrRepo.HostName = strings.Split(strings.Split(repoUrl, "@")[1], ":")[0]
	} else {
		gitrRepo.Protocol = strings.Split(repoUrl, "://")[0]
		gitrRepo.HostName = strings.Split(strings.Split(repoUrl, "://")[1], "/")[0]
	}
	gitrRepo.UrlPath = string(repoUrl[strings.Index(repoUrl, gitrRepo.HostName)+1+len(gitrRepo.HostName) : strings.Index(repoUrl, ".git")])
	gitrRepo.ScmProvider = getScmProvider(gitrRepo.HostName)
	gitrRepo.RepoPath = gitrRepo.UrlPath
	gitrRepo.RepoName = getRepoName(gitrRepo.RepoPath)
	gitrRepo.GitRemSshUrl = getGitRemSshUrl(gitrRepo)
	gitrRepo.GitRemHttpUrl = getGitRemHttpUrl(gitrRepo)
	return gitrRepo
}

func ParseBrowserUrl(browserUrl string) GitrRepo {
	var gitrRepo = GitrRepo{}
	gitrRepo.Protocol = strings.Split(browserUrl, "://")[0]
	gitrRepo.HostName = strings.Split(strings.Split(browserUrl, "://")[1], "/")[0]
	gitrRepo.UrlPath = string(browserUrl[strings.Index(browserUrl, gitrRepo.HostName)+1+len(gitrRepo.HostName):])
	gitrRepo.ScmProvider = getScmProvider(gitrRepo.HostName)
	gitrRepo.RepoPath = getRepoPath(gitrRepo.UrlPath)
	gitrRepo.RepoName = getRepoName(gitrRepo.RepoPath)
	if gitrRepo.ScmProvider != GitLab {
		gitrRepo.GitRemSshUrl = getGitRemSshUrl(gitrRepo)
		gitrRepo.GitRemHttpUrl = getGitRemHttpUrl(gitrRepo)
	} else {
		gitrRepo.GitRemSshUrl = ""
		gitrRepo.GitRemHttpUrl = ""
	}
	return gitrRepo
}

func ParseUrl(url string) GitrRepo {
	var gitrRepo GitrRepo
	if isGitUrl(url) {
		gitrRepo = ParseGitRemoteUrl(url)
	} else {
		gitrRepo = ParseBrowserUrl(url)
	}
	return gitrRepo
}