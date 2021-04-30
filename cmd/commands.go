package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func getPwd() string {
	pwd, _ := os.Getwd()
	return pwd
}

var wh = &webCmdHandler{dir: getPwd()}
var ch = &cloneCmdHandler{}

var branches = &cobra.Command{
	Use:   "branches",
	Short: "open branches of the repo in the browser",
	Run:   wh.gitrWebHandler(wh.branches),
}

var clone = &cobra.Command{
	Use:   "clone",
	Short: "clones repo to mimic folder structure to the scm repo hierarchy",
	Run:   ch.Clone,
}

var commits = &cobra.Command{
	Use:   "commits",
	Short: "open commits of the local branch of repo in the browser",
	Run:   wh.gitrWebHandler(wh.commits),
}

var issues = &cobra.Command{
	Use:   "issues",
	Short: "open issues of the repo in the browser",
	Run:   wh.gitrWebHandler(wh.issues),
}

var pipelines = &cobra.Command{
	Use:     "pipelines",
	Short:   "open pipelines/actions of the repo in the browser",
	Aliases: []string{"pipe"},
	Run:     wh.gitrWebHandler(wh.pipelines),
}

var prs = &cobra.Command{
	Use:   "prs",
	Short: "open prs/mrs of the repo in the browser",
	Run:   wh.gitrWebHandler(wh.prs),
}

var releases = &cobra.Command{
	Use:   "releases",
	Short: "open releases of the repo in the browser",
	Run:   wh.gitrWebHandler(wh.releases),
}

var rem = &cobra.Command{
	Use:   "rem",
	Short: "open local checkout branch of the repo in the browser",
	Run:   wh.gitrWebHandler(wh.rem),
}

var tags = &cobra.Command{
	Use:   "tags",
	Short: "open tags of the repo in the browser",
	Run:   wh.gitrWebHandler(wh.tags),
}

var web = &cobra.Command{
	Use:   "web",
	Short: "open home page of the repo in the browser",
	Run:   wh.gitrWebHandler(wh.web),
}

func init() {
	rootCmd.AddCommand(branches, clone, commits, issues, pipelines, prs, releases, rem, tags, web)
}
