package lib_test

import (
	gitr "github.com/swarupdonepudi/gitr/lib"
	"testing"
)

func TestIsGitUrl(t *testing.T) {
	gru := &gitr.GitrUtil{}
	var positiveUrlTests = []struct {
		url      string
		isGitUrl      bool
	}{
		{"git@github.com:swarupdonepudi/gitr.git", true},
		{"https://github.com/swarupdonepudi/gitr.git", true},
		{"https://github.com/swarupdonepudi/gitr", false},
		{"git@github.com:swarupdonepudi/gitr", false},
	}
	var negativeUrlTests = []struct {
		url      string
		isGitUrl      bool
	}{
		{"https://github.com/swarupdonepudi/gitr", false},
		{"git@github.com:swarupdonepudi/gitr", false},
	}
	t.Run("urls ending with .git should be git urls", func(t *testing.T) {
		for _, u := range positiveUrlTests {
			if gru.IsGitUrl(u.url) != u.isGitUrl {
				t.Errorf("expected url %s as git url", u.url)
			}
		}
	})
	t.Run("urls not ending with .git should not be git urls", func(t *testing.T) {
		for _, u := range negativeUrlTests {
			if gru.IsGitUrl(u.url) != u.isGitUrl {
				t.Errorf("expected url %s as not git url", u.url)
			}
		}
	})
}

func TestIsGitSshUrl(t *testing.T) {
	gru := &gitr.GitrUtil{}
	var positiveUrlTests = []struct {
		url         string
		isGitSshUrl bool
	}{
		{"git@github.com:swarupdonepudi/gitr.git", true},
		{"ssh://github.com/swarupdonepudi/gitr.git", true},
	}
	var negativeUrlTests = []struct {
		url         string
		isGitSshUrl bool
	}{
		{"https://github.com/swarupdonepudi/gitr", false},
		{"github.com:swarupdonepudi/gitr.git", false},
	}
	t.Run("urls prefixed with ssh or git should be git ssh urls", func(t *testing.T) {
		for _, u := range positiveUrlTests {
			if gru.IsGitSshUrl(u.url) != u.isGitSshUrl {
				t.Errorf("expected url %s as git url", u.url)
			}
		}
	})
	t.Run("urls not prefixed with ssh or git should not be git ssh urls", func(t *testing.T) {
		for _, u := range negativeUrlTests {
			if gru.IsGitSshUrl(u.url) != u.isGitSshUrl {
				t.Errorf("expected url %s as not git url", u.url)
			}
		}
	})
}

func TestIsGitHttpUrlHasUsername(t *testing.T) {
	gru := &gitr.GitrUtil{}
	var usernameTests = []struct {
		url         string
		hasUsername bool
	}{
		{"https://swarup@github.com:swarupdonepudi/gitr.git", true},
		{"https://swarupd@github.com:swarupdonepudi/gitr", true},
		{"https://github.com/swarupdonepudi/gitr", false},
		{"github.com:swarupdonepudi/gitr.git", false},
	}

	t.Run("username in http url", func(t *testing.T) {
		for _, u := range usernameTests {
			if gru.IsGitHttpUrlHasUsername(u.url) != u.hasUsername {
				t.Errorf("expected %v but received %v for %s ", u.hasUsername, gru.IsGitHttpUrlHasUsername(u.url), u.url)
			}
		}
	})
}

func TestIsGitRepoName(t *testing.T) {
	gru := &gitr.GitrUtil{}
	var repoNameTests = []struct {
		repoPath      string
		repoName string
	}{
		{"swarupdonepudi/gitr.git", "gitr.git"},
		{"parent/sub/sub2/project-name.git", "project-name.git"},
		{"parent/sub/sub2/sub/project-name.git", "project-name.git"},
		{"no-path.git", "no-path.git"},
		{"parent/sub/git-repo", "git-repo"},
		{"parent/git-repo", "git-repo"},
		{"git-repo", "git-repo"},
	}

	t.Run("repo name from repo path", func(t *testing.T) {
		for _, u := range repoNameTests {
			if gru.GetRepoName(u.repoPath) != u.repoName {
				t.Errorf("expected %s but got %s for %s path", u.repoName, gru.GetRepoName(u.repoPath), u.repoPath)
			}
		}
	})
}

