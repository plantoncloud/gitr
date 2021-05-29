package lib_test

import (
	gitr "github.com/swarupdonepudi/gitr/v2/pkg"
	"testing"
)

func TestRemUrls(t *testing.T) {
	var urlTests = []struct {
		provider    gitr.ScmProvider
		branch      string
		remote      string
		expectedUrl string
	}{
		{gitr.GitHub, "master", "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/tree/master"},
		{gitr.GitHub, "master", "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/tree/master"},
		{gitr.GitHub, "feat/custom-branch", "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/tree/feat/custom-branch"},
	}
	r := gitr.GitrWeb{
		Scheme: "https",
	}
	t.Run("validate remote urls", func(t *testing.T) {
		for _, u := range urlTests {
			r.Url = u.remote
			r.Provider = u.provider
			r.Branch = u.branch
			if r.GetRemUrl() != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetRemUrl())
			}
		}
	})
}

func TestPrsUrls(t *testing.T) {
	var urlTests = []struct {
		provider    gitr.ScmProvider
		remote      string
		expectedUrl string
	}{
		{gitr.GitHub, "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/pulls"},
		{gitr.GitHub, "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/pulls"},
		{gitr.GitLab, "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/merge_requests"},
		{gitr.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/merge_requests"},
	}
	r := gitr.GitrWeb{
		Scheme: "https",
	}
	t.Run("validate mr/pr urls", func(t *testing.T) {
		for _, u := range urlTests {
			r.Url = u.remote
			r.Provider = u.provider
			if r.GetPrsUrl() != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetPrsUrl())
			}
		}
	})
}

func TestIssuesUrls(t *testing.T) {
	var urlTests = []struct {
		provider    gitr.ScmProvider
		remote      string
		expectedUrl string
	}{
		{gitr.GitHub, "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/issues"},
		{gitr.GitHub, "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/issues"},
		{gitr.GitLab, "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/issues"},
		{gitr.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/issues"},
	}
	r := gitr.GitrWeb{
		Scheme: "https",
	}
	t.Run("validate issues urls", func(t *testing.T) {
		for _, u := range urlTests {
			r.Url = u.remote
			r.Provider = u.provider
			if r.GetIssuesUrl() != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetIssuesUrl())
			}
		}
	})
}

func TestTagsUrls(t *testing.T) {
	var urlTests = []struct {
		provider    gitr.ScmProvider
		remote      string
		expectedUrl string
	}{
		{gitr.GitHub, "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/tags"},
		{gitr.GitHub, "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/tags"},
		{gitr.GitLab, "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/tags"},
		{gitr.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/tags"},
	}
	r := gitr.GitrWeb{
		Scheme: "https",
	}
	t.Run("validate tags urls", func(t *testing.T) {
		for _, u := range urlTests {
			r.Url = u.remote
			r.Provider = u.provider
			if r.GetTagsUrl() != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetTagsUrl())
			}
		}
	})
}

func TestReleasesUrls(t *testing.T) {
	var urlTests = []struct {
		provider    gitr.ScmProvider
		remote      string
		expectedUrl string
	}{
		{gitr.GitHub, "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/releases"},
		{gitr.GitHub, "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/releases"},
		{gitr.GitLab, "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/releases"},
		{gitr.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/releases"},
	}
	r := gitr.GitrWeb{
		Scheme: "https",
	}
	t.Run("validate releases urls", func(t *testing.T) {
		for _, u := range urlTests {
			r.Url = u.remote
			r.Provider = u.provider
			if r.GetReleasesUrl() != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetReleasesUrl())
			}
		}
	})
}

func TestPipelinesUrls(t *testing.T) {
	var urlTests = []struct {
		provider    gitr.ScmProvider
		remote      string
		expectedUrl string
	}{
		{gitr.GitHub, "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/actions"},
		{gitr.GitHub, "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/actions"},
		{gitr.GitLab, "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/pipelines"},
		{gitr.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/pipelines"},
	}
	r := gitr.GitrWeb{
		Scheme: "https",
	}
	t.Run("validate pipelines urls", func(t *testing.T) {
		for _, u := range urlTests {
			r.Url = u.remote
			r.Provider = u.provider
			if r.GetPipelinesUrl() != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetPipelinesUrl())
			}
		}
	})
}

func TestBranchesUrls(t *testing.T) {
	var urlTests = []struct {
		provider    gitr.ScmProvider
		remote      string
		expectedUrl string
	}{
		{gitr.GitHub, "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/branches"},
		{gitr.GitHub, "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/branches"},
		{gitr.GitLab, "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/branches"},
		{gitr.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/branches"},
	}
	r := gitr.GitrWeb{
		Scheme: "https",
	}
	t.Run("validate branches urls", func(t *testing.T) {
		for _, u := range urlTests {
			r.Url = u.remote
			r.Provider = u.provider
			if r.GetBranchesUrl() != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetBranchesUrl())
			}
		}
	})
}

func TestCommitsUrls(t *testing.T) {
	var urlTests = []struct {
		provider    gitr.ScmProvider
		branch      string
		remote      string
		expectedUrl string
	}{
		{gitr.GitHub, "master", "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/commits/master"},
		{gitr.GitHub, "master", "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/commits/master"},
		{gitr.GitHub, "feat/custom-branch", "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr/commits/feat/custom-branch"},
		{gitr.GitLab, "master", "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/commits/master"},
		{gitr.GitLab, "master", "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/commits/master"},
		{gitr.GitLab, "feat/custom-branch", "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss/-/commits/feat/custom-branch"},
	}
	r := gitr.GitrWeb{
		Scheme: "https",
	}
	t.Run("validate commits urls", func(t *testing.T) {
		for _, u := range urlTests {
			r.Url = u.remote
			r.Branch = u.branch
			r.Provider = u.provider
			if r.GetCommitsUrl() != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetCommitsUrl())
			}
		}
	})
}

func TestWebUrls(t *testing.T) {
	var urlTests = []struct {
		provider    gitr.ScmProvider
		branch      string
		remote      string
		expectedUrl string
	}{
		{gitr.GitHub, "master", "git@github.com:swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr"},
		{gitr.GitHub, "master", "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr"},
		{gitr.GitHub, "feat/custom-branch", "https://github.com/swarupdonepudi/gitr.git", "https://github.com/swarupdonepudi/gitr"},
		{gitr.GitLab, "master", "git@gitlab.com:gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss"},
		{gitr.GitLab, "master", "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss"},
		{gitr.GitLab, "feat/custom-branch", "https://gitlab.com/gitlab-org/gitlab-foss.git", "https://gitlab.com/gitlab-org/gitlab-foss"},
	}
	r := gitr.GitrWeb{
		Scheme: "https",
	}
	t.Run("validate web urls", func(t *testing.T) {
		for _, u := range urlTests {
			r.Url = u.remote
			r.Branch = u.branch
			r.Provider = u.provider
			if r.GetWebUrl() != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, r.GetWebUrl())
			}
		}
	})
}
