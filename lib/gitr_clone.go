package lib

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

type GitrClone struct {
	Url    string
	CreDir bool
	Gc     *GitrConfig
}

func ParseCloneReq(args []string, creDir bool, gc *GitrConfig) *GitrClone {
	return &GitrClone{
		Url:    args[0],
		CreDir: creDir,
		Gc:     gc,
	}
}

func (c *GitrClone) PrintInfo() {
	gru := &GitrUtil{}
	var provider ScmProvider
	scmSystem, err := c.Gc.GetScmSystem(gru.GetHost(c.Url))
	if err == nil {
		provider = scmSystem.Provider
	}
	println("")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRow(table.Row{"remote", c.Url})
	t.AppendSeparator()
	t.AppendRow(table.Row{"provider", provider})
	t.AppendSeparator()
	t.AppendRow(table.Row{"host", gru.GetHost(c.Url)})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repo-name", gru.GetRepoName(gru.GetRepoPath(c.Url))})
	t.AppendSeparator()
	t.AppendRow(table.Row{"create-dir", c.Gc.Clone.AlwaysCreDir || c.CreDir})
	t.AppendSeparator()
	t.AppendRow(table.Row{"clone-path", c.GetClonePath()})
	t.AppendSeparator()
	t.Render()
	println("")
}

func (c *GitrClone) Clone() {
	gru := &GitrUtil{}
	if gru.IsGitUrl(c.Url) {
		if gru.IsGitSshUrl(c.Url) {
			err = c.sshClone()
		} else {
			err = c.httpClone()
		}
		if err != nil {
			fmt.Printf("error cloning the repo. %v\n", err)
		}
	} else {
		print("ssh clone using browser urls not support")
	}
}

func (c *GitrClone) GetClonePath() string {
	gru := &GitrUtil{}
	clonePath := ""
	if c.Gc.Get().Clone.AlwaysCreDir {
		if c.Gc.Get().Clone.IncludeHostForCreDir {
			clonePath = fmt.Sprintf("%s/%s", gru.GetHost(c.Url), gru.GetRepoPath(c.Url))
		} else {
			clonePath = gru.GetRepoPath(c.Url)
		}
	} else if c.CreDir {
		clonePath = gru.GetRepoPath(c.Url)
	} else {
		clonePath = gru.GetRepoName(gru.GetRepoPath(c.Url))
	}
	if c.Gc.Get().Clone.ScmHome != "" {
		clonePath = fmt.Sprintf("%s/%s", c.Gc.Get().Clone.ScmHome, clonePath)
	}
	return clonePath
}

func (c *GitrClone) httpClone() error {
	clonePath := c.GetClonePath()
	os.MkdirAll(clonePath, os.ModePerm)
	_, err := git.PlainClone(clonePath, false, &git.CloneOptions{
		URL:      c.Url,
		Progress: os.Stdout,
	})
	return err
}

func (c *GitrClone) sshClone() error {
	gru := &GitrUtil{}
	auth, sshErr := gru.SetUpSshAuth(gru.GetHost(c.Url))

	if sshErr != nil {
		return sshErr
	}
	clonePath := c.GetClonePath()
	os.MkdirAll(clonePath, os.ModePerm)
	_, err := git.PlainClone(clonePath, false, &git.CloneOptions{
		URL:      c.Url,
		Progress: os.Stdout,
		Auth:     auth,
	})
	return err
}
