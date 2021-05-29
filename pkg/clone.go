package pkg

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/jedib0t/go-pretty/v6/table"
	"log"
	"os"
)

func ParseCloneReq(args []string, creDir bool, gc *GitrConfig) *GitrClone {
	return &GitrClone{
		Url:    args[0],
		CreDir: creDir,
		Gc:     gc,
	}
}

func PrintCloneInfo() {
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

func Clone() {
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
		scmSystem, err := c.Gc.GetScmSystem(gru.GetHost(c.Url))
		if err != nil {
			log.Fatal(err)
		}
		if scmSystem.Provider == BitBucketDatacenter {
			println("clone using browser urls for bitbucket-datacenter is not supported")
			return
		}
		inputUrl := c.Url
		c.Url = fmt.Sprintf("git@%s:%s.git", gru.GetHost(c.Url), gru.GetRepoPath(c.Url))
		err = c.sshClone()
		if err != nil {

		}
	}
}

func GetClonePath(url string) string {
	gru := &GitrUtil{}
	clonePath := ""
	if c.Gc.Clone.AlwaysCreDir {
		if c.Gc.Clone.IncludeHostForCreDir {
			clonePath = fmt.Sprintf("%s/%s", gru.GetHost(url), gru.GetRepoPath(url))
		} else {
			clonePath = gru.GetRepoPath(url)
		}
	} else if c.CreDir {
		clonePath = gru.GetRepoPath(url)
	} else {
		clonePath = gru.GetRepoName(gru.GetRepoPath(url))
	}
	if c.Gc.Clone.ScmHome != "" {
		clonePath = fmt.Sprintf("%s/%s", c.Gc.Clone.ScmHome, clonePath)
	}
	return clonePath
}

func httpClone(url string) error {
	clonePath := c.GetClonePath(url)
	os.MkdirAll(clonePath, os.ModePerm)
	_, err := git.PlainClone(clonePath, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	return err
}

func sshClone(url string) error {
	gru := &GitrUtil{}
	auth, sshErr := gru.SetUpSshAuth(gru.GetHost(url))

	if sshErr != nil {
		return sshErr
	}
	clonePath := c.GetClonePath()
	os.MkdirAll(clonePath, os.ModePerm)
	_, err := git.PlainClone(clonePath, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
		Auth:     auth,
	})
	return err
}
