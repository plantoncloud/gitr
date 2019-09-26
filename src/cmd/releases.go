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

var releasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "Open Releases on SCM Web Interface",
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
			if gitrRepo.GetReleasesUrl() != "" {
				open.Run(gitrRepo.GetReleasesUrl())
			} else {
				println(fmt.Sprintf("SCM Provider %s does not support Releases", gitrRepo.ScmProvider))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(releasesCmd)
}
