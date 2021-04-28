package lib

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"log"
	"os"
)

var err error

type GitrWeb struct {
	Url      string
	Scheme   string //TODO: is this really needed?
	Provider ScmProvider
	Branch   string
}

func ScanRepo(dir string) *GitrWeb {
	gu := &GitUtil{}
	gtu := &GitrUtil{}
	r := &GitrWeb{}
	repo := gu.GetGitRepo(dir)
	if repo != nil {
		remoteUrl := gu.GetGitRemoteUrl(repo)
		r.Url = remoteUrl
		r.Branch = gu.GetGitBranch(repo)
		r.Scheme = "https"
		c := &GitrConfig{}
		scmSystem, err := c.GetScmSystem(gtu.GetHost(r.Url))
		if err != nil {
			log.Fatal(err)
		}
		r.Provider = scmSystem.Provider
	}
	return r
}

func (r *GitrWeb) PrintInfo() {
	gru := &GitrUtil{}
	println("")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRow(table.Row{"remote", r.Url})
	t.AppendSeparator()
	t.AppendRow(table.Row{"provider", r.Provider})
	t.AppendSeparator()
	t.AppendRow(table.Row{"host", gru.GetHost(r.Url)})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repo-path", gru.GetRepoPath(r.Url)})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repo-name", gru.GetRepoName(gru.GetRepoPath(r.Url))})
	t.AppendSeparator()
	t.AppendRow(table.Row{"branch", r.Branch})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-web", r.GetWebUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-remote", r.GetRemUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-commits", r.GetCommitsUrl()})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-branches", r.GetBranchesUrl()})
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

func (r *GitrWeb) GetWebUrl() string {
	gru := &GitrUtil{}
	switch r.Provider {
	default:
		return fmt.Sprintf("%s://%s/%s", r.Scheme, gru.GetHost(r.Url), gru.GetRepoPath(r.Url))
	}
}

func (r *GitrWeb) GetRemUrl() string {
	gru := &GitrUtil{}
	switch r.Provider {
	case GitLab:
		return fmt.Sprintf("%s://%s/%s/-/tree/%s", r.Scheme, gru.GetHost(r.Url), gru.GetRepoPath(r.Url), r.Branch)
	default:
		return fmt.Sprintf("%s://%s/%s/tree/%s", r.Scheme, gru.GetHost(r.Url), gru.GetRepoPath(r.Url), r.Branch)
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
	gru := &GitrUtil{}
	switch r.Provider {
	case GitLab:
		return fmt.Sprintf("%s://%s/%s/-/commits/%s", r.Scheme, gru.GetHost(r.Url), gru.GetRepoPath(r.Url), r.Branch)
	default:
		return fmt.Sprintf("%s://%s/%s/commits/%s", r.Scheme, gru.GetHost(r.Url), gru.GetRepoPath(r.Url), r.Branch)
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
