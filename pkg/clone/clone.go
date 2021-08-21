package clone

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/go-git/go-git/v5"
	ssh2 "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kevinburke/ssh_config"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	"github.com/swarupdonepudi/gitr/v2/pkg/file"
	"github.com/swarupdonepudi/gitr/v2/pkg/url"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"strings"
)

func Clone(cfg *config.GitrConfig, inputUrl string, creDir, dry bool) error {
	s, err := config.GetScmHost(cfg, url.GetHostname(inputUrl))
	if err != nil {
		return errors.Wrapf(err, "failed to clone git repo with %s url", inputUrl)
	}
	repoPath := url.GetRepoPath(inputUrl, s.Hostname, s.Provider)
	repoLocation := GetClonePath(inputUrl, creDir)
	if dry {
		printGitrCloneInfo(inputUrl, creDir || s.Clone.AlwaysCreDir, cfg)
		return nil
	}
	if file.IsDirExists(repoLocation) {
		println("\nrepo already exists. skipping cloning...")
	} else {
		if url.IsGitUrl(inputUrl) {
			if url.IsGitSshUrl(inputUrl) {
				if err := sshClone(inputUrl, repoLocation); err != nil {
					return errors.Wrap(err, "error cloning the repo")
				}
			} else {
				if err := httpClone(inputUrl, repoLocation); err != nil {
					return errors.Wrap(err, "error cloning the repo")
				}
			}
		} else {
			if s.Provider == config.BitBucketDatacenter || s.Provider == config.BitBucketCloud {
				println("gitr does not support clone using browser urls for bitbucket-datacenter & bitbucket.org")
				return nil
			}
			sshCloneUrl := GetSshCloneUrl(s.Hostname, repoPath)
			if err := sshClone(sshCloneUrl, repoLocation); err != nil {
				log.Warnf("failed to clone repo using ssh. trying http clone...")
				httpCloneUrl := GetHttpCloneUrl(s.Hostname, repoPath, s.Scheme)
				if err := httpClone(httpCloneUrl, repoLocation); err != nil {
					return errors.Wrap(err, "error cloning the repo using http")
				}
			}
		}
	}
	log.Infof("\nrepo path: %s\n", repoLocation)
	if cfg.CopyRepoPathCdCmdToClipboard {
		err := clipboard.WriteAll(fmt.Sprintf("cd %s", repoLocation))
		if err != nil {
			return errors.Wrap(err, "err copying repo path to clipboard")
		}
		log.Infof("\nnote: command to navigate to repo path has been added to clipboard. run cmd+v to paste the command\n\n")
	} else {
		fmt.Printf("\n*** run below command to navigate to repo path  ***\n\ncd %s\n\n", repoLocation)
	}
	return nil
}

func GetClonePath(inputUrl string, creDir bool) string {
	cfg, err := config.NewGitrConfig()
	if err != nil {
		log.Fatalf("failed to get gitr config. err: %v", err)
	}
	s, err := config.GetScmHost(cfg, url.GetHostname(inputUrl))
	if err != nil {
		log.Fatalf("failed to get scm host. err: %v", err)
	}
	repoPath := url.GetRepoPath(inputUrl, s.Hostname, s.Provider)
	repoName := url.GetRepoName(repoPath)
	scmHome := getScmHome(s.Clone.HomeDir, cfg.Scm.HomeDir)
	clonePath := ""
	if creDir {
		if s.Clone.IncludeHostForCreDir {
			clonePath = fmt.Sprintf("%s/%s", s.Hostname, repoPath)
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

func printGitrCloneInfo(inputUrl string, creDir bool, cfg *config.GitrConfig) {
	s, err := config.GetScmHost(cfg, url.GetHostname(inputUrl))
	repoPath := url.GetRepoPath(inputUrl, s.Hostname, s.Provider)
	repoName := url.GetRepoName(repoPath)
	scmHome := getScmHome(s.Clone.HomeDir, cfg.Scm.HomeDir)
	clonePath := GetClonePath(inputUrl, creDir)
	if err != nil {
		log.Fatal(err)
	}
	println("")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendRow(table.Row{"remote", inputUrl})
	t.AppendSeparator()
	t.AppendRow(table.Row{"provider", s.Provider})
	t.AppendSeparator()
	t.AppendRow(table.Row{"host", s.Hostname})
	t.AppendSeparator()
	t.AppendRow(table.Row{"repo-name", repoName})
	t.AppendSeparator()
	t.AppendRow(table.Row{"ssh-url", GetSshCloneUrl(s.Hostname, repoPath)})
	t.AppendSeparator()
	t.AppendRow(table.Row{"http-url", GetHttpCloneUrl(s.Hostname, repoPath, s.Scheme)})
	t.AppendSeparator()
	t.AppendRow(table.Row{"create-dir", s.Clone.AlwaysCreDir || creDir})
	t.AppendSeparator()
	t.AppendRow(table.Row{"scm-home", scmHome})
	t.AppendSeparator()
	t.AppendRow(table.Row{"clone-path", clonePath})
	t.AppendSeparator()
	t.Render()
	println("")
}

func getScmHome(scmHostHomeDir, scmHomeDir string) string {
	if scmHostHomeDir != "" {
		return scmHostHomeDir
	}
	if scmHomeDir != "" {
		return scmHomeDir
	}
	return file.GetPwd()
}
