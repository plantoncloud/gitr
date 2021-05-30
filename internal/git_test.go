package internal

import (
	"fmt"
	"os"
	"testing"
)

func TestGitUtilGetRepo(t *testing.T) {
	pwd, _ := os.Getwd()
	gr := getGitRepo(fmt.Sprintf("%s/git_test_data/r1-no-remote", pwd))
	t.Run("git sdk should load git repo object", func(t *testing.T) {
		if gr == nil {
			t.Errorf("expected git repo object but received nil")
		}
	})
	u := getGitRemoteUrl(gr)
	if u != "" {
		t.Errorf("expected empty remote string but received %s", u)
	}
}

func TestGitUtilGetRepoFromSubDir(t *testing.T) {
	pwd, _ := os.Getwd()
	gr := getGitRepo(fmt.Sprintf("%s/git_test_data/r1-no-remote/f1", pwd))
	t.Run("git sdk should load git repo object from sub folder", func(t *testing.T) {
		if gr == nil {
			t.Errorf("expected git repo object but received nil")
		}
	})
	t.Run("no-remote git repo object should have remote empty", func(t *testing.T) {
		if getGitRemoteUrl(gr) != "" {
			t.Errorf("expected empty remote string but received %s", getGitRemoteUrl(gr))
		}
	})
}

func TestGitUtilGetRepoWithRemote(t *testing.T) {
	pwd, _ := os.Getwd()

	remoteUrl := "https://github.com/swarupdonepudi/non-existent-repo.git"

	gr := getGitRepo(fmt.Sprintf("%s/git_test_data/r2-with-remote/f1/f1-1", pwd))
	t.Run("with-remote git repo object should not have remote empty", func(t *testing.T) {
		if getGitRemoteUrl(gr) == "" {
			t.Errorf("expected %s remote string but received empty string", remoteUrl)
		}
	})
	t.Run("loaded repo object should have correct branch", func(t *testing.T) {
		if getGitBranch(gr) != "master" {
			t.Errorf("expected master as the branch name but received %s", getGitBranch(gr))
		}
	})
}

func TestGitUtilGetRepoWithRemoteWithCustomBranch(t *testing.T) {
	pwd, _ := os.Getwd()
	customBranchName := "feat/test-branch"
	gr := getGitRepo(fmt.Sprintf("%s/git_test_data/r3-with-remote-custom-branch/f1", pwd))
	t.Run("loaded repo object with non default branch checked out should have correct branch", func(t *testing.T) {
		if getGitBranch(gr) != customBranchName {
			t.Errorf("expected %s as the branch name but received %s", customBranchName, getGitBranch(gr))
		}
	})
}
