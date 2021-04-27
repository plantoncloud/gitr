package lib_test

import (
	"fmt"
	gitr "github.com/swarupdonepudi/gitr/lib"
	"os"
	"testing"
)

func TestGitUtilGetRepo(t *testing.T) {
	gu := &gitr.GitUtil{}
	pwd, _ := os.Getwd()
	gr := gu.GetGitRepo(fmt.Sprintf("%s/test_data/r1-no-remote", pwd))
	if gr == nil {
		t.Errorf("expected git repo object but received nil")
	}
	u := gu.GetGitRemoteUrl(gr)
	if u != "" {
		t.Errorf("expected empty remote string but received %s", u)
	}
}

func TestGitUtilGetRepoFromSubDir(t *testing.T) {
	gu := &gitr.GitUtil{}
	pwd, _ := os.Getwd()
	gr := gu.GetGitRepo(fmt.Sprintf("%s/test_data/r1-no-remote/f1", pwd))
	if gr == nil {
		t.Errorf("expected git repo object but received nil")
	}
	if gu.GetGitRemoteUrl(gr) != "" {
		t.Errorf("expected empty remote string but received %s", gu.GetGitRemoteUrl(gr))
	}
	gr = gu.GetGitRepo(fmt.Sprintf("%s/test_data/r1-no-remote/f2", pwd))
	if gr == nil {
		t.Errorf("expected git repo object but received nil")
	}
	if gu.GetGitRemoteUrl(gr) != "" {
		t.Errorf("expected empty remote string but received %s", gu.GetGitRemoteUrl(gr))
	}
}

func TestGitUtilGetRepoWithRemote(t *testing.T) {
	gu := &gitr.GitUtil{}
	pwd, _ := os.Getwd()

	remoteUrl := "https://github.com/swarupdonepudi/non-existent-repo.git"

	gr := gu.GetGitRepo(fmt.Sprintf("%s/test_data/r2-with-remote/f1", pwd))
	if gr == nil {
		t.Errorf("expected git repo object but received nil")
	}
	if gu.GetGitRemoteUrl(gr) == "" {
		t.Errorf("expected %s remote string but received empty string", remoteUrl)
	}
	gr = gu.GetGitRepo(fmt.Sprintf("%s/test_data/r2-with-remote/f1/f1-1", pwd))
	if gu.GetGitRemoteUrl(gr) == "" {
		t.Errorf("expected %s remote string but received empty string", remoteUrl)
	}
	if gu.GetGitBranch(gr) != "master" {
		t.Errorf("expected master as the branch name but received %s", gu.GetGitBranch(gr))
	}
}

func TestGitUtilGetRepoWithRemoteWithCustomBranch(t *testing.T) {
	gu := &gitr.GitUtil{}
	pwd, _ := os.Getwd()
	customBranchName := "feat/test-branch"
	gr := gu.GetGitRepo(fmt.Sprintf("%s/test_data/r3-with-remote-custom-branch/f1", pwd))
	if gr == nil {
		t.Errorf("expected git repo object but received nil")
	}
	if gu.GetGitBranch(gr) != customBranchName {
		t.Errorf("expected %s as the branch name but received %s", customBranchName, gu.GetGitBranch(gr))
	}
}
