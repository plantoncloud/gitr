package gitr

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/gitr/v2/cmd/gitr/root"
	"github.com/swarupdonepudi/gitr/v2/internal/cli"
	"github.com/swarupdonepudi/gitr/v2/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "gitr",
	Short: "git rapid - the missing link b/w git cli & scm providers",
	Long:  "save time(a ton) by opening git repos on web browser right from the command line",
}

func init() {
	rootCmd.PersistentFlags().BoolP(string(cli.Dry), "", false, "dry run")
	rootCmd.AddCommand(
		root.Version,
		root.Config,
		root.Clone,
		root.Path,
		root.BranchesCmd,
		root.CommitsCmd,
		root.IssuesCmd,
		root.PipelinesCmd,
		root.PrsCmd,
		root.ReleasesCmd,
		root.RemCmd,
		root.TagsCmd,
		root.WebCmd,
	)
	if err := config.EnsureInitialConfig(); err != nil {
		log.Fatalf("failed to initialize config. err %v", err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to run command. err %v", err)
	}
}
