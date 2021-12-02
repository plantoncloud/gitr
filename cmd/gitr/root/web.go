package root

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/gitr/v2/internal/cli"
	"github.com/swarupdonepudi/gitr/v2/internal/web"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	"github.com/swarupdonepudi/gitr/v2/pkg/git"
	"github.com/swarupdonepudi/gitr/v2/pkg/url"
	"os"
)

type WebCmdName string

const (
	branches  WebCmdName = "branches"
	prs       WebCmdName = "prs"
	commits   WebCmdName = "commits"
	issues    WebCmdName = "issues"
	tags      WebCmdName = "tags"
	releases  WebCmdName = "releases"
	pipelines WebCmdName = "pipelines"
	webHome   WebCmdName = "web"
	rem       WebCmdName = "rem"
)

var BranchesCmd = &cobra.Command{
	Use:   string(branches),
	Short: "open branches of the repo in the browser",
	Run:   webHandler,
}

var WebCmd = &cobra.Command{
	Use:   string(webHome),
	Short: "open home page of the repo in the browser",
	Run:   webHandler,
}

var TagsCmd = &cobra.Command{
	Use:   string(tags),
	Short: "open tags of the repo in the browser",
	Run:   webHandler,
}

var RemCmd = &cobra.Command{
	Use:   string(rem),
	Short: "open local checkout branch of the repo in the browser",
	Run:   webHandler,
}

var ReleasesCmd = &cobra.Command{
	Use:   string(releases),
	Short: "open releases of the repo in the browser",
	Run:   webHandler,
}

var PrsCmd = &cobra.Command{
	Use:   string(prs),
	Short: "open prs/mrs of the repo in the browser",
	Run:   webHandler,
}

var PipelinesCmd = &cobra.Command{
	Use:     string(pipelines),
	Short:   "open pipelines/actions of the repo in the browser",
	Aliases: []string{"pipe"},
	Run:     webHandler,
}

var IssuesCmd = &cobra.Command{
	Use:   string(issues),
	Short: "open issues of the repo in the browser",
	Run:   webHandler,
}

var CommitsCmd = &cobra.Command{
	Use:   string(commits),
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
		log.Fatalf("failed to get scm host for %s url. err: %v", remoteUrl, err)
	}

	repoPath := url.GetRepoPath(remoteUrl, s.Hostname, s.Provider)
	repoName := url.GetRepoName(repoPath)
	webUrl := web.GetWebUrl(s.Provider, s.Scheme, s.Hostname, repoPath)

	if dry {
		web.PrintGitrWebInfo(s.Provider, s.Hostname, remoteUrl, webUrl, repoPath, repoName, branch)
		return
	}

	switch WebCmdName(cmd.Name()) {
	case branches:
		url.OpenInBrowser(web.GetBranchesUrl(s.Provider, webUrl))
	case prs:
		url.OpenInBrowser(web.GetPrsUrl(s.Provider, webUrl))
	case commits:
		url.OpenInBrowser(web.GetCommitsUrl(s.Provider, webUrl, branch))
	case issues:
		url.OpenInBrowser(web.GetIssuesUrl(s.Provider, webUrl))
	case tags:
		url.OpenInBrowser(web.GetTagsUrl(s.Provider, webUrl))
	case releases:
		url.OpenInBrowser(web.GetReleasesUrl(s.Provider, webUrl))
	case pipelines:
		url.OpenInBrowser(web.GetPipelinesUrl(s.Provider, webUrl))
	case webHome:
		url.OpenInBrowser(webUrl)
	case rem:
		url.OpenInBrowser(web.GetRemUrl(s.Provider, webUrl, branch))
	default:
		log.Fatal("unknown web command")
	}
}
