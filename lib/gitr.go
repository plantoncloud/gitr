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

type GitRepo struct {
	origin   string
	scheme   GitOriginScheme // http, https or ssh
	host     string
	repoPath string
	provider ScmProvider
	levels   []string
	repoName string
	branch   string
}

func (o *GitRepo) Scan() {
	pwd, _ := os.Getwd()
	gu := &GitUtil{}
	repo := gu.GetGitRepo(pwd)
	if repo != nil {
		remoteUrl := gu.GetGitRemoteUrl(repo)
		o.branch = gu.GetGitBranch(repo)
		o.parseGitRemoteUrl(remoteUrl)
		o.PrintTable()
	}
}

func (o *GitRepo) PrintTable() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	println("")
	t.AppendRow(table.Row{"origin", o.origin})
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
	t.AppendRow(table.Row{"branch", o.branch})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-remote", o.getWebUrl()})
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

func (o *GitRepo) getWebUrl() string {
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

func (o *GitRepo) GetRemUrl() string {
	switch o.provider {
	case BitBucketDatacenter:
		var project string
		if o.scheme == Http {
			project = o.levels[1]
		} else {
			project = o.levels[0]
		}
		return fmt.Sprintf("%s://%s/projects/%s/repos/%s", o.scheme, o.host, project, o.repoName)
	case GitLab:
		return fmt.Sprintf("%s://%s/%s/-/tree/%s", o.scheme, o.host, o.repoPath, o.branch)
	default:
		return fmt.Sprintf("%s://%s/%s/tree/%s", o.scheme, o.host, o.repoPath, o.branch)

	}
}

func (o *GitRepo) GetPrsUrl() string {
	switch o.provider {
	case GitHub:
		return fmt.Sprintf("%s/pulls", o.getWebUrl())
	case GitLab:
		return fmt.Sprintf("%s/merge_requests", o.getWebUrl())
	case BitBucketDatacenter, BitBucketCloud:
		return fmt.Sprintf("%s/pull-requests", o.getWebUrl())
	default:
		return ""
	}
}

func (o *GitRepo) GetBranchesUrl() string {
	switch o.provider {
	case GitLab:
		return fmt.Sprintf("%s/-/branches", o.getWebUrl())
	default:
		return fmt.Sprintf("%s/branches", o.getWebUrl())
	}
}

func (o *GitRepo) GetCommitsUrl() string {
	switch o.provider {
	case GitLab:
		return fmt.Sprintf("%s/-/commits", o.getWebUrl())
	default:
		return fmt.Sprintf("%s/commits", o.getWebUrl())
	}
}

func (o *GitRepo) GetTagsUrl() string {
	switch o.provider {
	case GitLab:
		return fmt.Sprintf("%s/-/tags", o.getWebUrl())
	default:
		return fmt.Sprintf("%s/tags", o.getWebUrl())
	}
}

func (o *GitRepo) GetIssuesUrl() string {
	switch o.provider {
	case BitBucketDatacenter, BitBucketCloud:
		return ""
	default:
		return fmt.Sprintf("%s/issues", o.getWebUrl())
	}
}

func (o *GitRepo) GetReleasesUrl() string {
	switch o.provider {
	case GitHub:
		return fmt.Sprintf("%s/releases", o.getWebUrl())
	default:
		return ""
	}
}

func (o *GitRepo) GetPipelinesUrl() string {
	switch o.provider {
	case GitLab:
		return fmt.Sprintf("%s/pipelines", o.getWebUrl())
	case GitHub:
		return fmt.Sprintf("%s/actions", o.getWebUrl())
	case BitBucketCloud:
		return fmt.Sprintf("%s/addon/pipelines/home", o.getWebUrl())
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

func (o *GitRepo) parseGitRemoteUrl(repoUrl string) {
	o.origin = repoUrl
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
