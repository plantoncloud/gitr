package cmd

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	url "leftbin.com/tools/gitr/pkg"
	util "leftbin.com/tools/gitr/pkg"
	"os"
)

var commitsCmd = &cobra.Command{
	Use:   "commits",
	Short: "open commits on scm web interface",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()
		repo := util.GetGitRepo(pwd)
		if repo != nil {
			remoteUrl := util.GetGitRemoteUrl(repo)
			gitrRepo := url.ParseGitRemoteUrl(remoteUrl)
			if viper.GetBool("debug") {
				println(gitrRepo.ToString())
			}
			if gitrRepo.GetCommitsUrl() != "" {
				open.Run(gitrRepo.GetCommitsUrl())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(commitsCmd)
}
