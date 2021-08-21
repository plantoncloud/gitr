package gitr

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/swarupdonepudi/gitr/v2/internal"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gitr",
	Short: "git rapid - the missing link b/w git cli & scm providers",
	Long:  `save time(a ton) by opening git repos on web browser right from the command line`,
}

func init() {
	rootCmd.AddCommand(
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
		versionCmd,
	)
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolP(string(internal.Dry), "d", false, "dry run")
	cloneCmd.PersistentFlags().BoolP(string(internal.CreDir), "c", false, "create directories")
	viper.BindPFlag(string(internal.Dry), rootCmd.PersistentFlags().Lookup(string(internal.Dry)))
	viper.BindPFlag(string(internal.CreDir), cloneCmd.PersistentFlags().Lookup(string(internal.CreDir)))
}

func initConfig() {
	internal.EnsureInitialConfig()
	config.LoadViperConfig()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
