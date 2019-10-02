package pkg

import (
	"fmt"
	"log"
	"os/user"
	"path/filepath"
	"strings"
	"regexp"
)

var err error

type InputUrlType string

const (
	Browser       InputUrlType = "browser"
	GitRemoteSsh  InputUrlType = "git-remote-ssh"
	GitRemoteHttp InputUrlType = "git-remote-http"
)

type GitrRepo struct {
	InputUrl      string
	InputUrlType  InputUrlType
	Protocol      string
	HostName      string
	UrlPath       string
	ScmProvider   ScmProvider
	levels        []string
	RepoName      string
	GitRemSshUrl  string
	GitRemHttpUrl string
}

func (c GitrRepo) ToString() string {
	return fmt.Sprintf("InputURL\t: \t%s" +
		"\nInputURLType\t: \t%s" +
		"\nScmProvider\t: \t%s"+
		"\nProtocol\t: \t%s"+
		"\nHostName\t: \t%s"+
		"\nUrlPath\t\t: \t%s"+
		"\nLevel 1\t\t: \t%s"+
		"\nLevel 2\t\t: \t%s"+
		"\nRepoName\t: \t%s"+
		"\nGitRemSshUrl\t: \t%s"+
		"\nGitRemHttpUrl\t: \t%s"+
		"\nWeb\t\t: \t%s"+
		"\nPRs\t\t: \t%s"+
		"\nBranches\t: \t%s"+
		"\nCommits\t\t: \t%s"+
		"\nIssues\t\t: \t%s"+
		"\nReleases\t: \t%s"+
		"\nPipelines\t: \t%s",
		c.InputUrl, c.InputUrlType, c.ScmProvider, c.Protocol, c.HostName,
		c.UrlPath, c.levels[0], c.levels[1], c.RepoName, c.GitRemSshUrl,
		c.GitRemHttpUrl, c.GetWebUrl(), c.GetPrsUrl(), c.GetBranchesUrl(),
		c.GetCommitsUrl(), c.GetIssuesUrl(), c.GetReleasesUrl(), c.GetPipelinesUrl())
}

func (gitrRepo GitrRepo) GetWebUrl() string {
	switch gitrRepo.ScmProvider {
	case BitBucketDatacenter:
		var project string
		if  gitrRepo.InputUrlType == GitRemoteHttp {
			project = gitrRepo.levels[1]
		} else {
			project = gitrRepo.levels[0]
		}
		return fmt.Sprintf("%s://%s/projects/%s/repos/%s", gitrRepo.Protocol, gitrRepo.HostName, project, gitrRepo.RepoName)
	default:
		return fmt.Sprintf("%s://%s/%s", gitrRepo.Protocol, gitrRepo.HostName, gitrRepo.UrlPath)
	}
}

