package lib

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"log"
	"os"
)

var err error

type GitRemoteScheme string

const (
	Ssh   GitRemoteScheme = "ssh"
	Https GitRemoteScheme = "https"
)

type GitrWeb struct {
	Url      string
	Scheme   GitRemoteScheme // http, https or ssh
	Provider ScmProvider
	Branch   string
}

func ScanRepo(dir string) *GitrWeb {
	gu := &GitUtil{}
	r := &GitrWeb{}
	repo := gu.GetGitRepo(dir)
	if repo != nil {
		remoteUrl := gu.GetGitRemoteUrl(repo)
		r.Url = remoteUrl
		r.Scheme = Https
		r.Branch = gu.GetGitBranch(repo)
		c := &GitrConfig{}
		r.Provider, err = c.GetScmProvider(GetHost(r.Url))
		if err != nil {
			log.Fatal(err)
		}
	}
	return r
}

func (r *GitrWeb) PrintInfo() {
	println("")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRow(table.Row{"remote", r.Url})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Provider", r.Provider})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Scheme", r.Scheme})
	t.AppendSeparator()
	t.AppendRow(table.Row{"host", GetHost(r.Url)})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repoPath", GetRepoPath(r.Url)})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repoName", GetRepoName(GetRepoPath(r.Url))})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Branch", r.Branch})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Url-home", r.GetWebUrl()})
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

func (r *GitrWeb) GetWebUrl() string {
	switch r.Provider {
	default:
		return fmt.Sprintf("%s://%s/%s", r.Scheme, GetHost(r.Url), GetRepoPath(r.Url))
	}
}

func (r *GitrWeb) GetRemUrl() string {
	switch r.Provider {
	case GitLab:
		return fmt.Sprintf("%s://%s/%s/-/tree/%s", r.Scheme, GetHost(r.Url), GetRepoPath(r.Url), r.Branch)
	default:
		return fmt.Sprintf("%s://%s/%s/tree/%s", r.Scheme, GetHost(r.Url), GetRepoPath(r.Url), r.Branch)
	}
}

func (r *GitrWeb) GetPrsUrl() string {
	switch r.Provider {
	case GitHub:
		return fmt.Sprintf("%s/pulls", r.GetWebUrl())
	case GitLab:
		return fmt.Sprintf("%s/-/merge_requests", r.GetWebUrl())
	case BitBucketDatacenter, BitBucketCloud:
		return fmt.Sprintf("%s/pull-requests", r.GetWebUrl())
	default:
		return ""
	}
}

func (r *GitrWeb) GetBranchesUrl() string {
	switch r.Provider {
	case GitLab:
		return fmt.Sprintf("%s/-/branches", r.GetWebUrl())
	default:
		return fmt.Sprintf("%s/branches", r.GetWebUrl())
	}
}

func (r *GitrWeb) GetCommitsUrl() string {
	switch r.Provider {
	case GitLab:
		return fmt.Sprintf("%s://%s/%s/-/commits/%s", r.Scheme, GetHost(r.Url), GetRepoPath(r.Url), r.Branch)
	default:
		return fmt.Sprintf("%s://%s/%s/commits/%s", r.Scheme, GetHost(r.Url), GetRepoPath(r.Url), r.Branch)
	}
}

func (r *GitrWeb) GetTagsUrl() string {
	switch r.Provider {
	case GitLab:
		return fmt.Sprintf("%s/-/tags", r.GetWebUrl())
	default:
		return fmt.Sprintf("%s/tags", r.GetWebUrl())
	}
}

func (r *GitrWeb) GetIssuesUrl() string {
	switch r.Provider {
	case BitBucketDatacenter, BitBucketCloud:
		return ""
	case GitLab:
		return fmt.Sprintf("%s/-/issues", r.GetWebUrl())
	default:
		return fmt.Sprintf("%s/issues", r.GetWebUrl())
	}
}

func (r *GitrWeb) GetReleasesUrl() string {
	switch r.Provider {
	case GitHub:
		return fmt.Sprintf("%s/releases", r.GetWebUrl())
	case GitLab:
		return fmt.Sprintf("%s/-/releases", r.GetWebUrl())
	default:
		return ""
	}
}

func (r *GitrWeb) GetPipelinesUrl() string {
	switch r.Provider {
	case GitLab:
		return fmt.Sprintf("%s/-/pipelines", r.GetWebUrl())
	case GitHub:
		return fmt.Sprintf("%s/actions", r.GetWebUrl())
	case BitBucketCloud:
		return fmt.Sprintf("%s/addon/pipelines/home", r.GetWebUrl())
	default:
		return ""
	}
}
