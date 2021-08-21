package gitr

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/gitr/v2/internal/cli"
	"github.com/swarupdonepudi/gitr/v2/pkg/clone"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "clones repo to mimic folder structure to the scm repo hierarchy",
	Run:   cloneHandler,
}

func init() {
	cloneCmd.PersistentFlags().BoolP(string(cli.CreDir), "", false, "cre folders to mimic repo path on scm")
}

func cloneHandler(cmd *cobra.Command, args []string) {
	if len(args) <= 0 {
		log.Fatalf("clone url required as argument")
	}
	inputUrl := args[0]
	dry, err := cmd.InheritedFlags().GetBool(string(cli.Dry))
	cli.HandleFlagErr(err, cli.Dry)
	creDir, err := cmd.PersistentFlags().GetBool(string(cli.CreDir))
	cli.HandleFlagErr(err, cli.CreDir)
	cfg, err := config.NewGitrConfig()
	if err != nil {
		log.Fatalf("failed to get gitr config. err: %v", err)
	}
	if err := clone.Clone(cfg, inputUrl, creDir, dry); err != nil {
		log.Fatalf("failed to clone repo. err: %v", err)
	}
}
