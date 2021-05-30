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
	s, err := config.GetScmSystem(cfg, url.GetHost(inputUrl))
	repoPath := url.GetRepoPath(inputUrl, s.Hostname, s.Provider)
	repoName := url.GetRepoName(repoPath)
	clonePath := clone.GetClonePath(s.Hostname, repoPath, repoName, s.Clone.HomeDir, creDir || s.Clone.AlwaysCreDir, s.Clone.IncludeHostForCreDir)
	if err != nil {
		log.Fatal(err)
	}
	println("")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRow(table.Row{"remote", inputUrl})
	t.AppendSeparator()
	t.AppendRow(table.Row{"provider", s.Provider})
	t.AppendSeparator()
	t.AppendRow(table.Row{"host", s.Hostname})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repo-name", repoName})
	t.AppendSeparator()
	t.AppendRow(table.Row{"create-dir", s.Clone.AlwaysCreDir || creDir})
	t.AppendSeparator()
	t.AppendRow(table.Row{"clone-path", clonePath})
	t.AppendSeparator()
	t.Render()
	println("")
}
