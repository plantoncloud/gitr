package lib

import (
	"errors"
	"fmt"
	ssh2 "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/kevinburke/ssh_config"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

type GitrUtil struct{}

func (gru *GitrUtil) IsGitUrl(repoUrl string) bool {
	return strings.HasSuffix(repoUrl, ".git")
}

func (gru *GitrUtil) IsGitSshUrl(repoUrl string) bool {
	return strings.HasPrefix(repoUrl, "ssh://") || strings.HasPrefix(repoUrl, "git@")
}

func (gru *GitrUtil) IsGitHttpUrlHasUsername(repoUrl string) bool {
	matched, err := regexp.MatchString("https*:\\/\\/.*@+.*", repoUrl)
	if err != nil {
		println(err.Error())
		return false
	} else {
		return matched
	}
}

func (gru *GitrUtil) GetRepoName(repoPath string) string {
	return strings.Split(repoPath, "/")[strings.Count(repoPath, "/")]
}

func (gru *GitrUtil) GetHost(url string) string {
	if url != "" {
		if gru.IsGitSshUrl(url) {
			if strings.HasPrefix(url, "ssh://") {
				return strings.Split(strings.Split(url, "@")[1], "/")[0]
			} else {
				return strings.Split(strings.Split(url, "@")[1], ":")[0]
			}
		} else if gru.IsGitHttpUrlHasUsername(url) {
			return strings.Split(strings.Split(url, "@")[1], "/")[0]
		} else {
			return strings.Split(strings.Split(url, "://")[1], "/")[0]
		}
	} else {
		return ""
	}
}

func (gru *GitrUtil) GetRepoPath(url string) string {
	return url[strings.Index(url, gru.GetHost(url))+1+len(gru.GetHost(url)) : strings.Index(url, ".git")]
}

func (gru *GitrUtil) fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (gru *GitrUtil) SetUpSshAuth(hostname string) (*ssh2.PublicKeys, error) {
	keyFilePath := ssh_config.Get(hostname, "IdentityFile")
	homeDir, _ := os.UserHomeDir()
	if strings.HasSuffix(keyFilePath, "identity") {
		var defaultSshKey = fmt.Sprintf("%s/.ssh/id_rsa", homeDir)
		if gru.fileExists(gru.getAbsolutePath(defaultSshKey)) {
			keyFilePath = defaultSshKey
		} else {
			return nil, errors.New("ssh auth not found")
		}
	}
	pem, _ := ioutil.ReadFile(gru.getAbsolutePath(keyFilePath))
	signer, _ := ssh.ParsePrivateKey(pem)
	return &ssh2.PublicKeys{User: "git", Signer: signer}, nil
}

func (gru *GitrUtil) getAbsolutePath(pemFilePath string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if strings.HasPrefix(pemFilePath, "~/") {
		pemFilePath = filepath.Join(dir, pemFilePath[2:])
	}
	return pemFilePath
}
