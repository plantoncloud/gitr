package git

import (
	"fmt"
	"os"
	"testing"
)

func TestGitUtilGetRepo(t *testing.T) {
	pwd, _ := os.Getwd()
	gr, err := GetGitRepo(fmt.Sprintf("%s/git_test_data/r1-no-remote", pwd))
	if err != nil {
		t.Errorf("not expecting error but got err: %s", err.Error())
	}
	t.Run("git sdk should load git repo object", func(t *testing.T) {
		if gr == nil {
			t.Errorf("expected git repo object but received nil")
		}
	})
	u, err := GetGitRemoteUrl(gr)
	if err != nil {
		t.Errorf("not expecting error but got err: %s", err.Error())
	}
	if u != "" {
		t.Errorf("expected empty remote string but received %s", u)
	}
}

func TestGitUtilGetRepoFromSubDir(t *testing.T) {
	pwd, _ := os.Getwd()
	gr, err := GetGitRepo(fmt.Sprintf("%s/git_test_data/r1-no-remote/f1", pwd))
	if err != nil {
		t.Errorf("not expecting error but got err: %s", err.Error())
	}
	t.Run("git sdk should load git repo object from sub folder", func(t *testing.T) {
		if gr == nil {
			t.Errorf("expected git repo object but received nil")
		}
	})
	t.Run("no-remote git repo object should return error", func(t *testing.T) {
		_, err := GetGitRemoteUrl(gr)
		if err == nil {
			t.Errorf("expecting error but did not get any")
		}
	})
}

func TestGitUtilGetRepoWithRemote(t *testing.T) {
	pwd, _ := os.Getwd()

	remoteUrl := "https://github.com/swarupdonepudi/non-existent-repo.git"

	gr, err := GetGitRepo(fmt.Sprintf("%s/git_test_data/r2-with-remote/f1/f1-1", pwd))
	if err != nil {
		t.Errorf("not expecting error but got err: %s", err.Error())
	}
	t.Run("with-remote git repo object should not have remote empty", func(t *testing.T) {
		u, err := GetGitRemoteUrl(gr)
		if err != nil {
			t.Errorf("not expecting error but got err: %s", err.Error())
		}
		if u == "" {
			t.Errorf("expected %s remote string but received empty string", remoteUrl)
		}
	})
	t.Run("loaded repo object should have correct branch", func(t *testing.T) {
		u, err := GetGitRemoteUrl(gr)
		if err != nil {
			t.Errorf("not expecting error but got err: %s", err.Error())
		}
		if u != "master" {
			t.Errorf("expected master as the branch name but received %s", u)
		}
	})
}

func TestGitUtilGetRepoWithRemoteWithCustomBranch(t *testing.T) {
	pwd, _ := os.Getwd()
	customBranchName := "feat/test-branch"
	gr, err := GetGitRepo(fmt.Sprintf("%s/git_test_data/r3-with-remote-custom-branch/f1", pwd))
	if err != nil {
		t.Errorf("not expecting error but got err: %s", err.Error())
	}
	t.Run("loaded repo object with non default branch checked out should have correct branch", func(t *testing.T) {
		u, err := GetGitBranch(gr)
		if err != nil {
			t.Errorf("not expecting error but got err: %s", err.Error())
		}
		if u != customBranchName {
			t.Errorf("expected %s as the branch name but received %s", customBranchName, u)
		}
	})
}
