package cmd

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	url "gitr/src/pkg"
	util "gitr/src/pkg"
	"os"
)

var prsCmd = &cobra.Command{
	Use:   "prs",
	Short: "Open Pull Requests on SCM Web Interface",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()
		repo := util.GetGitRepo(pwd)
		if repo != nil {
			remoteUrl := util.GetGitRemoteUrl(repo)
			repoUrl := url.Parse(remoteUrl)
			open.Run(repoUrl.GetPrsUrl())
		}
	},
}

func init() {
	rootCmd.AddCommand(prsCmd)
}
