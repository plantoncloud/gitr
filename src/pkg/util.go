package pkg

import (
	"gopkg.in/src-d/go-git.v4"
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
