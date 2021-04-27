package cmd

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	gitr "github.com/swarupdonepudi/gitr/lib"
	"os"
)

type cmdHandler struct {
	r *gitr.RemoteRepo
}

func getPwd() string {
	pwd, _ := os.Getwd()
	return pwd
}

func openBrowser(url string) {
	if url != "" {
		_ = open.Run(url)
	}
}

func (h *cmdHandler) remoteRepoHandler(handler func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		h.r = gitr.ScanRepo(getPwd())
		if viper.GetBool("debug") {
			h.r.PrintInfo()
		}
		handler(cmd, args)
	}
}

func (h *cmdHandler) branches(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(getPwd())
	openBrowser(r.GetBranchesUrl())
}

func (h *cmdHandler) commits(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(getPwd())
	openBrowser(r.GetCommitsUrl())
}

func (h *cmdHandler) issues(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(getPwd())
	openBrowser(r.GetIssuesUrl())
}

func (h *cmdHandler) pipelines(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(getPwd())
	openBrowser(r.GetPipelinesUrl())
}

func (h *cmdHandler) prs(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(getPwd())
	openBrowser(r.GetPrsUrl())
}

func (h *cmdHandler) releases(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(getPwd())
	openBrowser(r.GetReleasesUrl())
}

func (h *cmdHandler) rem(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(getPwd())
	openBrowser(r.GetRemUrl())
}

func (h *cmdHandler) tags(cmd *cobra.Command, args []string) {
	r := gitr.ScanRepo(getPwd())
	openBrowser(r.GetTagsUrl())
}
