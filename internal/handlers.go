package internal

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/swarupdonepudi/gitr/v2/pkg/clone"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	"github.com/swarupdonepudi/gitr/v2/pkg/file"
	"github.com/swarupdonepudi/gitr/v2/pkg/url"
	"log"
)

type GitrFlag string

type CmdName string

const (
	Dry       GitrFlag = "dry"
	CreDir    GitrFlag = "create-dir"
	Branches  CmdName  = "branches"
	Prs       CmdName  = "prs"
	Commits   CmdName  = "commits"
	Issues    CmdName  = "issues"
	Tags      CmdName  = "tags"
	Releases  CmdName  = "releases"
	Pipelines CmdName  = "pipelines"
	Web       CmdName  = "web"
	Rem       CmdName  = "rem"
	Clone     CmdName  = "clone"
)

func CloneHandler(cmd *cobra.Command, args []string) {
	inputUrl := args[0]
	creDir := viper.GetBool(string(CreDir))
	cfg := config.GetGitrConfig()
	if viper.GetBool(string(Dry)) {
		PrintGitrCloneInfo(inputUrl, creDir || cfg.Clone.AlwaysCreDir, cfg)
		return
	}
	clone.Clone(inputUrl, creDir || cfg.Clone.AlwaysCreDir, cfg)
}

func WebHandler(cmd *cobra.Command, args []string) {
	r := getGitRepo(file.GetPwd())

	remoteUrl := getGitRemoteUrl(r)
	branch := getGitBranch(r)

	s, err := config.GetScmSystem(config.GetGitrConfig(), url.GetHost(remoteUrl))
	if err != nil {
		log.Fatal(err)
	}

	repoPath := url.GetRepoPath(remoteUrl, s.Hostname, s.Provider)
	repoName := url.GetRepoName(repoPath)
	webUrl := GetWebUrl(s.Provider, s.Scheme, s.Hostname, repoPath)

	if viper.GetBool(string(Dry)) {
		PrintGitrWebInfo(s.Provider, s.Hostname, remoteUrl, webUrl, repoPath, repoName, branch)
		return
	}

	switch CmdName(cmd.Name()) {
	case Branches:
		url.OpenInBrowser(GetBranchesUrl(s.Provider, webUrl))
	case Prs:
		url.OpenInBrowser(GetPrsUrl(s.Provider, webUrl))
	case Commits:
		url.OpenInBrowser(GetCommitsUrl(s.Provider, webUrl, branch))
	case Issues:
		url.OpenInBrowser(GetIssuesUrl(s.Provider, webUrl))
	case Tags:
		url.OpenInBrowser(GetTagsUrl(s.Provider, webUrl))
	case Releases:
		url.OpenInBrowser(GetReleasesUrl(s.Provider, webUrl))
	case Pipelines:
		url.OpenInBrowser(GetPipelinesUrl(s.Provider, webUrl))
	case Web:
		url.OpenInBrowser(webUrl)
	case Rem:
		url.OpenInBrowser(GetRemUrl(s.Provider, webUrl, branch))
	default:
		log.Fatal("unknown web command")
	}
}
