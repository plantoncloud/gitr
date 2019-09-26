package cmd

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	url "gitr/src/pkg"
	util "gitr/src/pkg"
	"os"
)

var pipelinesCmd = &cobra.Command{
	Use:   "pipelines",
	Short: "Open Pipelines on SCM Web Interface",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()
		repo := util.GetGitRepo(pwd)
		if repo != nil {
			remoteUrl := util.GetGitRemoteUrl(repo)
			repoUrl := url.ParseGitRemoteUrl(remoteUrl)
			if repoUrl.GetPipelinesUrl() != "" {
				open.Run(repoUrl.GetPipelinesUrl())
			} else {
				println(fmt.Sprintf("SCM Provider %s does not support Pipelines", repoUrl.ScmProvider))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pipelinesCmd)
}
