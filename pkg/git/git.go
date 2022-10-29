package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
)

// GetGitRepo returns a git repository by walking the file system upwards from the provided directory
// and returns an error when no git repository is found before reaching the "/"(root) directory
func GetGitRepo(folder string) (*git.Repository, error) {
	for true {
		gitRepo, err := git.PlainOpen(folder)
		if err != nil {
			if folder == "/" {
				break
			}
			folder = filepath.Dir(folder)
			continue
		}
		return gitRepo, nil
	}
	return nil, errors.New("git repository not found in the folder tree")
}

// GetGitRemoteUrl returns the first url in the first remote found in the git repository object
// and returns an errors either if there is no remotes or if there is no urls for the first remote.
func GetGitRemoteUrl(r *git.Repository) (string, error) {
	remotes, err := r.Remotes()
	if err != nil {
		log.Fatalf("failed to get remote url from git repo. err: %v", err)
	}
	if len(remotes) == 0 {
		return "", errors.New("no remotes found")
	}
	if len(remotes[0].Config().URLs) == 0 {
		return "", errors.Errorf("urls not found for %s remote", remotes[0].Config().Name)
	}
	return remotes[0].Config().URLs[0], nil
}

// GetGitBranch returns the name of the current branch
func GetGitBranch(r *git.Repository) (string, error) {
	head, err := r.Head()
	if err != nil {
		return "", errors.Wrap(err, "failed to get head from git repo")
	}
	return strings.ReplaceAll(head.Name().String(), "refs/heads/", ""), nil
}
