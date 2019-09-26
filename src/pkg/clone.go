package pkg

import (
	"errors"
	"io/ioutil"
	"log"
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

func set_up_ssh_auth(hostname string) (*ssh2.PublicKeys, error) {
	pkeyfilepath := ssh_config.Get(hostname, "IdentityFile")
	if strings.HasSuffix(pkeyfilepath, "identity") {
		var defaultSshKey = "~/.ssh/id_rsa"
		if fileExists(getAbsolutePath(defaultSshKey)) {
			pkeyfilepath = ("~/.ssh/id_rsa")
		} else {
			return nil, errors.New("ssh auth not found")
		}
	}
	pkeyfile := getAbsolutePath(pkeyfilepath)
	pem, _ := ioutil.ReadFile(pkeyfile)
	signer, _ := ssh.ParsePrivateKey(pem)
	return &ssh2.PublicKeys{User: "git", Signer: signer}, nil
}

func ssh_clone(clone_url_object RepoUrl) error {
	auth, ssh_err := set_up_ssh_auth(clone_url_object.HostName)

	if ssh_err != nil {
		return ssh_err
	}

	os.Mkdir(clone_url_object.RepoName, os.ModePerm)
	_, err := git.PlainClone(clone_url_object.RepoName, false, &git.CloneOptions{
		URL:      clone_url_object.GetSshCloneUrl(),
		Progress: os.Stdout,
		Auth:     auth,
	})
	return err
}

func http_clone(clone_url_object RepoUrl) error {
	os.Mkdir(clone_url_object.RepoName, os.ModePerm)
	_, err := git.PlainClone(clone_url_object.RepoName, false, &git.CloneOptions{
		URL:      clone_url_object.GetHttpCloneUrl(),
		Progress: os.Stdout,
	})
	return err
}

func CloneRepo(clone_url string) {
	gitrRepo := Parse(clone_url)
	print(gitrRepo.
	errSsh := ssh_clone(gitrRepo)

	if errSsh != nil {
		println("Failed SSH. Trying HTTP Clone")
		err_http := http_clone(gitrRepo)
		if err_http != nil {
			log.Fatal(err_http)
		}
	}
}
