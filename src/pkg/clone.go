package pkg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/kevinburke/ssh_config"
	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	ssh2 "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

type ScmProvider string

const (
	GitHub    ScmProvider = "github.com"
	GitLab    ScmProvider = "gitlab.com"
	BitBucket ScmProvider = "bitbucket.com"
)

type CloneUrl struct {
	protocol     string
	hostname     string
	url_path     string
	scm_provider ScmProvider
	repopath     string
	reponame     string
	sshCloneUrl  string
	httpCloneUrl string
}

func (c CloneUrl) get_ssh_clone_url() string {
	return fmt.Sprintf("git@%s:%s.git", c.hostname, c.repopath)
}

func (c CloneUrl) get_http_clone_url() string {
	return fmt.Sprintf("%s://%s/%s.git", c.protocol, c.hostname, c.repopath)
}

func get_absolute_path(pemFilePath string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if strings.HasPrefix(pemFilePath, "~/") {
		pemFilePath = filepath.Join(dir, pemFilePath[2:])
	}
	return pemFilePath
}

func is_browser_url(clone_url string) bool {
	return strings.HasSuffix(clone_url, ".git")
}

func get_scm_provider(hostname string) ScmProvider {
	switch hostname {
	case "github.com":
		return GitHub
	case "gitlab.com":
		return GitLab
	case "bitbucket.org":
		return BitBucket
	default:
		return GitHub
	}
}

func get_repo_path(clone_url string, url_path string, scm_provider ScmProvider) string {
	org_or_team := strings.Split(url_path, "/")[1]
	repo_name := strings.Split(url_path, "/")[2]
	return fmt.Sprintf("%s/%s", org_or_team, repo_name)
}

func get_repo_name(repopath string) string {
	return string(repopath[strings.LastIndex(repopath, "/")+1 : len(repopath)])
}

func parse_clone_url(clone_url string) CloneUrl {
	var clone_url_object = CloneUrl{}
	if !is_browser_url(clone_url) {
		clone_url_object.protocol = strings.Split(clone_url, "://")[0]
		clone_url_object.hostname = strings.Split(strings.Split(clone_url, "://")[1], "/")[0]
		clone_url_object.url_path = string(clone_url[strings.Index(clone_url, clone_url_object.hostname)+len(clone_url_object.hostname) : len(clone_url)])
		clone_url_object.scm_provider = get_scm_provider(clone_url_object.hostname)
		clone_url_object.repopath = get_repo_path(clone_url, clone_url_object.url_path, clone_url_object.scm_provider)
		clone_url_object.reponame = get_repo_name(clone_url_object.repopath)
	}
	return clone_url_object
}

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
		if fileExists(get_absolute_path(defaultSshKey)) {
			pkeyfilepath = ("~/.ssh/id_rsa")
		} else {
			return nil, errors.New("ssh auth not found")
		}
	}
	println(pkeyfilepath)
	pkeyfile := get_absolute_path(pkeyfilepath)
	println(pkeyfile)
	pem, _ := ioutil.ReadFile(pkeyfile)
	signer, _ := ssh.ParsePrivateKey(pem)
	return &ssh2.PublicKeys{User: "git", Signer: signer}, nil
}

func ssh_clone(clone_url_object CloneUrl) error {
	auth, ssh_err := set_up_ssh_auth(clone_url_object.hostname)

	if ssh_err != nil {
		return ssh_err
	}

	os.Mkdir(clone_url_object.reponame, os.ModePerm)
	_, err := git.PlainClone(clone_url_object.reponame, false, &git.CloneOptions{
		URL:      clone_url_object.get_ssh_clone_url(),
		Progress: os.Stdout,
		Auth:     auth,
	})
	return err
}

func http_clone(clone_url_object CloneUrl) error {
	os.Mkdir(clone_url_object.reponame, os.ModePerm)
	_, err := git.PlainClone(clone_url_object.reponame, false, &git.CloneOptions{
		URL:      clone_url_object.get_http_clone_url(),
		Progress: os.Stdout,
	})
	return err
}

func CloneRepo(clone_url string) {
	clone_url_object := parse_clone_url(clone_url)
	err_ssh := ssh_clone(clone_url_object)

	if err_ssh != nil {
		println("Failed SSH. Trying HTTP Clone")
		err_http := http_clone(clone_url_object)
		if err_http != nil {
			log.Fatal(err_http)
		}
	}
}
