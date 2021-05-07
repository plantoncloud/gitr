package cmd

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	gitr "github.com/swarupdonepudi/gitr/lib"
)

type webCmdHandler struct {
	r         *gitr.GitrWeb
	dir       string
	urlToOpen string
}

type cloneCmdHandler struct{}

func (ch *cloneCmdHandler) Clone(cmd *cobra.Command, args []string) {
	gc := &gitr.GitrConfig{}
	c := gitr.ParseCloneReq(args, viper.GetBool("create-dir"), gc.Get())
	if viper.GetBool("dry") {
		c.PrintInfo()
	} else {
		c.Clone()
	}
}

func openBrowser(url string) {
	if url != "" && !viper.GetBool("dry") {
		_ = open.Run(url)
	}
}

func (h *webCmdHandler) gitrWebHandler(handler func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		h.r = gitr.ScanRepo(h.dir)
		if viper.GetBool("dry") {
			h.r.PrintInfo()
		}
		handler(cmd, args)
		if h.urlToOpen != "" {
			openBrowser(h.urlToOpen)
		}
	}
}

func (h *webCmdHandler) branches(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	h.urlToOpen = r.GetBranchesUrl()
}

func (h *webCmdHandler) commits(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	h.urlToOpen = r.GetCommitsUrl()
}

func (h *webCmdHandler) issues(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	h.urlToOpen = r.GetIssuesUrl()
}

func (h *webCmdHandler) pipelines(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	h.urlToOpen = r.GetPipelinesUrl()
}

func (h *webCmdHandler) prs(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	h.urlToOpen = r.GetPrsUrl()
}

func (h *webCmdHandler) releases(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	h.urlToOpen = r.GetReleasesUrl()
}

func (h *webCmdHandler) rem(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	h.urlToOpen = r.GetRemUrl()
}

func (h *webCmdHandler) web(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	h.urlToOpen = r.GetWebUrl()
}

func (h *webCmdHandler) tags(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(h.dir)
	h.urlToOpen = r.GetTagsUrl()
}
