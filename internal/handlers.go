package internal

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	"github.com/swarupdonepudi/gitr/v2/pkg/file"
	"github.com/swarupdonepudi/gitr/v2/pkg/url"
	"log"
)

type GitrFlag string

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

type CmdName string

func CloneHandler(cmd *cobra.Command, args []string) {
	//gc := &gitr.GitrConfig{}
	//c := gitr.ParseCloneReq(args, viper.GetBool("create-dir"), gc.Get())
	//if cmd.Flag("dry").Value() {
	//	c.PrintInfo()
	//} else {
	//	c.Clone()
	//}
}

func WebHandler(cmd *cobra.Command, args []string) {
	var urlToOpen string
	r := getGitRepo(file.GetPwd())
	remoteUrl := getGitRemoteUrl(r)
	branch := getGitBranch(r)
	scmSystem, err := config.GetScmSystem(config.GetGitrConfig(), url.GetHost(remoteUrl))
	if err != nil {
		log.Fatal(err)
	}
	if viper.GetBool("dry") {
		PrintGitrWebInfo(scmSystem, remoteUrl, branch)
		return
	}
	webUrl := GetWebUrl(scmSystem.Provider, remoteUrl)
	url.OpenInBrowser(urlToOpen)
	switch CmdName(cmd.Name()) {
	case Branches:
		url.OpenInBrowser(GetBranchesUrl(scmSystem.Provider, webUrl))
	case Prs:
		url.OpenInBrowser(GetPrsUrl(scmSystem.Provider, webUrl))
	case Commits:
		url.OpenInBrowser(GetCommitsUrl(scmSystem.Provider, webUrl, branch))
	case Issues:
		url.OpenInBrowser(GetIssuesUrl(scmSystem.Provider, webUrl))
	case Tags:
		url.OpenInBrowser(GetTagsUrl(scmSystem.Provider, webUrl))
	case Releases:
		url.OpenInBrowser(GetReleasesUrl(scmSystem.Provider, webUrl))
	case Pipelines:
		url.OpenInBrowser(GetPipelinesUrl(scmSystem.Provider, webUrl))
	case Web:
		url.OpenInBrowser(webUrl)
	case Rem:
		url.OpenInBrowser(GetRemUrl(scmSystem.Provider, webUrl, branch))
	default:
		log.Fatal("unknown web command")
	}
}
