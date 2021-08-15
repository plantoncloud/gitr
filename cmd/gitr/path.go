package gitr

import (
	"github.com/spf13/cobra"
	h "github.com/swarupdonepudi/gitr/v2/internal"
)

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "prints the path to which the repo is cloned/will be cloned",
	Run:   h.PathHandler,
}
