package clone

import (
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
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

func Clone(inputUrl, scmHome string, creDir, copyCloneLocationCdCmdToClipboard bool, s *config.ScmHost) {
	repoPath := url.GetRepoPath(inputUrl, s.Hostname, s.Provider)
	repoName := url.GetRepoName(repoPath)
	clonePath := GetClonePath(s.Hostname, repoPath, repoName, scmHome, creDir || s.Clone.AlwaysCreDir, s.Clone.IncludeHostForCreDir)
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
		if s.Provider == config.BitBucketDatacenter || s.Provider == config.BitBucketCloud {
			println("gitr does not support clone using browser urls for bitbucket-datacenter & bitbucket.org")
			return
		}
		sshCloneUrl := GetSshCloneUrl(s.Hostname, repoPath)
		if err := sshClone(sshCloneUrl, clonePath); err != nil {
			fmt.Println("error cloning the repo using ssh. trying http clone...")
			httpCloneUrl := GetHttpCloneUrl(s.Hostname, repoPath, s.Scheme)
			if err := httpClone(httpCloneUrl, clonePath); err != nil {
				log.Fatalf("error cloning the repo using http. %v\n", err)
			}
		}
	}
	fmt.Printf("\ncloned path: %s\n", clonePath)
	if copyCloneLocationCdCmdToClipboard {
		err := clipboard.WriteAll(fmt.Sprintf("cd %s", clonePath))
		if err != nil {
			log.Fatalf("err copying cloned path to clipboard. %v\n", err)
		}
		fmt.Printf("\nnote: command to navigate to cloned location has been added to clipboard. run cmd+v to paste the command\n\n")
	} else {
		fmt.Printf("\n*** run below command to navigate to cloned location  ***\n\ncd %s\n\n", clonePath)
	}
}

func GetClonePath(scmHost, repoPath, repoName, scmHome string, creDir, includeHostForCreDir bool) string {
	clonePath := ""
	if creDir {
		if includeHostForCreDir {
			clonePath = fmt.Sprintf("%s/%s", scmHost, repoPath)
		} else {
			clonePath = repoPath
		}
	} else {
		clonePath = repoName
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
	auth, sshErr := setUpSshAuth(url.GetHostname(repoUrl))
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

func GetSshCloneUrl(hostname, repoPath string) string {
	return fmt.Sprintf("git@%s:%s.git", hostname, repoPath)
}

func GetHttpCloneUrl(hostname, repoPath string, scheme config.HttpScheme) string {
	return fmt.Sprintf("%s://%s/%s.git", scheme, hostname, repoPath)
}
