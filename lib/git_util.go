package lib

import (
	"github.com/go-git/go-git/v5"
	"log"
	"path/filepath"
	"strings"
)

type GitUtil struct{}

func (g *GitUtil) GetGitRepo(folder string) *git.Repository {
	for true {
		repo, err := git.PlainOpen(folder)
		if err != nil {
			if folder == "/" {
				return nil
			} else {
				folder = filepath.Dir(folder)
			}
		}
		if repo != nil {
			return repo
		}
	}
	return nil
}

func (g *GitUtil) GetGitRemoteUrl(repo *git.Repository) string {
	remotes, err := repo.Remotes()
	if err != nil {
		log.Fatal(err)
	}
	if len(remotes) == 0 {
		return ""
	}
	return remotes[0].Config().URLs[0]
}

func (g *GitUtil) GetGitBranch(repo *git.Repository) string {
	head, err := repo.Head()
	if err != nil {
		log.Fatal(err)
	}
	return strings.ReplaceAll(head.Name().String(), "refs/heads/", "")
}
