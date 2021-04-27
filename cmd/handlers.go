package cmd

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	gitr "github.com/swarupdonepudi/gitr/lib"
)

type cmdHandler struct{}

func (h *cmdHandler) branches(cmd *cobra.Command, args []string) {
	o := &gitr.GitOrigin{}
	o.ScanOrigin()
	openBrowser(o.GetBranchesUrl())
}

func (h *cmdHandler) commits(cmd *cobra.Command, args []string) {
	o := &gitr.GitOrigin{}
	o.ScanOrigin()
	openBrowser(o.GetCommitsUrl())
}

func (h *cmdHandler) issues(cmd *cobra.Command, args []string) {
	o := &gitr.GitOrigin{}
	o.ScanOrigin()
	openBrowser(o.GetIssuesUrl())
}

func (h *cmdHandler) pipelines(cmd *cobra.Command, args []string) {
	o := &gitr.GitOrigin{}
	o.ScanOrigin()
	openBrowser(o.GetPipelinesUrl())
}

func (h *cmdHandler) prs(cmd *cobra.Command, args []string) {
	o := &gitr.GitOrigin{}
	o.ScanOrigin()
	openBrowser(o.GetPrsUrl())
}

func (h *cmdHandler) releases(cmd *cobra.Command, args []string) {
	o := &gitr.GitOrigin{}
	o.ScanOrigin()
	openBrowser(o.GetReleasesUrl())
}

func (h *cmdHandler) rem(cmd *cobra.Command, args []string) {
	o := &gitr.GitOrigin{}
	o.ScanOrigin()

	openBrowser(o.GetWebUrl())
}

func openBrowser(url string) {
	if url != "" {
		open.Run(url)
	}
}

func (h *cmdHandler) tags(cmd *cobra.Command, args []string) {
	o := &gitr.GitOrigin{}
	o.ScanOrigin()
	openBrowser(o.GetTagsUrl())
}
