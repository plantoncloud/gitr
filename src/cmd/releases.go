package cmd

import (
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	url "gitr/src/pkg"
	util "gitr/src/pkg"
	"os"
)

var releasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "Open Releases on SCM Web Interface",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()
		repo := util.GetGitRepo(pwd)
		if repo != nil {
			remoteUrl := util.GetGitRemoteUrl(repo)
			repoUrl := url.ParseGitRemoteUrl(remoteUrl)
			if(repoUrl.ScmProvider == url.GitHub) {
				open.Run(repoUrl.GetReleasesUrl())
			} else {
				println(fmt.Sprintf("SCM Provider %s does not support Releases", repoUrl.ScmProvider))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(releasesCmd)
}
