package web

import (
	"github.com/plantoncloud/gitr/pkg/config"
	"testing"
)

func TestRemUrls(t *testing.T) {
	var urlTests = []struct {
		provider    config.ScmProvider
		webUrl      string
		branch      string
		expectedUrl string
	}{
		{config.GitHub, "https://github.com/plantoncloud/gitr", "master", "https://github.com/plantoncloud/gitr/tree/master"},
		{config.GitHub, "https://github.com/plantoncloud/gitr", "feat/custom-branch", "https://github.com/plantoncloud/gitr/tree/feat/custom-branch"},
		{config.BitBucketCloud, "https://bitbucket.org/ramamohanraju/demo-project", "master", "https://bitbucket.org/ramamohanraju/demo-project/branch/master"},
	}
	t.Run("validate remote urls", func(t *testing.T) {
		for _, u := range urlTests {
			returnedUrl := GetRemUrl(u.provider, u.webUrl, u.branch)
			if returnedUrl != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, returnedUrl)
			}
		}
	})
}

func TestPrsUrls(t *testing.T) {
	var urlTests = []struct {
		provider    config.ScmProvider
		webUrl      string
		expectedUrl string
	}{
		{config.GitHub, "https://github.com/plantoncloud/gitr", "https://github.com/plantoncloud/gitr/pulls"},
		{config.GitHub, "https://github.com/plantoncloud/gitr", "https://github.com/plantoncloud/gitr/pulls"},
		{config.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss", "https://gitlab.com/gitlab-org/gitlab-foss/-/merge_requests"},
		{config.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api", "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api/-/merge_requests"},
	}
	t.Run("validate mr/pr urls", func(t *testing.T) {
		for _, u := range urlTests {
			returnedUrl := GetPrsUrl(u.provider, u.webUrl)
			if returnedUrl != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, returnedUrl)
			}
		}
	})
}

func TestIssuesUrls(t *testing.T) {
	var urlTests = []struct {
		provider    config.ScmProvider
		webUrl      string
		expectedUrl string
	}{
		{config.GitHub, "https://github.com/plantoncloud/gitr", "https://github.com/plantoncloud/gitr/issues"},
		{config.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api", "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api/-/issues"},
	}
	t.Run("validate issues urls", func(t *testing.T) {
		for _, u := range urlTests {
			returnedUrl := GetIssuesUrl(u.provider, u.webUrl)
			if returnedUrl != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, returnedUrl)
			}
		}
	})
}

func TestTagsUrls(t *testing.T) {
	var urlTests = []struct {
		provider    config.ScmProvider
		webUrl      string
		expectedUrl string
	}{
		{config.GitHub, "https://github.com/plantoncloud/gitr", "https://github.com/plantoncloud/gitr/tags"},
		{config.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api", "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api/-/tags"},
	}
	t.Run("validate tags urls", func(t *testing.T) {
		for _, u := range urlTests {
			returnedUrl := GetTagsUrl(u.provider, u.webUrl)
			if returnedUrl != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, returnedUrl)
			}
		}
	})
}

func TestReleasesUrls(t *testing.T) {
	var urlTests = []struct {
		provider    config.ScmProvider
		webUrl      string
		expectedUrl string
	}{
		{config.GitHub, "https://github.com/plantoncloud/gitr", "https://github.com/plantoncloud/gitr/releases"},
		{config.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api", "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api/-/releases"},
	}
	t.Run("validate releases urls", func(t *testing.T) {
		for _, u := range urlTests {
			returnedUrl := GetReleasesUrl(u.provider, u.webUrl)
			if returnedUrl != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, returnedUrl)
			}
		}
	})
}

func TestPipelinesUrls(t *testing.T) {
	var urlTests = []struct {
		provider    config.ScmProvider
		webUrl      string
		expectedUrl string
	}{
		{config.GitHub, "https://github.com/plantoncloud/gitr", "https://github.com/plantoncloud/gitr/actions"},
		{config.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api", "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api/-/pipelines"},
	}
	t.Run("validate pipelines urls", func(t *testing.T) {
		for _, u := range urlTests {
			returnedUrl := GetPipelinesUrl(u.provider, u.webUrl)
			if returnedUrl != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, returnedUrl)
			}
		}
	})
}

func TestBranchesUrls(t *testing.T) {
	var urlTests = []struct {
		provider    config.ScmProvider
		webUrl      string
		expectedUrl string
	}{
		{config.GitHub, "https://github.com/plantoncloud/gitr", "https://github.com/plantoncloud/gitr/branches"},
		{config.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api", "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api/-/branches"},
	}
	t.Run("validate branches urls", func(t *testing.T) {
		for _, u := range urlTests {
			returnedUrl := GetBranchesUrl(u.provider, u.webUrl)
			if returnedUrl != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, returnedUrl)
			}
		}
	})
}

func TestCommitsUrls(t *testing.T) {
	var urlTests = []struct {
		provider    config.ScmProvider
		webUrl      string
		branch      string
		expectedUrl string
	}{
		{config.GitHub, "https://github.com/plantoncloud/gitr", "master", "https://github.com/plantoncloud/gitr/commits/master"},
		{config.GitHub, "https://github.com/plantoncloud/gitr", "feat/custom", "https://github.com/plantoncloud/gitr/commits/feat/custom"},
		{config.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api", "main", "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api/-/commits/main"},
		{config.GitLab, "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api", "feat/custom", "https://gitlab.com/gitlab-org/gitlab-foss/gitlab-foss-api/-/commits/feat/custom"},
	}
	t.Run("validate commits urls", func(t *testing.T) {
		for _, u := range urlTests {
			returnedUrl := GetCommitsUrl(u.provider, u.webUrl, u.branch)
			if returnedUrl != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, returnedUrl)
			}
		}
	})
}

func TestWebUrls(t *testing.T) {
	var urlTests = []struct {
		provider    config.ScmProvider
		scheme      config.HttpScheme
		host        string
		repoPath    string
		expectedUrl string
	}{
		{config.GitHub, config.Https, "github.com", "swarupdonepudi/gitr", "https://github.com/plantoncloud/gitr"},
		{config.GitHub, config.Http, "github.com", "swarupdonepudi/gitr", "http://github.com/plantoncloud/gitr"},
		{config.GitLab, config.Https, "gitlab.com", "gitlab-org/gitlab-foss", "https://gitlab.com/gitlab-org/gitlab-foss"},
		{config.GitLab, config.Http, "gitlab.com", "gitlab-org/gitlab-foss", "http://gitlab.com/gitlab-org/gitlab-foss"},
		{config.GitLab, config.Https, "gitlab.com", "gitlab-org/gitlab-foss", "https://gitlab.com/gitlab-org/gitlab-foss"},
	}
	t.Run("validate web urls", func(t *testing.T) {
		for _, u := range urlTests {
			returnedUrl := GetWebUrl(u.provider, u.scheme, u.host, u.repoPath)
			if returnedUrl != u.expectedUrl {
				t.Errorf("expecting %s but got %s", u.expectedUrl, returnedUrl)
			}
		}
	})
}
