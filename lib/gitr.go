package lib

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"log"
	"os"
	"strings"
)

var err error

type GitRemoteScheme string

const (
	Ssh   GitRemoteScheme = "ssh"
	Https GitRemoteScheme = "https"
)

type RemoteRepo struct {
	url      string
	scheme   GitRemoteScheme // http, https or ssh
	provider ScmProvider
	branch   string
}

func ScanRepo(dir string) *RemoteRepo {
	gu := &GitUtil{}
	r := &RemoteRepo{}
	repo := gu.GetGitRepo(dir)
	if repo != nil {
		remoteUrl := gu.GetGitRemoteUrl(repo)
		r.url = remoteUrl
		r.scheme = Https
		r.branch = gu.GetGitBranch(repo)
		c := &GitrConfig{}
		r.provider, err = c.GetScmProvider(r.getHost())
		if err != nil {
			log.Fatal(err)
		}
	}
	return r
}

func (r *RemoteRepo) PrintInfo() {
	println("")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRow(table.Row{"remote", r.url})
	t.AppendSeparator()
	t.AppendRow(table.Row{"provider", r.provider})
	t.AppendSeparator()
	t.AppendRow(table.Row{"scheme", r.scheme})
	t.AppendSeparator()
	t.AppendRow(table.Row{"host", r.getHost()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repoPath", r.getRepoPath()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repoName", r.getRepoName()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"branch", r.branch})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-home", r.getWebUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-remote", r.GetRemUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-commits", r.GetCommitsUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-tags", r.GetTagsUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-releases", r.GetReleasesUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-pipelines", r.GetPipelinesUrl()})
	t.AppendSeparator()
	t.Render()
	println("")
}

func (r *RemoteRepo) getWebUrl() string {
	switch r.provider {
	default:
		return fmt.Sprintf("%s://%s/%s", r.scheme, r.getHost(), r.getRepoPath())
	}
}

func (r *RemoteRepo) GetRemUrl() string {
	switch r.provider {
	case GitLab:
		return fmt.Sprintf("%s://%s/%s/-/tree/%s", r.scheme, r.getHost(), r.getRepoPath(), r.branch)
	default:
		return fmt.Sprintf("%s://%s/%s/tree/%s", r.scheme, r.getHost(), r.getRepoPath(), r.branch)
	}
}

func (r *RemoteRepo) GetPrsUrl() string {
	switch r.provider {
	case GitHub:
		return fmt.Sprintf("%s/pulls", r.getWebUrl())
	case GitLab:
		return fmt.Sprintf("%s/merge_requests", r.getWebUrl())
	case BitBucketDatacenter, BitBucketCloud:
		return fmt.Sprintf("%s/pull-requests", r.getWebUrl())
	default:
		return ""
	}
}

func (r *RemoteRepo) GetBranchesUrl() string {
	switch r.provider {
	case GitLab:
		return fmt.Sprintf("%s/-/branches", r.getWebUrl())
	default:
		return fmt.Sprintf("%s/branches", r.getWebUrl())
	}
}

func (r *RemoteRepo) GetCommitsUrl() string {
	switch r.provider {
	case GitLab:
		return fmt.Sprintf("%s://%s/%s/-/commits/%s", r.scheme, r.getHost(), r.getRepoPath(), r.branch)
	default:
		return fmt.Sprintf("%s://%s/%s/commits/%s", r.scheme, r.getHost(), r.getRepoPath(), r.branch)
	}
}

func (r *RemoteRepo) GetTagsUrl() string {
	switch r.provider {
	case GitLab:
		return fmt.Sprintf("%s/-/tags", r.getWebUrl())
	default:
		return fmt.Sprintf("%s/tags", r.getWebUrl())
	}
}

func (r *RemoteRepo) GetIssuesUrl() string {
	switch r.provider {
	case BitBucketDatacenter, BitBucketCloud:
		return ""
	default:
		return fmt.Sprintf("%s/issues", r.getWebUrl())
	}
}

func (r *RemoteRepo) GetReleasesUrl() string {
	switch r.provider {
	case GitHub:
		return fmt.Sprintf("%s/releases", r.getWebUrl())
	default:
		return ""
	}
}

func (r *RemoteRepo) GetPipelinesUrl() string {
	switch r.provider {
	case GitLab:
		return fmt.Sprintf("%s/pipelines", r.getWebUrl())
	case GitHub:
		return fmt.Sprintf("%s/actions", r.getWebUrl())
	case BitBucketCloud:
		return fmt.Sprintf("%s/addon/pipelines/home", r.getWebUrl())
	default:
		return ""
	}
}

func (r *RemoteRepo) getHost() string {
	if r.url != "" {
		if isGitSshUrl(r.url) {
			if strings.HasPrefix(r.url, "ssh://") {
				return strings.Split(strings.Split(r.url, "@")[1], "/")[0]
			} else {
				return strings.Split(strings.Split(r.url, "@")[1], ":")[0]
			}
		} else if isGitHttpUrlHasUsername(r.url) {
			return strings.Split(strings.Split(r.url, "@")[1], "/")[0]
		} else {
			return strings.Split(strings.Split(r.url, "://")[1], "/")[0]
		}
	} else {
		return ""
	}
}

func (r *RemoteRepo) getRepoName() string {
	if r.getRepoPath() != "" {
		levels := strings.Split(r.getRepoPath(), "/")
		if len(levels) < 2 {
			log.Fatal("failed to parse repo name")
		}
		return levels[1]
	} else {
		return ""
	}
}

func (r *RemoteRepo) getRepoPath() string {
	return r.url[strings.Index(r.url, r.getHost())+1+len(r.getHost()) : strings.Index(r.url, ".git")]
}
