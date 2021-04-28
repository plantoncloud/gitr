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
}

func ParseCloneReq(args []string, creDir bool) *GitrClone {
	return &GitrClone{
		Url:    args[0],
		CreDir: creDir,
	}
}

func (c *GitrClone) PrintInfo() {
	gc := &GitrConfig{}
	gru := &GitrUtil{}
	scmProvider, err := gc.GetScmProvider(gru.GetHost(c.Url))
	if err != nil {
		scmProvider = "error"
	}
	println("")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRow(table.Row{"url", c.Url})
	t.AppendSeparator()
	t.AppendRow(table.Row{"create-dir", gc.Clone.AlwaysCreDir || c.CreDir})
	t.AppendSeparator()
	t.AppendRow(table.Row{"scm-provider", scmProvider})
	t.AppendSeparator()
	t.AppendRow(table.Row{"scm-host", gru.GetHost(c.Url)})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repo-name", gru.GetRepoName(gru.GetRepoPath(c.Url))})
	t.AppendSeparator()
	t.AppendRow(table.Row{"clone-path", c.setupClonePath()})
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

func (c *GitrClone) setupClonePath() string {
	gc := &GitrConfig{}
	gru := &GitrUtil{}
	clonePath := ""
	if gc.Get().Clone.AlwaysCreDir {
		if gc.Get().Clone.IncludeHostForCreDir {
			clonePath = fmt.Sprintf("%s/%s", gru.GetHost(c.Url), gru.GetRepoPath(c.Url))
		} else {
			clonePath = gru.GetRepoPath(c.Url)
		}
	} else if c.CreDir {
		clonePath = gru.GetRepoPath(c.Url)
	} else {
		clonePath = gru.GetRepoName(gru.GetRepoPath(c.Url))
	}
	if gc.Get().Clone.ScmHome != "" {
		clonePath = fmt.Sprintf("%s/%s", gc.Get().Clone.ScmHome, clonePath)
	}
	os.MkdirAll(clonePath, os.ModePerm)
	return clonePath
}

func (c *GitrClone) httpClone() error {
	clonePath := c.setupClonePath()
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
	clonePath := c.setupClonePath()
	_, err := git.PlainClone(clonePath, false, &git.CloneOptions{
		URL:      c.Url,
		Progress: os.Stdout,
		Auth:     auth,
	})
	return err
}
