package cmd

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	url "leftbin.com/tools/gitr/pkg"
	util "leftbin.com/tools/gitr/pkg"
	"os"
)

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "open tags on scm web interface",
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
			if gitrRepo.GetTagsUrl() != "" {
				open.Run(gitrRepo.GetTagsUrl())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(tagsCmd)
}
