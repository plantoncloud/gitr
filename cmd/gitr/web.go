package gitr

import (
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/gitr/v2/internal/cli"
	"github.com/swarupdonepudi/gitr/v2/internal/git"
	"github.com/swarupdonepudi/gitr/v2/internal/web"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	"github.com/swarupdonepudi/gitr/v2/pkg/url"
	"log"
	"os"
)

type WebCmd string

const (
	Branches  WebCmd = "branches"
	Prs       WebCmd = "prs"
	Commits   WebCmd = "commits"
	Issues    WebCmd = "issues"
	Tags      WebCmd = "tags"
	Releases  WebCmd = "releases"
	Pipelines WebCmd = "pipelines"
	Web       WebCmd = "web"
	Rem       WebCmd = "rem"
)

var branchesCmd = &cobra.Command{
	Use:   string(Branches),
	Short: "open branches of the repo in the browser",
	Run:   webHandler,
}

var webCmd = &cobra.Command{
	Use:   string(Web),
	Short: "open home page of the repo in the browser",
	Run:   webHandler,
}

var tagsCmd = &cobra.Command{
	Use:   string(Tags),
	Short: "open tags of the repo in the browser",
	Run:   webHandler,
}

var remCmd = &cobra.Command{
	Use:   string(Rem),
	Short: "open local checkout branch of the repo in the browser",
	Run:   webHandler,
}

var releasesCmd = &cobra.Command{
	Use:   string(Releases),
	Short: "open releases of the repo in the browser",
	Run:   webHandler,
}

var prsCmd = &cobra.Command{
	Use:   string(Prs),
	Short: "open prs/mrs of the repo in the browser",
	Run:   webHandler,
}

var pipelinesCmd = &cobra.Command{
	Use:     string(Pipelines),
	Short:   "open pipelines/actions of the repo in the browser",
	Aliases: []string{"pipe"},
	Run:     webHandler,
}

var issuesCmd = &cobra.Command{
	Use:   string(Issues),
	Short: "open issues of the repo in the browser",
	Run:   webHandler,
}

var commitsCmd = &cobra.Command{
	Use:   string(Commits),
	Short: "open commits of the local branch of repo in the browser",
	Run:   webHandler,
}

func webHandler(cmd *cobra.Command, args []string) {
	dry, err := cmd.InheritedFlags().GetBool(string(cli.Dry))
	cli.HandleFlagErr(err, cli.Dry)
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current dir. err: %v", err)
	}
	r := git.GetGitRepo(pwd)
	remoteUrl := git.GetGitRemoteUrl(r)
	branch := git.GetGitBranch(r)
	cfg, err := config.NewGitrConfig()
	if err != nil {
		log.Fatalf("failed to get gitr config. err: %v", err)
	}
	s, err := config.GetScmHost(cfg, url.GetHostname(remoteUrl))
	if err != nil {
		log.Fatal(err)
	}

	repoPath := url.GetRepoPath(remoteUrl, s.Hostname, s.Provider)
	repoName := url.GetRepoName(repoPath)
	webUrl := web.GetWebUrl(s.Provider, s.Scheme, s.Hostname, repoPath)

	if dry {
		web.PrintGitrWebInfo(s.Provider, s.Hostname, remoteUrl, webUrl, repoPath, repoName, branch)
		return
	}

	switch WebCmd(cmd.Name()) {
	case Branches:
		url.OpenInBrowser(web.GetBranchesUrl(s.Provider, webUrl))
	case Prs:
		url.OpenInBrowser(web.GetPrsUrl(s.Provider, webUrl))
	case Commits:
		url.OpenInBrowser(web.GetCommitsUrl(s.Provider, webUrl, branch))
	case Issues:
		url.OpenInBrowser(web.GetIssuesUrl(s.Provider, webUrl))
	case Tags:
		url.OpenInBrowser(web.GetTagsUrl(s.Provider, webUrl))
	case Releases:
		url.OpenInBrowser(web.GetReleasesUrl(s.Provider, webUrl))
	case Pipelines:
		url.OpenInBrowser(web.GetPipelinesUrl(s.Provider, webUrl))
	case Web:
		url.OpenInBrowser(webUrl)
	case Rem:
		url.OpenInBrowser(web.GetRemUrl(s.Provider, webUrl, branch))
	default:
		log.Fatal("unknown web command")
	}
}
