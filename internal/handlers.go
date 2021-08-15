package internal

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/swarupdonepudi/gitr/v2/pkg/clone"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	"github.com/swarupdonepudi/gitr/v2/pkg/file"
	"github.com/swarupdonepudi/gitr/v2/pkg/url"
	"gopkg.in/yaml.v2"
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
	Config    CmdName  = "config"
)

func ConfigHandler(cmd *cobra.Command, args []string) {
	cfg := config.GetGitrConfig()
	d, err := yaml.Marshal(&cfg)
	fmt.Printf("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%s\n", string(d))
}

func CloneHandler(cmd *cobra.Command, args []string) {
	inputUrl := args[0]
	creDir := viper.GetBool(string(CreDir))
	cfg := config.GetGitrConfig()
	s, err := config.GetScmHost(cfg, url.GetHostname(inputUrl))
	if err != nil {
		log.Fatal(err)
	}
	scmHome := getScmHome(s.Clone.HomeDir, cfg.Scm.HomeDir)
	if viper.GetBool(string(Dry)) {
		printGitrCloneInfo(inputUrl, creDir || s.Clone.AlwaysCreDir, cfg)
		return
	}
	clone.Clone(inputUrl, scmHome, creDir || s.Clone.AlwaysCreDir, cfg.CopyRepoPathCdCmdToClipboard, s)
}

func WebHandler(cmd *cobra.Command, args []string) {
	r := getGitRepo(file.GetPwd())

	remoteUrl := getGitRemoteUrl(r)
	branch := getGitBranch(r)

	s, err := config.GetScmHost(config.GetGitrConfig(), url.GetHostname(remoteUrl))
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

func PathHandler(cmd *cobra.Command, args []string) {
	inputUrl := args[0]
	creDir := viper.GetBool(string(CreDir))
	cfg := config.GetGitrConfig()
	s, err := config.GetScmHost(cfg, url.GetHostname(inputUrl))
	if err != nil {
		log.Fatalf("failed to get scm host. err: %v", err)
	}
	repoPath := url.GetRepoPath(inputUrl, s.Hostname, s.Provider)
	repoName := url.GetRepoName(repoPath)
	scmHome := getScmHome(s.Clone.HomeDir, cfg.Scm.HomeDir)
	repoLocation := clone.GetClonePath(s.Hostname, repoPath, repoName, scmHome, creDir || s.Clone.AlwaysCreDir, s.Clone.IncludeHostForCreDir)
	fmt.Println(repoLocation)
	if err := clipboard.WriteAll(fmt.Sprintf("cd %s", repoLocation)); err != nil {
		log.Fatalf("err copying repo path to clipboard. %v\n", err)
	}
}
