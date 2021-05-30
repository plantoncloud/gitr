package gitr

import (
	"github.com/spf13/cobra"
	h "github.com/swarupdonepudi/gitr/v2/internal"
)

var branchesCmd = &cobra.Command{
	Use:   string(h.Branches),
	Short: "open branches of the repo in the browser",
	Run:   h.WebHandler,
}

var webCmd = &cobra.Command{
	Use:   string(h.Web),
	Short: "open home page of the repo in the browser",
	Run:   h.WebHandler,
}

var tagsCmd = &cobra.Command{
	Use:   string(h.Tags),
	Short: "open tags of the repo in the browser",
	Run:   h.WebHandler,
}

var remCmd = &cobra.Command{
	Use:   string(h.Rem),
	Short: "open local checkout branch of the repo in the browser",
	Run:   h.WebHandler,
}

var releasesCmd = &cobra.Command{
	Use:   string(h.Releases),
	Short: "open releases of the repo in the browser",
	Run:   h.WebHandler,
}

var prsCmd = &cobra.Command{
	Use:   string(h.Prs),
	Short: "open prs/mrs of the repo in the browser",
	Run:   h.WebHandler,
}

var pipelinesCmd = &cobra.Command{
	Use:     string(h.Pipelines),
	Short:   "open pipelines/actions of the repo in the browser",
	Aliases: []string{"pipe"},
	Run:     h.WebHandler,
}

var issuesCmd = &cobra.Command{
	Use:   string(h.Issues),
	Short: "open issues of the repo in the browser",
	Run:   h.WebHandler,
}

var commitsCmd = &cobra.Command{
	Use:   string(h.Commits),
	Short: "open commits of the local branch of repo in the browser",
	Run:   h.WebHandler,
}
