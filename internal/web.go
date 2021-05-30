package internal

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	"github.com/swarupdonepudi/gitr/v2/pkg/url"
	"os"
)

func PrintGitrWebInfo(scmSystem *config.ScmSystem, remoteUrl, branch string) {
	repoPath := url.GetRepoPath(remoteUrl)
	repoName := url.GetRepoName(remoteUrl)
	webUrl := GetWebUrl(scmSystem.Provider, scmSystem.Scheme, remoteUrl)
	println("")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRow(table.Row{"remote", remoteUrl})
	t.AppendSeparator()
	t.AppendRow(table.Row{"provider", scmSystem.Provider})
	t.AppendSeparator()
	t.AppendRow(table.Row{"host", scmSystem.Hostname})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repo-path", repoPath})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repo-name", repoName})
	t.AppendSeparator()
	t.AppendRow(table.Row{"branch", branch})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-web", webUrl})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-remote", GetRemUrl(scmSystem.Provider, webUrl, branch)})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-commits", GetCommitsUrl(scmSystem.Provider, repoPath, branch)})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-branches", GetBranchesUrl(scmSystem.Provider, webUrl)})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-tags", GetTagsUrl(scmSystem.Provider, webUrl)})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-releases", GetReleasesUrl(scmSystem.Provider, webUrl)})
	t.AppendSeparator()
	t.AppendRow(table.Row{"url-pipelines", GetPipelinesUrl(scmSystem.Provider, webUrl)})
	t.AppendSeparator()
	t.Render()
	println("")
}

func GetWebUrl(p config.ScmProvider, scheme config.HttpScheme, remoteUrl string) string {
	switch p {
	default:
		return fmt.Sprintf("%s://%s/%s", scheme, url.GetHost(remoteUrl), url.GetRepoPath(remoteUrl))
	}
}

func GetRemUrl(p config.ScmProvider, webUrl, repoBranch string) string {
	switch p {
	case config.GitLab:
		return fmt.Sprintf("%s/-/tree/%s", webUrl, repoBranch)
	default:
		return fmt.Sprintf("%s/tree/%s", webUrl, repoBranch)
	}
}

func GetPrsUrl(p config.ScmProvider, webUrl string) string {
	switch p {
	case config.GitHub:
		return fmt.Sprintf("%s/pulls", webUrl)
	case config.GitLab:
		return fmt.Sprintf("%s/-/merge_requests", webUrl)
	case config.BitBucketDatacenter, config.BitBucketCloud:
		return fmt.Sprintf("%s/pull-requests", webUrl)
	default:
		return ""
	}
}

func GetBranchesUrl(p config.ScmProvider, webUrl string) string {
	switch p {
	case config.GitLab:
		return fmt.Sprintf("%s/-/branches", webUrl)
	default:
		return fmt.Sprintf("%s/branches", webUrl)
	}
}

func GetCommitsUrl(p config.ScmProvider, webUrl, repoBranch string) string {
	switch p {
	case config.GitLab:
		return fmt.Sprintf("%s/-/commits/%s", webUrl, repoBranch)
	default:
		return fmt.Sprintf("%s/commits/%s", webUrl, repoBranch)
	}
}

func GetTagsUrl(p config.ScmProvider, webUrl string) string {
	switch p {
	case config.GitLab:
		return fmt.Sprintf("%s/-/tags", webUrl)
	default:
		return fmt.Sprintf("%s/tags", webUrl)
	}
}

func GetIssuesUrl(p config.ScmProvider, webUrl string) string {
	switch p {
	case config.BitBucketDatacenter, config.BitBucketCloud:
		return ""
	case config.GitLab:
		return fmt.Sprintf("%s/-/issues", webUrl)
	default:
		return fmt.Sprintf("%s/issues", webUrl)
	}
}

func GetReleasesUrl(p config.ScmProvider, webUrl string) string {
	switch p {
	case config.GitHub:
		return fmt.Sprintf("%s/releases", webUrl)
	case config.GitLab:
		return fmt.Sprintf("%s/-/releases", webUrl)
	default:
		return ""
	}
}

func GetPipelinesUrl(p config.ScmProvider, webUrl string) string {
	switch p {
	case config.GitLab:
		return fmt.Sprintf("%s/-/pipelines", webUrl)
	case config.GitHub:
		return fmt.Sprintf("%s/actions", webUrl)
	case config.BitBucketCloud:
		return fmt.Sprintf("%s/addon/pipelines/home", webUrl)
	default:
		return ""
	}
}
