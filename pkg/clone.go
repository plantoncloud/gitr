package pkg

import (
	"errors"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kevinburke/ssh_config"
	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	ssh2 "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func setUpSshAuth(hostname string) (*ssh2.PublicKeys, error) {
	pkeyfilepath := ssh_config.Get(hostname, "IdentityFile")
	if strings.HasSuffix(pkeyfilepath, "identity") {
		var defaultSshKey = "~/.ssh/id_rsa"
		if fileExists(getAbsolutePath(defaultSshKey)) {
			pkeyfilepath = "~/.ssh/id_rsa"
		} else {
			return nil, errors.New("ssh auth not found")
		}
	}
	pkeyfile := getAbsolutePath(pkeyfilepath)
	pem, _ := ioutil.ReadFile(pkeyfile)
	signer, _ := ssh.ParsePrivateKey(pem)
	return &ssh2.PublicKeys{User: "git", Signer: signer}, nil
}

func sshClone(gitrRepo GitrRepo) error {
	auth, sshErr := setUpSshAuth(gitrRepo.HostName)

	if sshErr != nil {
		return sshErr
	}

	os.Mkdir(gitrRepo.RepoName, os.ModePerm)
	_, err := git.PlainClone(gitrRepo.RepoName, false, &git.CloneOptions{
		URL:      gitrRepo.GitRemSshUrl,
		Progress: os.Stdout,
		Auth:     auth,
	})
	return err
}

func httpClone(gitrRepo GitrRepo) error {
	os.Mkdir(gitrRepo.RepoName, os.ModePerm)
	_, err := git.PlainClone(gitrRepo.RepoName, false, &git.CloneOptions{
		URL:      gitrRepo.GitRemHttpUrl,
		Progress: os.Stdout,
	})
	return err
}

func CloneRepo(cloneUrl string) {
	gitrRepo := ParseUrl(cloneUrl)
	if viper.GetBool("debug") {
		println(gitrRepo.ToString())
	}
	if gitrRepo.GitRemSshUrl == "" && gitrRepo.ScmProvider == GitLab {
		println("Clone operation using Browser URLs for Gitlab repos is currently not supported by gitr. Working on it")
		os.Exit(0)
	} else {
		errSsh := sshClone(gitrRepo)
		if errSsh != nil && viper.GetBool("debug") {
			println(errSsh.Error())
		}
		if errSsh != nil {
			if strings.Contains(errSsh.Error(), "handshake failed") {
				println("SSH Handshake Failed. Trying " + strings.ToUpper(gitrRepo.Protocol) + " Clone")
				errHttp := httpClone(gitrRepo)
				if errHttp != nil {
					println("Gitr failed to clone using " + gitrRepo.GitRemHttpUrl + " url. Error : " + errHttp.Error() + ". Use the below command to clone.\n git clone " + gitrRepo.GitRemHttpUrl)
				}
			} else if strings.Contains(errSsh.Error(), "remote repository is empty") {
				println("Gitr can not clone empty repositories. Run the below command to clone. \n git clone " + gitrRepo.GitRemSshUrl)
			} else {
				println(errSsh.Error())
			}
		}
	}
}
