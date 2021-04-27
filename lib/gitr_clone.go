package lib

import (
	"github.com/go-git/go-git/v5"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

type GitrClone struct {
	Url      string
	Scheme   GitRemoteScheme // http, https or ssh
	Provider ScmProvider
	CreDir   bool
}

func ParseCloneReq(args []string, creDir bool) *GitrClone {
	return &GitrClone{
		Url:      args[0],
		Scheme:   "some-scheme",
		Provider: "some-provider",
		CreDir:   creDir,
	}
}

func (c *GitrClone) PrintInfo() {
	println("")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRow(table.Row{"url", c.Url})
	t.AppendSeparator()
	t.AppendRow(table.Row{"create-dir", c.CreDir})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Provider", c.Provider})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Scheme", c.Scheme})
	t.AppendSeparator()
	t.Render()
	println("")
}

func (c *GitrClone) Clone() {
	c.httpClone()
}

func (c *GitrClone) httpClone() error {
	repoName := GetRepoName(GetRepoPath(c.Url))
	os.Mkdir(repoName, os.ModePerm)
	_, err := git.PlainClone(repoName, false, &git.CloneOptions{
		URL:      c.Url,
		Progress: os.Stdout,
	})
	return err
}
