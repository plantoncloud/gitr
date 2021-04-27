package lib

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"log"
	"os"
	"regexp"
	"strings"
)

var err error

type GitOriginScheme string

const (
	Ssh  GitOriginScheme = "ssh"
	Http GitOriginScheme = "http"
)

type GitOrigin struct {
	url      string
	scheme   GitOriginScheme // http, https or ssh
	host     string
	repoPath string
	provider ScmProvider
	levels   []string
	repoName string
}

func (o *GitOrigin) ScanOrigin() {
	pwd, _ := os.Getwd()
	repo := GetGitRepo(pwd)
	if repo != nil {
		remoteUrl := GetGitRemoteUrl(repo)
		o.parseGitRemoteUrl(remoteUrl)
		o.PrintTable()
	}
}

func (o *GitOrigin) PrintTable() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	println("")
	t.AppendRow(table.Row{"url", o.url})
	t.AppendSeparator()
	t.AppendRow(table.Row{"provider", o.provider})
	t.AppendSeparator()
	t.AppendRow(table.Row{"scheme", o.scheme})
	t.AppendSeparator()
	t.AppendRow(table.Row{"host", o.host})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repoPath", o.repoPath})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repoName", o.repoName})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-remote", o.GetWebUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-commits", o.GetCommitsUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-tags", o.GetTagsUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-releases", o.GetReleasesUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-pipelines", o.GetPipelinesUrl()})
	t.AppendSeparator()
	t.Render()
}

func (o *GitOrigin) GetWebUrl() string {
	switch o.provider {
	case BitBucketDatacenter:
		var project string
		if o.scheme == Http {
			project = o.levels[1]
		} else {
			project = o.levels[0]
		}
		return fmt.Sprintf("%s://%s/projects/%s/repos/%s", o.scheme, o.host, project, o.repoName)
	default:
		return fmt.Sprintf("%s://%s/%s", o.scheme, o.host, o.repoPath)
	}
}

func (o *GitOrigin) GetPrsUrl() string {
	switch o.provider {
	case GitHub:
		return fmt.Sprintf("%s/pulls", o.GetWebUrl())
	case GitLab:
		return fmt.Sprintf("%s/merge_requests", o.GetWebUrl())
	case BitBucketDatacenter, BitBucketCloud:
		return fmt.Sprintf("%s/pull-requests", o.GetWebUrl())
	default:
		return ""
	}
}

func (o *GitOrigin) GetBranchesUrl() string {
	switch o.provider {
	case GitLab:
		return fmt.Sprintf("%s/-/branches", o.GetWebUrl())
	default:
		return fmt.Sprintf("%s/branches", o.GetWebUrl())
	}
}

func (o *GitOrigin) GetCommitsUrl() string {
	switch o.provider {
	case GitLab:
		return fmt.Sprintf("%s/-/commits", o.GetWebUrl())
	default:
		return fmt.Sprintf("%s/commits", o.GetWebUrl())
	}
}

func (o *GitOrigin) GetTagsUrl() string {
	switch o.provider {
	case GitLab:
		return fmt.Sprintf("%s/-/tags", o.GetWebUrl())
	default:
		return fmt.Sprintf("%s/tags", o.GetWebUrl())
	}
}

func (o *GitOrigin) GetIssuesUrl() string {
	switch o.provider {
	case BitBucketDatacenter, BitBucketCloud:
		return ""
	default:
		return fmt.Sprintf("%s/issues", o.GetWebUrl())
	}
}

func (o *GitOrigin) GetReleasesUrl() string {
	switch o.provider {
	case GitHub:
		return fmt.Sprintf("%s/releases", o.GetWebUrl())
	default:
		return ""
	}
}

func (o *GitOrigin) GetPipelinesUrl() string {
	switch o.provider {
	case GitLab:
		return fmt.Sprintf("%s/pipelines", o.GetWebUrl())
	case GitHub:
		return fmt.Sprintf("%s/actions", o.GetWebUrl())
	case BitBucketCloud:
		return fmt.Sprintf("%s/addon/pipelines/home", o.GetWebUrl())
	default:
		return ""
	}
}

func isGitSshUrl(repoUrl string) bool {
	return strings.HasPrefix(repoUrl, "ssh://") || strings.HasPrefix(repoUrl, "git@")
}

func isGitHttpUrlHasUsername(repoUrl string) bool {
	matched, err := regexp.MatchString("https*:\\/\\/.*@+.*", repoUrl)
	if err != nil {
		println(err.Error())
		return false
	} else {
		return matched
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

func getRepoName(scheme GitOriginScheme, scmProvider ScmProvider, levels []string) string {
	if scmProvider == BitBucketDatacenter {
		if scheme == Http {
			return levels[2]
		} else {
			return levels[1]
		}
	} else {
		return levels[1]
	}
}

func (o *GitOrigin) parseGitRemoteUrl(repoUrl string) {
	o.url = repoUrl
	o.scheme = Http
	o.host = getHostName(repoUrl)

	o.repoPath = repoUrl[strings.Index(repoUrl, o.host)+1+len(o.host) : strings.Index(repoUrl, ".git")]

	c := &GitrConfig{}
	o.provider, err = c.GetScmProvider(o.host)

	if err != nil {
		log.Fatal(err)
	}

	o.levels = getLevels(o.repoPath)
	o.repoName = getRepoName(o.scheme, o.provider, o.levels)
}
