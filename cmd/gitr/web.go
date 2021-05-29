package gitr

import (
	"github.com/spf13/cobra"
	handlers "github.com/swarupdonepudi/gitr/v2/internal"
)

var branchesCmd = &cobra.Command{
	Use:   "branches",
	Short: "open branches of the repo in the browser",
	Run:   handlers.WebHandler(handlers.Branches),
}

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "open home page of the repo in the browser",
	Run:   handlers.WebHandler(handlers.Web),
}

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "open tags of the repo in the browser",
	Run:   handlers.WebHandler(handlers.Tags),
}

var remCmd = &cobra.Command{
	Use:   "rem",
	Short: "open local checkout branch of the repo in the browser",
	Run:   handlers.WebHandler(handlers.Rem),
}

var releasesCmd = &cobra.Command{
	Use:   "releases",
	Short: "open releases of the repo in the browser",
	Run:   handlers.WebHandler(handlers.Releases),
}

var prsCmd = &cobra.Command{
	Use:   "prs",
	Short: "open prs/mrs of the repo in the browser",
	Run:   handlers.WebHandler(handlers.Prs),
}

var pipelinesCmd = &cobra.Command{
	Use:     "pipelines",
	Short:   "open pipelines/actions of the repo in the browser",
	Aliases: []string{"pipe"},
	Run:     handlers.WebHandler(handlers.Pipelines),
}

var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "open issues of the repo in the browser",
	Run:   handlers.WebHandler(handlers.Issues),
}

var commitsCmd = &cobra.Command{
	Use:   "commits",
	Short: "open commits of the local branch of repo in the browser",
	Run:   handlers.WebHandler(handlers.Commits),
}
