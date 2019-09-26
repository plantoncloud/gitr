package cmd

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	url "gitr/src/pkg"
	util "gitr/src/pkg"
	"os"
)

var remCmd = &cobra.Command{
	Use:   "rem",
	Short: "Opens the repo on the SCM Web interface",
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
			if gitrRepo.GetWebUrl() != "" {
				open.Run(gitrRepo.GetWebUrl())
			} else {
				fmt.Println("No remote Web URL found for git remote url " + remoteUrl)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(remCmd)
}
