package internal

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/swarupdonepudi/gitr/v2/pkg"
	"os"
)

func PrintGitrWebInfo(scmSystem *pkg.ScmSystem, remoteUrl, branch string) {
	repoPath := getRepoPath(remoteUrl)
	repoName := getRepoName(remoteUrl)
	webUrl := GetWebUrl(scmSystem.Provider, remoteUrl)
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
	t.AppendRow(table.Row{"url-commits", GetCommitsUrl(scmSystem, repoPath, branch)})
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

func GetWebUrl(p pkg.ScmProvider, remoteUrl string) string {
	switch p {
	default:
		return fmt.Sprintf("https://%s/%s", getHost(remoteUrl), getRepoPath(remoteUrl))
	}
}

func GetRemUrl(p pkg.ScmProvider, webUrl, repoBranch string) string {
	switch p {
	case pkg.GitLab:
		return fmt.Sprintf("%s/-/tree/%s", webUrl, repoBranch)
	default:
		return fmt.Sprintf("%s/tree/%s", webUrl, repoBranch)
	}
}

func GetPrsUrl(p pkg.ScmProvider, webUrl string) string {
	switch p {
	case pkg.GitHub:
		return fmt.Sprintf("%s/pulls", webUrl)
	case pkg.GitLab:
		return fmt.Sprintf("%s/-/merge_requests", webUrl)
	case pkg.BitBucketDatacenter, pkg.BitBucketCloud:
		return fmt.Sprintf("%s/pull-requests", webUrl)
	default:
		return ""
	}
}

func GetBranchesUrl(p pkg.ScmProvider, webUrl string) string {
	switch p {
	case pkg.GitLab:
		return fmt.Sprintf("%s/-/branches", webUrl)
	default:
		return fmt.Sprintf("%s/branches", webUrl)
	}
}

func GetCommitsUrl(s *pkg.ScmSystem, webUrl, repoBranch string) string {
	switch s.Provider {
	case pkg.GitLab:
		return fmt.Sprintf("%s/-/commits/%s", webUrl, repoBranch)
	default:
		return fmt.Sprintf("%s/commits/%s", webUrl, repoBranch)
	}
}

func GetTagsUrl(p pkg.ScmProvider, webUrl string) string {
	switch p {
	case pkg.GitLab:
		return fmt.Sprintf("%s/-/tags", webUrl)
	default:
		return fmt.Sprintf("%s/tags", webUrl)
	}
}

func GetIssuesUrl(p pkg.ScmProvider, webUrl string) string {
	switch p {
	case pkg.BitBucketDatacenter, pkg.BitBucketCloud:
		return ""
	case pkg.GitLab:
		return fmt.Sprintf("%s/-/issues", webUrl)
	default:
		return fmt.Sprintf("%s/issues", webUrl)
	}
}

func GetReleasesUrl(p pkg.ScmProvider, webUrl string) string {
	switch p {
	case pkg.GitHub:
		return fmt.Sprintf("%s/releases", webUrl)
	case pkg.GitLab:
		return fmt.Sprintf("%s/-/releases", webUrl)
	default:
		return ""
	}
}

func GetPipelinesUrl(p pkg.ScmProvider, webUrl string) string {
	switch p {
	case pkg.GitLab:
		return fmt.Sprintf("%s/-/pipelines", webUrl)
	case pkg.GitHub:
		return fmt.Sprintf("%s/actions", webUrl)
	case pkg.BitBucketCloud:
		return fmt.Sprintf("%s/addon/pipelines/home", webUrl)
	default:
		return ""
	}
}
