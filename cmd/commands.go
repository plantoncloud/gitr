package cmd

import (
	"github.com/spf13/cobra"
)

var h = &cmdHandler{}

var branches = &cobra.Command{
	Use:   "branches",
	Short: "open branches on scm web interface",
	Long:  ``,
	Run:   h.branches,
}

var commits = &cobra.Command{
	Use:   "commits",
	Short: "open commits on scm web interface",
	Long:  ``,
	Run:   h.commits,
}

var issues = &cobra.Command{
	Use:   "issues",
	Short: "open issues on scm web interface",
	Long:  ``,
	Run:   h.issues,
}
var pipelines = &cobra.Command{
	Use:     "pipelines",
	Short:   "open pipelines on scm web interface",
	Long:    ``,
	Aliases: []string{"pipe"},
	Run:     h.pipelines,
}

var prs = &cobra.Command{
	Use:   "prs",
	Short: "open pull requests on scm web interface",
	Long:  ``,
	Run:   h.prs,
}

var releases = &cobra.Command{
	Use:   "releases",
	Short: "open releases on scm web interface",
	Long:  ``,
	Run:   h.releases,
}

var rem = &cobra.Command{
	Use:   "rem",
	Short: "opens the repo on the scm web interface",
	Long:  ``,
	Run:   h.rem,
}

var tags = &cobra.Command{
	Use:   "tags",
	Short: "open tags on scm web interface",
	Long:  ``,
	Run:   h.tags,
}

func init() {
	rootCmd.AddCommand(branches)
	rootCmd.AddCommand(commits)
	rootCmd.AddCommand(issues)
	rootCmd.AddCommand(pipelines)
	rootCmd.AddCommand(prs)
	rootCmd.AddCommand(releases)
	rootCmd.AddCommand(rem)
	rootCmd.AddCommand(tags)
}
