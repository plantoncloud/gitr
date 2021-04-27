package cmd

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	gitr "github.com/swarupdonepudi/gitr/lib"
)

type cmdHandler struct{}

func openBrowser(url string) {
	if url != "" {
		open.Run(url)
	}
}

func (h *cmdHandler) branches(cmd *cobra.Command, args []string) {
	o := &gitr.GitRepo{}
	o.Scan()
	openBrowser(o.GetBranchesUrl())
}

func (h *cmdHandler) commits(cmd *cobra.Command, args []string) {
	o := &gitr.GitRepo{}
	o.Scan()
	openBrowser(o.GetCommitsUrl())
}

func (h *cmdHandler) issues(cmd *cobra.Command, args []string) {
	o := &gitr.GitRepo{}
	o.Scan()
	openBrowser(o.GetIssuesUrl())
}

func (h *cmdHandler) pipelines(cmd *cobra.Command, args []string) {
	o := &gitr.GitRepo{}
	o.Scan()
	openBrowser(o.GetPipelinesUrl())
}

func (h *cmdHandler) prs(cmd *cobra.Command, args []string) {
	o := &gitr.GitRepo{}
	o.Scan()
	openBrowser(o.GetPrsUrl())
}

func (h *cmdHandler) releases(cmd *cobra.Command, args []string) {
	o := &gitr.GitRepo{}
	o.Scan()
	openBrowser(o.GetReleasesUrl())
}

func (h *cmdHandler) rem(cmd *cobra.Command, args []string) {
	o := &gitr.GitRepo{}
	o.Scan()
	openBrowser(o.GetRemUrl())
}

func (h *cmdHandler) tags(cmd *cobra.Command, args []string) {
	o := &gitr.GitRepo{}
	o.Scan()
	openBrowser(o.GetTagsUrl())
}
