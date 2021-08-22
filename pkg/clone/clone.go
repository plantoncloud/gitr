package clone

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/go-git/go-git/v5"
	ssh2 "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kevinburke/ssh_config"
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
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
	repoLocation, err := GetClonePath(cfg, inputUrl, creDir)
	if err != nil {

		return errors.Wrap(err, "failed to get clone path")
	}
	if dry {
		err := printGitrCloneInfo(cfg, inputUrl, creDir || s.Clone.AlwaysCreDir)
		if err != nil {
			return errors.Wrap(err, "failed to print gitr clone info")
		}
		return nil
	}
	if file.IsDirExists(repoLocation) {
		log.Info("repo already exists. skipping cloning...")
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
				log.Warn("gitr does not support clone using browser urls for bitbucket-datacenter & bitbucket.org")
				return nil
			}
			sshCloneUrl := GetSshCloneUrl(s.Hostname, repoPath)
			if err := sshClone(sshCloneUrl, repoLocation); err != nil {
				log.Warn("failed to clone repo using ssh. trying http clone...")
				httpCloneUrl := GetHttpCloneUrl(s.Hostname, repoPath, s.Scheme)
				if err := httpClone(httpCloneUrl, repoLocation); err != nil {
					return errors.Wrap(err, "error cloning the repo using http")
				}
			}
		}
	}
	log.Infof("repo path: %s", repoLocation)
	if cfg.CopyRepoPathCdCmdToClipboard {
		err := clipboard.WriteAll(fmt.Sprintf("cd %s", repoLocation))
		if err != nil {
			return errors.Wrap(err, "err copying repo path to clipboard")
		}
		log.Info("command to navigate to repo path has been added to clipboard. run cmd+v to paste the command")
	} else {
		log.Infof("run this command to navigate to repo path: cd %s", repoLocation)
	}
	return nil
}

func GetClonePath(cfg *config.GitrConfig, inputUrl string, creDir bool) (string, error) {
	s, err := config.GetScmHost(cfg, url.GetHostname(inputUrl))
	if err != nil {
		log.Fatalf("failed to get scm host. err: %v", err)
	}
	repoPath := url.GetRepoPath(inputUrl, s.Hostname, s.Provider)
	repoName := url.GetRepoName(repoPath)
	scmHome, err := getScmHome(s.Clone.HomeDir, cfg.Scm.HomeDir)
	if err != nil {
		return "", errors.Wrap(err, "failed to get scm home dir")
	}
	clonePath := ""
	if creDir || s.Clone.AlwaysCreDir {
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
	return clonePath, nil
}

func setUpSshAuth(hostname string) (*ssh2.PublicKeys, error) {
	keyFilePath := ssh_config.Get(hostname, "IdentityFile")
	homeDir, _ := os.UserHomeDir()
	if strings.HasSuffix(keyFilePath, "identity") {
		var defaultSshKey = fmt.Sprintf("%s/.ssh/id_rsa", homeDir)
		absPath, err := file.GetAbsPath(defaultSshKey)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get abs path of %s", defaultSshKey)
		}
		if file.IsFileExists(absPath) {
			keyFilePath = defaultSshKey
		} else {
			return nil, errors.New("ssh auth not found")
		}
	}
	keyFileAbsPath, err := file.GetAbsPath(keyFilePath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get abs path of %s", keyFilePath)
	}
	pem, _ := ioutil.ReadFile(keyFileAbsPath)
	signer, _ := ssh.ParsePrivateKey(pem)
	return &ssh2.PublicKeys{User: "git", Signer: signer}, nil
}

func httpClone(url, clonePath string) error {
	if err := os.MkdirAll(clonePath, os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to created dir %s", clonePath)
	}
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
	if err := os.MkdirAll(clonePath, os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to created dir %s", clonePath)
	}
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

func printGitrCloneInfo(cfg *config.GitrConfig, inputUrl string, creDir bool) error {
	s, err := config.GetScmHost(cfg, url.GetHostname(inputUrl))
	repoPath := url.GetRepoPath(inputUrl, s.Hostname, s.Provider)
	repoName := url.GetRepoName(repoPath)
	scmHome, err := getScmHome(s.Clone.HomeDir, cfg.Scm.HomeDir)
	if err != nil {
		return errors.Wrap(err, "failed to get scm home dir")
	}
	clonePath, err := GetClonePath(cfg, inputUrl, creDir)
	if err != nil {
		return errors.Wrap(err, "failed to get clone path")
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
	return nil
}

func getScmHome(scmHostHomeDir, scmHomeDir string) (string, error) {
	if scmHostHomeDir != "" {
		return scmHostHomeDir, nil
	}
	if scmHomeDir != "" {
		return scmHomeDir, nil
	}
	getwd, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "failed to get current dir")
	}
	return getwd, nil
}
