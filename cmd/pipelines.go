package cmd

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	url "leftbin.com/tools/gitr/pkg"
	util "leftbin.com/tools/gitr/pkg"
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
			gitrRepo := url.ParseGitRemoteUrl(remoteUrl)
			if viper.GetBool("debug") {
				println(gitrRepo.ToString())
			}
			if gitrRepo.GetPipelinesUrl() != "" {
				open.Run(gitrRepo.GetPipelinesUrl())
			} else {
				println(fmt.Sprintf("SCM Provider %s does not support Pipelines", gitrRepo.ScmProvider))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(pipelinesCmd)
}
