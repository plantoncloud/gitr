package pkg

import (
	"gopkg.in/src-d/go-git.v4"
	"log"
	"path/filepath"
)

func GetGitRepo(folder string) *git.Repository {
	var hasParent = true
	for hasParent {
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

func GetGitRemoteUrl(repo *git.Repository) string {
	remotes, err := repo.Remotes()
	if err != nil {
		log.Fatal(err)
	}
	remoteUrl := remotes[0].Config().URLs[0]
	return remoteUrl
}
