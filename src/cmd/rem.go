package cmd

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
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
			repoUrl := url.Parse(remoteUrl)
			if repoUrl.GetWebUrl() != "" {
				open.Run(repoUrl.GetWebUrl())
			} else {
				fmt.Println("No remote Web URL found for git remote url " + remoteUrl)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(remCmd)
}
