package lib_test

import (
	"github.com/swarupdonepudi/gitr/lib"
	"testing"
)

func TestRemUrls(t *testing.T) {
	var urlTests = []struct {
		provider    lib.ScmProvider
		branch      string
		remote      string
		expectedUrl string
	}{
		{lib.GitHub, "master", "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/tree/master"},
		{lib.GitHub, "master", "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/tree/master"},
		{lib.GitHub, "feat/custom-branch", "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/tree/feat/custom-branch"},
	}
	r := lib.RemoteRepo{
		Scheme: "https",
	}
	for _, u := range urlTests {
		r.Url = u.remote
		r.Provider = u.provider
		r.Branch = u.branch
		if r.GetRemUrl() != u.expectedUrl {
			t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetRemUrl())
		}
	}
}

func TestPrsUrls(t *testing.T) {
	var urlTests = []struct {
		provider    lib.ScmProvider
		remote      string
		expectedUrl string
	}{
		{lib.GitHub, "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/pulls"},
		{lib.GitHub, "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/pulls"},
		{lib.GitLab, "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/merge_requests"},
		{lib.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/merge_requests"},
	}
	r := lib.RemoteRepo{
		Scheme: "https",
	}
	for _, u := range urlTests {
		r.Url = u.remote
		r.Provider = u.provider
		if r.GetPrsUrl() != u.expectedUrl {
			t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetPrsUrl())
		}
	}
}

func TestIssuesUrls(t *testing.T) {
	var urlTests = []struct {
		provider    lib.ScmProvider
		remote      string
		expectedUrl string
	}{
		{lib.GitHub, "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/issues"},
		{lib.GitHub, "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/issues"},
		{lib.GitLab, "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/issues"},
		{lib.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/issues"},
	}
	r := lib.RemoteRepo{
		Scheme: "https",
	}
	for _, u := range urlTests {
		r.Url = u.remote
		r.Provider = u.provider
		if r.GetIssuesUrl() != u.expectedUrl {
			t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetIssuesUrl())
		}
	}
}

func TestTagsUrls(t *testing.T) {
	var urlTests = []struct {
		provider    lib.ScmProvider
		remote      string
		expectedUrl string
	}{
		{lib.GitHub, "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/tags"},
		{lib.GitHub, "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/tags"},
		{lib.GitLab, "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/tags"},
		{lib.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/tags"},
	}
	r := lib.RemoteRepo{
		Scheme: "https",
	}
	for _, u := range urlTests {
		r.Url = u.remote
		r.Provider = u.provider
		if r.GetTagsUrl() != u.expectedUrl {
			t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetTagsUrl())
		}
	}
}

func TestReleasesUrls(t *testing.T) {
	var urlTests = []struct {
		provider    lib.ScmProvider
		remote      string
		expectedUrl string
	}{
		{lib.GitHub, "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/releases"},
		{lib.GitHub, "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/releases"},
		{lib.GitLab, "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/releases"},
		{lib.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/releases"},
	}
	r := lib.RemoteRepo{
		Scheme: "https",
	}
	for _, u := range urlTests {
		r.Url = u.remote
		r.Provider = u.provider
		if r.GetReleasesUrl() != u.expectedUrl {
			t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetReleasesUrl())
		}
	}
}

func TestPipelinesUrls(t *testing.T) {
	var urlTests = []struct {
		provider    lib.ScmProvider
		remote      string
		expectedUrl string
	}{
		{lib.GitHub, "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/actions"},
		{lib.GitHub, "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/actions"},
		{lib.GitLab, "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/pipelines"},
		{lib.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/pipelines"},
	}
	r := lib.RemoteRepo{
		Scheme: "https",
	}
	for _, u := range urlTests {
		r.Url = u.remote
		r.Provider = u.provider
		if r.GetPipelinesUrl() != u.expectedUrl {
			t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetPipelinesUrl())
		}
	}
}

func TestBranchesUrls(t *testing.T) {
	var urlTests = []struct {
		provider    lib.ScmProvider
		remote      string
		expectedUrl string
	}{
		{lib.GitHub, "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/branches"},
		{lib.GitHub, "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/branches"},
		{lib.GitLab, "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/branches"},
		{lib.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/branches"},
	}
	r := lib.RemoteRepo{
		Scheme: "https",
	}
	for _, u := range urlTests {
		r.Url = u.remote
		r.Provider = u.provider
		if r.GetBranchesUrl() != u.expectedUrl {
			t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetBranchesUrl())
		}
	}
}

func TestCommitsUrls(t *testing.T) {
	var urlTests = []struct {
		provider    lib.ScmProvider
		branch      string
		remote      string
		expectedUrl string
	}{
		{lib.GitHub, "master", "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/commits/master"},
		{lib.GitHub, "master", "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/commits/master"},
		{lib.GitHub, "feat/custom-branch", "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/commits/feat/custom-branch"},
		{lib.GitLab, "master", "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/commits/master"},
		{lib.GitLab, "master", "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/commits/master"},
		{lib.GitLab, "feat/custom-branch", "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/commits/feat/custom-branch"},
	}
	r := lib.RemoteRepo{
		Scheme: "https",
	}
	for _, u := range urlTests {
		r.Url = u.remote
		r.Branch = u.branch
		r.Provider = u.provider
		if r.GetCommitsUrl() != u.expectedUrl {
			t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetCommitsUrl())
		}
	}
}
