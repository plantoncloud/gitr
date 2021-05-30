package url_test

import (
	"github.com/swarupdonepudi/gitr/v2/pkg/config"
	"github.com/swarupdonepudi/gitr/v2/pkg/url"
	"testing"
)

func TestIsGitUrl(t *testing.T) {
	var positiveUrlTests = []struct {
		url      string
		isGitUrl bool
	}{
		{"git@github.com:swarupdonepudi/gitr.git", true},
		{"https://github.com/swarupdonepudi/gitr.git", true},
		{"https://github.com/swarupdonepudi/gitr", false},
		{"git@github.com:swarupdonepudi/gitr", false},
	}
	var negativeUrlTests = []struct {
		url      string
		isGitUrl bool
	}{
		{"https://github.com/swarupdonepudi/gitr", false},
		{"git@github.com:swarupdonepudi/gitr", false},
	}
	t.Run("urls ending with .git should be git urls", func(t *testing.T) {
		for _, u := range positiveUrlTests {
			if url.IsGitUrl(u.url) != u.isGitUrl {
				t.Errorf("expected url %s as git url", u.url)
			}
		}
	})
	t.Run("urls not ending with .git should not be git urls", func(t *testing.T) {
		for _, u := range negativeUrlTests {
			if url.IsGitUrl(u.url) != u.isGitUrl {
				t.Errorf("expected url %s as not git url", u.url)
			}
		}
	})
}

func TestIsGitSshUrl(t *testing.T) {
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
			if url.IsGitSshUrl(u.url) != u.isGitSshUrl {
				t.Errorf("expected url %s as git url", u.url)
			}
		}
	})
	t.Run("urls not prefixed with ssh or git should not be git ssh urls", func(t *testing.T) {
		for _, u := range negativeUrlTests {
			if url.IsGitSshUrl(u.url) != u.isGitSshUrl {
				t.Errorf("expected url %s as not git url", u.url)
			}
		}
	})
}

func TestIsGitHttpUrlHasUsername(t *testing.T) {
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
			if url.IsGitHttpUrlHasUsername(u.url) != u.hasUsername {
				t.Errorf("expected %v but received %v for %s ", u.hasUsername, url.IsGitHttpUrlHasUsername(u.url), u.url)
			}
		}
	})
}

func TestIsGitRepoName(t *testing.T) {
	var repoNameTests = []struct {
		repoPath string
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
			if url.GetRepoName(u.repoPath) != u.repoName {
				t.Errorf("expected %s but got %s for %s path", u.repoName, url.GetRepoName(u.repoPath), u.repoPath)
			}
		}
	})
}

func TestGetRepoPath(t *testing.T) {
	var repoNameTests = []struct {
		url          string
		host         string
		provider     config.ScmProvider
		expectedPath string
	}{
		{"https://github.com/swarupdonepudi/gitr/blob/master/.gitattributes", "github.com", config.GitHub, "swarupdonepudi/gitr"},
		{"git@github.com/swarupdonepudi/gitr.git", "github.com", config.GitHub, "swarupdonepudi/gitr"},
		{"git@gitlab.com/swarupdonepudi/sample-repo.git", "gitlab.com", config.GitLab, "swarupdonepudi/sample-repo"},
		{"git@gitlab.com/swarupdonepudi/subgroup/sample-repo.git", "gitlab.com", config.GitLab, "swarupdonepudi/subgroup/sample-repo"},
		{"https://gitlab.com/swarupdonepudi/subgroup/subgroup2/repo-name/-/tree/master/.gitattributes", "gitlab.com", config.GitLab, "swarupdonepudi/subgroup/subgroup2/repo-name"},
		{"https://gitlab.com/swarupdonepudi/subgroup/repo-name/-/tree/master/.gitattributes", "gitlab.com", config.GitLab, "swarupdonepudi/subgroup/repo-name"},
	}

	t.Run("repo name from repo path", func(t *testing.T) {
		for _, u := range repoNameTests {
			returnedPath := url.GetRepoPath(u.url, u.host, u.provider)
			if returnedPath != u.expectedPath {
				t.Errorf("expected %s but got %s for %s url", u.expectedPath, returnedPath, u.url)
			}
		}
	})
}
