package cmd

import (
	"github.com/spf13/cobra"
	url "gitr/src/pkg"
	util "gitr/src/pkg"
	"log"
	"os"
	"github.com/skratchdot/open-golang/open"
)

var remCmd = &cobra.Command{
	Use:   "rem",
	Short: "Opens the repo on the SCM Web interface",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()
		repo := util.GetGitRepo(pwd)
		if repo != nil {
			remotes, err2 := repo.Remotes()
			if err2 != nil {
				println("2")
				log.Fatal(err2)
			}
			remoteUrl := remotes[0].Config().URLs[0]
			repoUrl := url.Parse(remoteUrl)
			open.Run(repoUrl.GetWebUrl())
		}
	},
}

func init() {
	rootCmd.AddCommand(remCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
