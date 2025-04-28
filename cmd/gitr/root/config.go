package root

import (
	"github.com/plantoncloud/gitr/cmd/gitr/root/config"
	"github.com/spf13/cobra"
)

var Config = &cobra.Command{
	Use:   "config",
	Short: "see gitr config",
}

func init() {
	Config.AddCommand(config.Init, config.Show, config.Edit)
}
