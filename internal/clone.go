package internal

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/swarupdonepudi/gitr/v2/pkg/clone"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	"github.com/swarupdonepudi/gitr/v2/pkg/url"
	"log"
	"os"
)

func PrintGitrCloneInfo(inputUrl string, creDir bool, cfg *config.GitrConfig) {
	clonePath := clone.GetClonePath(inputUrl, cfg.Clone.ScmHome, creDir || cfg.Clone.AlwaysCreDir, cfg.Clone.IncludeHostForCreDir)
	scmSystem, err := config.GetScmSystem(cfg, url.GetHost(inputUrl))
	if err != nil {
		log.Fatal(err)
	}
	println("")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRow(table.Row{"remote", inputUrl})
	t.AppendSeparator()
	t.AppendRow(table.Row{"provider", scmSystem.Provider})
	t.AppendSeparator()
	t.AppendRow(table.Row{"host", scmSystem.Hostname})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repo-name", url.GetRepoName(url.GetRepoPath(inputUrl))})
	t.AppendSeparator()
	t.AppendRow(table.Row{"create-dir", cfg.Clone.AlwaysCreDir || creDir})
	t.AppendSeparator()
	t.AppendRow(table.Row{"clone-path", clonePath})
	t.AppendSeparator()
	t.Render()
	println("")
}
