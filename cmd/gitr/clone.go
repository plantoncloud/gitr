package gitr

import (
	"github.com/spf13/cobra"
	handlers "github.com/swarupdonepudi/gitr/v2/internal"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "clones repo to mimic folder structure to the scm repo hierarchy",
	Run:   handlers.Clone,
}

