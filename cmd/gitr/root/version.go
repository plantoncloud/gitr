package root

import (
	"fmt"
	"github.com/plantoncloud/gitr/pkg/version"
	"github.com/spf13/cobra"
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
