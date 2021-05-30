package gitr

import (
	"github.com/spf13/cobra"
	h "github.com/swarupdonepudi/gitr/v2/internal"
)

var cloneCmd = &cobra.Command{
	Use:   string(h.Clone),
	Short: "clones repo to mimic folder structure to the scm repo hierarchy",
	Run:   h.CloneHandler,
}
