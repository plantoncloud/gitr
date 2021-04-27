package cmd

import (
	"github.com/spf13/cobra"
)

var h = &cmdHandler{}

var branches = &cobra.Command{
	Use:   "branches",
	Short: "open branches on scm web interface",
	Run:   h.remoteRepoHandler(h.branches),
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
	Short: "opens the repo on the scm web interface",
	Run:   h.remoteRepoHandler(h.rem),
}

var tags = &cobra.Command{
	Use:   "tags",
	Short: "open tags on scm web interface",
	Run:   h.remoteRepoHandler(h.tags),
}

func init() {
	rootCmd.AddCommand(branches, commits, issues, pipelines, prs, releases, rem, tags)
}
