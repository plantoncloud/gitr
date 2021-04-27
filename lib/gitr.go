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
	Url      string
	Scheme   GitRemoteScheme // http, https or ssh
	Provider ScmProvider
	Branch   string
}

func ScanRepo(dir string) *RemoteRepo {
	gu := &GitUtil{}
	r := &RemoteRepo{}
	repo := gu.GetGitRepo(dir)
	if repo != nil {
		remoteUrl := gu.GetGitRemoteUrl(repo)
		r.Url = remoteUrl
		r.Scheme = Https
		r.Branch = gu.GetGitBranch(repo)
		c := &GitrConfig{}
		r.Provider, err = c.GetScmProvider(r.getHost())
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
	t.AppendRow(table.Row{"remote", r.Url})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Provider", r.Provider})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Scheme", r.Scheme})
	t.AppendSeparator()
	t.AppendRow(table.Row{"host", r.getHost()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repoPath", r.getRepoPath()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repoName", r.getRepoName()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Branch", r.Branch})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Url-home", r.getWebUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Url-remote", r.GetRemUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Url-commits", r.GetCommitsUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Url-branches", r.GetBranchesUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Url-tags", r.GetTagsUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Url-releases", r.GetReleasesUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Url-pipelines", r.GetPipelinesUrl()})
	t.AppendSeparator()
	t.Render()
	println("")
}

func (r *RemoteRepo) getWebUrl() string {
	switch r.Provider {
	default:
		return fmt.Sprintf("%s://%s/%s", r.Scheme, r.getHost(), r.getRepoPath())
	}
}

func (r *RemoteRepo) GetRemUrl() string {
	switch r.Provider {
	case GitLab:
		return fmt.Sprintf("%s://%s/%s/-/tree/%s", r.Scheme, r.getHost(), r.getRepoPath(), r.Branch)
	default:
		return fmt.Sprintf("%s://%s/%s/tree/%s", r.Scheme, r.getHost(), r.getRepoPath(), r.Branch)
	}
}

func (r *RemoteRepo) GetPrsUrl() string {
	switch r.Provider {
	case GitHub:
		return fmt.Sprintf("%s/pulls", r.getWebUrl())
	case GitLab:
		return fmt.Sprintf("%s/-/merge_requests", r.getWebUrl())
	case BitBucketDatacenter, BitBucketCloud:
		return fmt.Sprintf("%s/pull-requests", r.getWebUrl())
	default:
		return ""
	}
}

func (r *RemoteRepo) GetBranchesUrl() string {
	switch r.Provider {
	case GitLab:
		return fmt.Sprintf("%s/-/branches", r.getWebUrl())
	default:
		return fmt.Sprintf("%s/branches", r.getWebUrl())
	}
}

func (r *RemoteRepo) GetCommitsUrl() string {
	switch r.Provider {
	case GitLab:
		return fmt.Sprintf("%s://%s/%s/-/commits/%s", r.Scheme, r.getHost(), r.getRepoPath(), r.Branch)
	default:
		return fmt.Sprintf("%s://%s/%s/commits/%s", r.Scheme, r.getHost(), r.getRepoPath(), r.Branch)
	}
}

func (r *RemoteRepo) GetTagsUrl() string {
	switch r.Provider {
	case GitLab:
		return fmt.Sprintf("%s/-/tags", r.getWebUrl())
	default:
		return fmt.Sprintf("%s/tags", r.getWebUrl())
	}
}

func (r *RemoteRepo) GetIssuesUrl() string {
	switch r.Provider {
	case BitBucketDatacenter, BitBucketCloud:
		return ""
	case GitLab:
		return fmt.Sprintf("%s/-/issues", r.getWebUrl())
	default:
		return fmt.Sprintf("%s/issues", r.getWebUrl())
	}
}

func (r *RemoteRepo) GetReleasesUrl() string {
	switch r.Provider {
	case GitHub:
		return fmt.Sprintf("%s/releases", r.getWebUrl())
	case GitLab:
		return fmt.Sprintf("%s/-/releases", r.getWebUrl())
	default:
		return ""
	}
}

func (r *RemoteRepo) GetPipelinesUrl() string {
	switch r.Provider {
	case GitLab:
		return fmt.Sprintf("%s/-/pipelines", r.getWebUrl())
	case GitHub:
		return fmt.Sprintf("%s/actions", r.getWebUrl())
	case BitBucketCloud:
		return fmt.Sprintf("%s/addon/pipelines/home", r.getWebUrl())
	default:
		return ""
	}
}

func (r *RemoteRepo) getHost() string {
	if r.Url != "" {
		if isGitSshUrl(r.Url) {
			if strings.HasPrefix(r.Url, "ssh://") {
				return strings.Split(strings.Split(r.Url, "@")[1], "/")[0]
			} else {
				return strings.Split(strings.Split(r.Url, "@")[1], ":")[0]
			}
		} else if isGitHttpUrlHasUsername(r.Url) {
			return strings.Split(strings.Split(r.Url, "@")[1], "/")[0]
		} else {
			return strings.Split(strings.Split(r.Url, "://")[1], "/")[0]
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
	return r.Url[strings.Index(r.Url, r.getHost())+1+len(r.getHost()) : strings.Index(r.Url, ".git")]
}
