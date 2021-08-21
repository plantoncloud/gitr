package gitr

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/gitr/v2/internal/cli"
	"github.com/swarupdonepudi/gitr/v2/internal/config"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gitr",
	Short: "git rapid - the missing link b/w git cli & scm providers",
	Long:  "save time(a ton) by opening git repos on web browser right from the command line",
}

func init() {
	rootCmd.PersistentFlags().BoolP(string(cli.Dry), "", false, "dry run")
	rootCmd.AddCommand(
		versionCmd,
		configCmd,
		cloneCmd,
		pathCmd,
		branchesCmd,
		commitsCmd,
		issuesCmd,
		pipelinesCmd,
		prsCmd,
		releasesCmd,
		remCmd,
		tagsCmd,
		webCmd,
	)
	initConfig()
}

func initConfig() {
	config.EnsureInitialConfig()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
