package clone

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	ssh2 "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/kevinburke/ssh_config"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	"github.com/swarupdonepudi/gitr/v2/pkg/file"
	"github.com/swarupdonepudi/gitr/v2/pkg/url"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Clone(inputUrl string, creDir bool, cfg *config.GitrConfig) {
	s, err := config.GetScmSystem(cfg, url.GetHost(inputUrl))
	if err != nil {
		log.Fatal(err)
	}
	clonePath := GetClonePath(inputUrl, cfg.Clone.ScmHome, creDir || cfg.Clone.AlwaysCreDir, cfg.Clone.IncludeHostForCreDir)
	if url.IsGitUrl(inputUrl) {
		if url.IsGitSshUrl(inputUrl) {
			if err := sshClone(inputUrl, clonePath); err != nil {
				log.Fatalf("error cloning the repo. %v\n", err)
			}
		} else {
			if err := httpClone(inputUrl, clonePath); err != nil {
				log.Fatalf("error cloning the repo. %v\n", err)
			}
		}
	} else {
		if s.Provider == config.BitBucketDatacenter {
			println("clone using browser urls for bitbucket-datacenter is not supported")
			return
		}
		sshCloneUrl := fmt.Sprintf("git@%s:%s.git", s.Hostname, url.GetRepoPath(inputUrl))
		if err := sshClone(sshCloneUrl, clonePath); err != nil {
			fmt.Println("error cloning the repo using ssh. trying http clone...")
			httpCloneUrl := fmt.Sprintf("%s://%s/%s.git", s.Scheme, s.Hostname, url.GetRepoPath(inputUrl))
			if err := httpClone(httpCloneUrl, clonePath); err != nil {
				log.Fatalf("error cloning the repo using http. %v\n", err)
			}
		}
	}
}

func GetClonePath(repoUrl, scmHome string, creDir, includeHostForCreDir bool) string {
	clonePath := ""
	if creDir {
		if includeHostForCreDir {
			clonePath = fmt.Sprintf("%s/%s", url.GetHost(repoUrl), url.GetRepoPath(repoUrl))
		} else {
			clonePath = url.GetRepoPath(repoUrl)
		}
	} else {
		clonePath = url.GetRepoName(url.GetRepoPath(repoUrl))
	}
	if scmHome != "" {
		clonePath = fmt.Sprintf("%s/%s", scmHome, clonePath)
	}
	return clonePath
}

func setUpSshAuth(hostname string) (*ssh2.PublicKeys, error) {
	keyFilePath := ssh_config.Get(hostname, "IdentityFile")
	homeDir, _ := os.UserHomeDir()
	if strings.HasSuffix(keyFilePath, "identity") {
		var defaultSshKey = fmt.Sprintf("%s/.ssh/id_rsa", homeDir)
		if file.IsFileExists(file.GetAbsPath(defaultSshKey)) {
			keyFilePath = defaultSshKey
		} else {
			return nil, errors.New("ssh auth not found")
		}
	}
	pem, _ := ioutil.ReadFile(file.GetAbsPath(keyFilePath))
	signer, _ := ssh.ParsePrivateKey(pem)
	return &ssh2.PublicKeys{User: "git", Signer: signer}, nil
}

func httpClone(url, clonePath string) error {
	os.MkdirAll(clonePath, os.ModePerm)
	_, err := git.PlainClone(clonePath, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	return err
}

func sshClone(repoUrl, clonePath string) error {
	auth, sshErr := setUpSshAuth(url.GetHost(repoUrl))
	if sshErr != nil {
		return sshErr
	}
	os.MkdirAll(clonePath, os.ModePerm)
	_, err := git.PlainClone(clonePath, false, &git.CloneOptions{
		URL:      repoUrl,
		Progress: os.Stdout,
		Auth:     auth,
	})
	return err
}
