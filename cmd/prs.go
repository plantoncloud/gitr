package cmd

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	url "github.com/swarupdonepudi/gitr/pkg"
	util "github.com/swarupdonepudi/gitr/pkg"
	"os"
)

var prsCmd = &cobra.Command{
	Use:   "prs",
	Short: "open pull requests on scm web interface",
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
			if gitrRepo.GetPrsUrl() != "" {
				open.Run(gitrRepo.GetPrsUrl())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(prsCmd)
}
