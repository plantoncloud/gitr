package gitr

import (
	"github.com/spf13/cobra"
	h "github.com/swarupdonepudi/gitr/v2/internal"
)

var configCmd = &cobra.Command{
	Use:   string(h.Config),
	Short: "see gitr config",
	Run:   h.ConfigHandler,
}
