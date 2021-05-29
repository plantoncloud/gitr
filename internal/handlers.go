package internal

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	gitr "github.com/swarupdonepudi/gitr/v2/pkg"
	"log"
)

type GitrFlag string

const (
	Dry    GitrFlag = "dry"
	CreDir GitrFlag = "create-dir"
)

var urlToOpen string

func Clone(cmd *cobra.Command, args []string) {
	//gc := &gitr.GitrConfig{}
	//c := gitr.ParseCloneReq(args, viper.GetBool("create-dir"), gc.Get())
	//if cmd.Flag("dry").Value() {
	//	c.PrintInfo()
	//} else {
	//	c.Clone()
	//}
}

func openBrowser(url string) {
	if url != "" && !viper.GetBool("dry") {
		_ = open.Run(url)
	}
}

func WebHandler(handler func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		r := getGitRepo(GetPwd())
		remoteUrl := getGitRemoteUrl(r)
		branch := getGitBranch(r)
		scmSystem, err := gitr.GetScmSystem(getHost(remoteUrl))
		if err != nil {
			log.Fatal(err)
		}
		if viper.GetBool("dry") {
			PrintGitrWebInfo(scmSystem, remoteUrl, branch)
		}
		handler(cmd, args)
		openBrowser(urlToOpen)
	}
}

func Branches(cmd *cobra.Command, args []string) {
	r := getGitRepo(GetPwd())
	remoteUrl := getGitRemoteUrl(r)
	scmSystem, err := gitr.GetScmSystem(getHost(remoteUrl))
	if err != nil {
		log.Fatal(err)
	}
	webUrl := GetWebUrl(scmSystem.Provider, remoteUrl)
	urlToOpen = GetBranchesUrl(scmSystem.Provider, webUrl)
}

func Commits(cmd *cobra.Command, args []string) {
	r := getGitRepo(GetPwd())
	remoteUrl := getGitRemoteUrl(r)
	scmSystem, err := gitr.GetScmSystem(getHost(remoteUrl))
	if err != nil {
		log.Fatal(err)
	}
	webUrl := GetWebUrl(scmSystem.Provider, remoteUrl)
	urlToOpen = GetCommitsUrl(scmSystem, webUrl, getGitBranch(r))
}

func Issues(cmd *cobra.Command, args []string) {
	r := getGitRepo(GetPwd())
	remoteUrl := getGitRemoteUrl(r)
	scmSystem, err := gitr.GetScmSystem(getHost(remoteUrl))
	if err != nil {
		log.Fatal(err)
	}
	webUrl := GetWebUrl(scmSystem.Provider, remoteUrl)
	urlToOpen = GetIssuesUrl(scmSystem.Provider, webUrl)
}

func Pipelines(cmd *cobra.Command, args []string) {
	r := getGitRepo(GetPwd())
	remoteUrl := getGitRemoteUrl(r)
	scmSystem, err := gitr.GetScmSystem(getHost(remoteUrl))
	if err != nil {
		log.Fatal(err)
	}
	webUrl := GetWebUrl(scmSystem.Provider, remoteUrl)
	urlToOpen = GetPipelinesUrl(scmSystem.Provider, webUrl)
}

func Prs(cmd *cobra.Command, args []string) {
	r := getGitRepo(GetPwd())
	remoteUrl := getGitRemoteUrl(r)
	scmSystem, err := gitr.GetScmSystem(getHost(remoteUrl))
	if err != nil {
		log.Fatal(err)
	}
	webUrl := GetWebUrl(scmSystem.Provider, remoteUrl)
	urlToOpen = GetPrsUrl(scmSystem.Provider, webUrl)
}

func Releases(cmd *cobra.Command, args []string) {
	r := getGitRepo(GetPwd())
	remoteUrl := getGitRemoteUrl(r)
	scmSystem, err := gitr.GetScmSystem(getHost(remoteUrl))
	if err != nil {
		log.Fatal(err)
	}
	webUrl := GetWebUrl(scmSystem.Provider, remoteUrl)
	urlToOpen = GetReleasesUrl(scmSystem.Provider, webUrl)
}

func Rem(cmd *cobra.Command, args []string) {
	r := getGitRepo(GetPwd())
	remoteUrl := getGitRemoteUrl(r)
	scmSystem, err := gitr.GetScmSystem(getHost(remoteUrl))
	if err != nil {
		log.Fatal(err)
	}
	webUrl := GetWebUrl(scmSystem.Provider, remoteUrl)
	branch := getGitBranch(r)
	urlToOpen = GetRemUrl(scmSystem.Provider, webUrl, branch)
}

func Web(cmd *cobra.Command, args []string) {
	r := getGitRepo(GetPwd())
	remoteUrl := getGitRemoteUrl(r)
	scmSystem, err := gitr.GetScmSystem(getHost(remoteUrl))
	if err != nil {
		log.Fatal(err)
	}
	urlToOpen = GetWebUrl(scmSystem.Provider, remoteUrl)
}

func Tags(cmd *cobra.Command, args []string) {
	r := getGitRepo(GetPwd())
	remoteUrl := getGitRemoteUrl(r)
	scmSystem, err := gitr.GetScmSystem(getHost(remoteUrl))
	if err != nil {
		log.Fatal(err)
	}
	webUrl := GetWebUrl(scmSystem.Provider, remoteUrl)
	urlToOpen = GetTagsUrl(scmSystem.Provider, webUrl)
}
