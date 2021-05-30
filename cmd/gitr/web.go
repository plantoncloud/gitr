package gitr

import (
	"github.com/spf13/cobra"
	handlers "github.com/swarupdonepudi/gitr/v2/internal"
)

var branchesCmd = &cobra.Command{
	Use:   string(handlers.Branches),
	Short: "open branches of the repo in the browser",
	Run:   handlers.WebHandler,
}

var webCmd = &cobra.Command{
	Use:   string(handlers.Web),
	Short: "open home page of the repo in the browser",
	Run:   handlers.WebHandler,
}

var tagsCmd = &cobra.Command{
	Use:   string(handlers.Tags),
	Short: "open tags of the repo in the browser",
	Run:   handlers.WebHandler,
}

var remCmd = &cobra.Command{
	Use:   string(handlers.Rem),
	Short: "open local checkout branch of the repo in the browser",
	Run:   handlers.WebHandler,
}

var releasesCmd = &cobra.Command{
	Use:   string(handlers.Releases),
	Short: "open releases of the repo in the browser",
	Run:   handlers.WebHandler,
}

var prsCmd = &cobra.Command{
	Use:   string(handlers.Prs),
	Short: "open prs/mrs of the repo in the browser",
	Run:   handlers.WebHandler,
}

var pipelinesCmd = &cobra.Command{
	Use:     string(handlers.Pipelines),
	Short:   "open pipelines/actions of the repo in the browser",
	Aliases: []string{"pipe"},
	Run:     handlers.WebHandler,
}

var issuesCmd = &cobra.Command{
	Use:   string(handlers.Issues),
	Short: "open issues of the repo in the browser",
	Run:   handlers.WebHandler,
}

var commitsCmd = &cobra.Command{
	Use:   string(handlers.Commits),
	Short: "open commits of the local branch of repo in the browser",
	Run:   handlers.WebHandler,
}
