package git

import (
	"github.com/go-git/go-git/v5"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
)

func GetGitRepo(folder string) *git.Repository {
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

func GetGitRemoteUrl(r *git.Repository) string {
	remotes, err := r.Remotes()
	if err != nil {
		log.Fatalf("failed to get remote url from git repo. err: %v", err)
	}
	if len(remotes) == 0 {
		return ""
	}
	return remotes[0].Config().URLs[0]
}

func GetGitBranch(r *git.Repository) string {
	head, err := r.Head()
	if err != nil {
		log.Fatalf("failed to get head from git repo. err: %v", err)
	}
	return strings.ReplaceAll(head.Name().String(), "refs/heads/", "")
}
