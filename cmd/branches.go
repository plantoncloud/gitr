package cmd

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	url "github.com/swarupdonepudi/gitr/pkg"
	util "github.com/swarupdonepudi/gitr/pkg"
	"os"
)

var branchesCmd = &cobra.Command{
	Use:   "branches",
	Short: "open branches on scm web interface",
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
			if gitrRepo.GetBranchesUrl() != "" {
				open.Run(gitrRepo.GetBranchesUrl())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(branchesCmd)
}
