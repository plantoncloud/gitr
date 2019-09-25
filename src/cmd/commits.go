package cmd

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	url "gitr/src/pkg"
	util "gitr/src/pkg"
	"os"
)

var commitsCmd = &cobra.Command{
	Use:   "commits",
	Short: "Open commits on SCM Web Interface",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()
		repo := util.GetGitRepo(pwd)
		if repo != nil {
			remoteUrl := util.GetGitRemoteUrl(repo)
			repoUrl := url.Parse(remoteUrl)
			if repoUrl.GetCommitsUrl() != "" {
				open.Run(repoUrl.GetCommitsUrl())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(commitsCmd)
}
