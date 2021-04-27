package lib

import (
	"github.com/go-git/go-git/v5"
	"github.com/jedib0t/go-pretty/v6/table"
	"log"
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
	gru := &GitrUtil{}
	if gru.IsGitUrl(c.Url) {
		if gru.IsGitSshUrl(c.Url) {
			err = c.sshClone()
		} else {
			err = c.httpClone()
		}
		if err != nil {
			log.Fatal("error cloning the repo", err)
		}
	} else {
		print("ssh clone using browser urls not implemented yet")
	}
}

func (c *GitrClone) setupClonePath() string {
	gru := &GitrUtil{}
	if c.CreDir {
		os.MkdirAll(gru.GetRepoPath(c.Url), os.ModePerm)
		return gru.GetRepoPath(c.Url)
	} else {
		repoName := gru.GetRepoName(gru.GetRepoPath(c.Url))
		os.Mkdir(repoName, os.ModePerm)
		return repoName
	}
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