func (c GitrRepo) GetPrsUrl() string {
	switch c.ScmProvider {
	case GitHub:
		return fmt.Sprintf("%s/pulls", c.GetWebUrl())
	case GitLab:
		return fmt.Sprintf("%s/merge_requests", c.GetWebUrl())
	case BitBucketDatacenter, BitBucketCloud:
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
	case BitBucketDatacenter, BitBucketCloud:
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
	case BitBucketCloud:
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
	return strings.HasPrefix(repoUrl, "ssh://") || strings.HasPrefix(repoUrl, "git@")
}

func isGitHttpUrlHasUsername(repoUrl string) bool {
	matched, err := regexp.MatchString("https*:\\/\\/.*@+.*",repoUrl)
	if err != nil {
		println(err.Error())
		return false
	} else {
		return matched
	}
}

func getGitRemSshUrl(gitrRepo GitrRepo) string {
	switch gitrRepo.ScmProvider {
	case BitBucketDatacenter:
		return fmt.Sprintf("ssh://git@%s/%s/%s.git", gitrRepo.HostName, gitrRepo.levels[1], gitrRepo.RepoName)
	default:
		return fmt.Sprintf("git@%s:%s/%s.git", gitrRepo.HostName, gitrRepo.levels[0], gitrRepo.RepoName)
	}

}

func getGitRemHttpUrl(gitrRepo GitrRepo) string {
	switch gitrRepo.ScmProvider {
	case BitBucketDatacenter:
		return fmt.Sprintf("%s://%s/scm/%s/%s.git", gitrRepo.Protocol, gitrRepo.HostName, gitrRepo.levels[1], gitrRepo.RepoName)
	default:
		return fmt.Sprintf("%s://%s/%s/%s.git", gitrRepo.Protocol, gitrRepo.HostName, gitrRepo.levels[0], gitrRepo.RepoName)
	}
}

func getHostName(url string) string {
	if isGitSshUrl(url) {
		if strings.HasPrefix(url, "ssh://") {
			return strings.Split(strings.Split(url, "@")[1], "/")[0]
		} else {
			return strings.Split(strings.Split(url, "@")[1], ":")[0]
		}
	} else if isGitHttpUrlHasUsername(url) {
		return strings.Split(strings.Split(url, "@")[1], "/")[0]
	} else {
		return strings.Split(strings.Split(url, "://")[1], "/")[0]
	}
}

func getLevels(urlPath string) []string {
	levels := strings.Split(urlPath, "/")
	return levels
}

func getRepoName(inputUrlType InputUrlType, scmProvider ScmProvider, levels []string) string {
	if scmProvider == BitBucketDatacenter {
		if inputUrlType == GitRemoteHttp {
			return levels[2]
		} else if inputUrlType == Browser {
			return levels[3]
		} else {
			return levels[1]
		}
	} else {
		return levels[1]
	}
}

func ParseGitRemoteUrl(repoUrl string) GitrRepo {
	var gitrRepo = GitrRepo{}
	gitrRepo.InputUrl = repoUrl
	if isGitSshUrl(repoUrl) {
		gitrRepo.InputUrlType = GitRemoteSsh
		gitrRepo.Protocol = "https"
	} else {
		gitrRepo.InputUrlType = GitRemoteHttp
		gitrRepo.Protocol = strings.Split(repoUrl, "://")[0]
	}
	gitrRepo.HostName = getHostName(repoUrl)

	gitrRepo.UrlPath = string(repoUrl[strings.Index(repoUrl, gitrRepo.HostName)+1+len(gitrRepo.HostName) : strings.Index(repoUrl, ".git")])

	gitrRepo.ScmProvider, err = getScmProvider(gitrRepo.HostName)

	if err != nil {
		log.Fatal(err)
	}

	gitrRepo.levels = getLevels(gitrRepo.UrlPath)
	gitrRepo.RepoName = getRepoName(gitrRepo.InputUrlType, gitrRepo.ScmProvider, gitrRepo.levels)
	gitrRepo.GitRemSshUrl = getGitRemSshUrl(gitrRepo)
	gitrRepo.GitRemHttpUrl = getGitRemHttpUrl(gitrRepo)
	return gitrRepo
}

func ParseBrowserUrl(browserUrl string) GitrRepo {
	var gitrRepo = GitrRepo{}
	gitrRepo.InputUrl = browserUrl
	gitrRepo.InputUrlType = Browser
	gitrRepo.Protocol = strings.Split(browserUrl, "://")[0]
	gitrRepo.HostName = strings.Split(strings.Split(browserUrl, "://")[1], "/")[0]
	gitrRepo.UrlPath = string(browserUrl[strings.Index(browserUrl, gitrRepo.HostName)+1+len(gitrRepo.HostName):])

	gitrRepo.ScmProvider, err = getScmProvider(gitrRepo.HostName)

	if err != nil {
		log.Fatal(err)
	}

	gitrRepo.levels = getLevels(gitrRepo.UrlPath)
	gitrRepo.RepoName = getRepoName(gitrRepo.InputUrlType, gitrRepo.ScmProvider, gitrRepo.levels)
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
