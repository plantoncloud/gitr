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
	Short: "open branches on scm web interface",
	Run:   h.remoteRepoHandler(h.branches),
}

var clone = &cobra.Command{
	Use:   "clone",
	Short: "clones repo using clone and browser urls",
	Run:   h.clone,
}

var commits = &cobra.Command{
	Use:   "commits",
	Short: "open commits on scm web interface",
	Run:   h.remoteRepoHandler(h.commits),
}

var issues = &cobra.Command{
	Use:   "issues",
	Short: "open issues on scm web interface",
	Run:   h.remoteRepoHandler(h.issues),
}

var pipelines = &cobra.Command{
	Use:     "pipelines",
	Short:   "open pipelines on scm web interface",
	Aliases: []string{"pipe"},
	Run:     h.remoteRepoHandler(h.pipelines),
}

var prs = &cobra.Command{
	Use:   "prs",
	Short: "open pull requests on scm web interface",
	Run:   h.remoteRepoHandler(h.prs),
}

var releases = &cobra.Command{
	Use:   "releases",
	Short: "open releases on scm web interface",
	Run:   h.remoteRepoHandler(h.releases),
}

var rem = &cobra.Command{
	Use:   "rem",
	Short: "opens the local branch of the repo on the scm web interface",
	Run:   h.remoteRepoHandler(h.rem),
}

var remHome = &cobra.Command{
	Use:     "rem-home",
	Aliases: []string{"remh"},
	Short:   "opens the home page of the repo on the scm web interface",
	Run:     h.remoteRepoHandler(h.remHome),
}

var tags = &cobra.Command{
	Use:   "tags",
	Short: "open tags on scm web interface",
	Run:   h.remoteRepoHandler(h.tags),
}

func init() {
	rootCmd.AddCommand(branches, clone, commits, issues, pipelines, prs, releases, rem, remHome, tags)
}
