package cmd

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	url "gitr/src/pkg"
	util "gitr/src/pkg"
	"os"
)

var branchesCmd = &cobra.Command{
	Use:   "branches",
	Short: "Open branches on SCM Web Interface",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()
		repo := util.GetGitRepo(pwd)
		if repo != nil {
			remoteUrl := util.GetGitRemoteUrl(repo)
			repoUrl := url.ParseGitRemoteUrl(remoteUrl)
			if repoUrl.GetBranchesUrl() != "" {
				open.Run(repoUrl.GetBranchesUrl())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(branchesCmd)
}
