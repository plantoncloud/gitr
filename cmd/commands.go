package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func getPwd() string {
	pwd, _ := os.Getwd()
	return pwd
}

var h = &cmdHandler{dir: getPwd()}

var branches = &cobra.Command{
	Use:   "branches",
	Short: "open branches of the repo in the browser",
	Run:   h.gitrWebHandler(h.branches),
}

var clone = &cobra.Command{
	Use:   "clone",
	Short: "clones repo to mimic folder structure to the scm repo hierarchy",
	Run:   h.clone,
}

var commits = &cobra.Command{
	Use:   "commits",
	Short: "open commits of the local branch of repo in the browser",
	Run:   h.gitrWebHandler(h.commits),
}

var issues = &cobra.Command{
	Use:   "issues",
	Short: "open issues of the repo in the browser",
	Run:   h.gitrWebHandler(h.issues),
}

var pipelines = &cobra.Command{
	Use:     "pipelines",
	Short:   "open pipelines/actions of the repo in the browser",
	Aliases: []string{"pipe"},
	Run:     h.gitrWebHandler(h.pipelines),
}

var prs = &cobra.Command{
	Use:   "prs",
	Short: "open prs/mrs of the repo in the browser",
	Run:   h.gitrWebHandler(h.prs),
}

var releases = &cobra.Command{
	Use:   "releases",
	Short: "open releases of the repo in the browser",
	Run:   h.gitrWebHandler(h.releases),
}

var rem = &cobra.Command{
	Use:   "rem",
	Short: "open local checkout branch of the repo in the browser",
	Run:   h.gitrWebHandler(h.rem),
}

var tags = &cobra.Command{
	Use:   "tags",
	Short: "open tags of the repo in the browser",
	Run:   h.gitrWebHandler(h.tags),
}

var web = &cobra.Command{
	Use:   "web",
	Short: "open home page of the repo in the browser",
	Run:   h.gitrWebHandler(h.web),
}

func init() {
	rootCmd.AddCommand(branches, clone, commits, issues, pipelines, prs, releases, rem, tags, web)
}
