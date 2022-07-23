package root

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/swarupdonepudi/gitr/pkg/version"
)

var Version = &cobra.Command{
	Use:     "version",
	Short:   "check the version of the cli",
	Aliases: []string{"v"},
	Run:     versionHandler,
}

func versionHandler(cmd *cobra.Command, args []string) {
	fmt.Println(fmt.Sprintf("version %s", version.Version))
}
