package cmd

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	gitr "github.com/swarupdonepudi/gitr/lib"
)

type cmdHandler struct {
	r   *gitr.GitrWeb
	dir string
}

func openBrowser(url string) {
	if url != "" {
		_ = open.Run(url)
	}
}

func (h *cmdHandler) gitrWebHandler(handler func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		h.r = gitr.ScanRepo(h.dir)
		if viper.GetBool("debug") {
			h.r.PrintInfo()
		}
		handler(cmd, args)
	}
}

func (h *cmdHandler) branches(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	openBrowser(r.GetBranchesUrl())
}

func (h *cmdHandler) clone(cmd *cobra.Command, args []string) {
	c := gitr.ParseCloneReq(args, viper.GetBool("create-dir"))
	if viper.GetBool("debug") {
		c.PrintInfo()
	}
	c.Clone()
}

func (h *cmdHandler) commits(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	openBrowser(r.GetCommitsUrl())
}

func (h *cmdHandler) issues(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	openBrowser(r.GetIssuesUrl())
}

func (h *cmdHandler) pipelines(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	openBrowser(r.GetPipelinesUrl())
}

func (h *cmdHandler) prs(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	openBrowser(r.GetPrsUrl())
}

func (h *cmdHandler) releases(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	openBrowser(r.GetReleasesUrl())
}

func (h *cmdHandler) rem(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	openBrowser(r.GetRemUrl())
}

func (h *cmdHandler) web(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	openBrowser(r.GetWebUrl())
}

func (h *cmdHandler) tags(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	openBrowser(r.GetTagsUrl())
}
