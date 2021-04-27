package lib

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var err error

type GitOriginType string

const (
	Browser GitOriginType = "browser"
	Ssh     GitOriginType = "ssh"
	Http    GitOriginType = "http"
)

type GitOrigin struct {
	urlType  GitOriginType
	url      string
	scheme   string // http, https or ssh
	host     string
	repoPath string
	provider ScmProvider
	levels   []string
	repoName string
}

func (o *GitOrigin) ToString() string {
	return fmt.Sprintf("originUrlType\t: \t%s"+
		"\noriginUrl\t: \t%s"+
		"\nprovider\t: \t%s"+
		"\nscheme\t: \t%s"+
		"\nhost\t: \t%s"+
		"\nrepoPath\t\t: \t%s"+
		"\nrepoName\t: \t%s"+
		"\nweb\t\t: \t%s"+
		"\nprs\t\t: \t%s"+
		"\nbranches\t: \t%s"+
		"\ncommits\t\t: \t%s"+
		"\nissues\t\t: \t%s"+
		"\nreleases\t: \t%s"+
		"\npipelines\t: \t%s"+
		"\ntags\t: \t%s",
		o.urlType, o.url, o.provider, o.scheme, o.host,
		o.repoName, o.repoName,
		o.GetWebUrl(), o.GetPrsUrl(), o.GetBranchesUrl(),
		o.GetCommitsUrl(), o.GetIssuesUrl(), o.GetReleasesUrl(), o.GetPipelinesUrl(), o.GetTagsUrl())
}

func (o *GitOrigin) ScanOrigin() {
	pwd, _ := os.Getwd()
	repo := GetGitRepo(pwd)
	if repo != nil {
		remoteUrl := GetGitRemoteUrl(repo)
		o.parseGitRemoteUrl(remoteUrl)
		print(o.ToString())
	}
}

func (o *GitOrigin) GetWebUrl() string {
	switch o.provider {
	case BitBucketDatacenter:
		var project string
		if o.urlType == Http {
			project = o.levels[1]
		} else {
			project = o.levels[0]
		}
		return fmt.Sprintf("%s://%s/projects/%s/repos/%s", o.scheme, o.host, project, o.repoName)
	default:
		return fmt.Sprintf("%s://%s/%s", o.scheme, o.host, o.repoPath)
	}
}

func (o *GitOrigin) GetPrsUrl() string {
	switch o.provider {
	case GitHub:
		return fmt.Sprintf("%s/pulls", o.GetWebUrl())
	case GitLab:
		return fmt.Sprintf("%s/merge_requests", o.GetWebUrl())
	case BitBucketDatacenter, BitBucketCloud:
		return fmt.Sprintf("%s/pull-requests", o.GetWebUrl())
	default:
		return ""
	}
}

func (o *GitOrigin) GetBranchesUrl() string {
	switch o.provider {
	case GitLab:
		return fmt.Sprintf("%s/-/branches", o.GetWebUrl())
	default:
		return fmt.Sprintf("%s/branches", o.GetWebUrl())
	}
}

func (o *GitOrigin) GetCommitsUrl() string {
	switch o.provider {
	case GitLab:
		return fmt.Sprintf("%s/-/commits", o.GetWebUrl())
	default:
		return fmt.Sprintf("%s/commits", o.GetWebUrl())
	}
}

func (o *GitOrigin) GetTagsUrl() string {
	switch o.provider {
	case GitLab:
		return fmt.Sprintf("%s/-/tags", o.GetWebUrl())
	default:
		return fmt.Sprintf("%s/tags", o.GetWebUrl())
	}
}

func (o *GitOrigin) GetIssuesUrl() string {
	switch o.provider {
	case BitBucketDatacenter, BitBucketCloud:
		return ""
	default:
		return fmt.Sprintf("%s/issues", o.GetWebUrl())
	}
}

func (o *GitOrigin) GetReleasesUrl() string {
	switch o.provider {
	case GitHub:
		return fmt.Sprintf("%s/releases", o.GetWebUrl())
	default:
		return ""
	}
}

func (o *GitOrigin) GetPipelinesUrl() string {
	switch o.provider {
	case GitLab:
		return fmt.Sprintf("%s/pipelines", o.GetWebUrl())
	case BitBucketCloud:
		return fmt.Sprintf("%s/addon/pipelines/home", o.GetWebUrl())
	default:
		return ""
	}
}

func isGitSshUrl(repoUrl string) bool {
	return strings.HasPrefix(repoUrl, "ssh://") || strings.HasPrefix(repoUrl, "git@")
}

func isGitHttpUrlHasUsername(repoUrl string) bool {
	matched, err := regexp.MatchString("https*:\\/\\/.*@+.*", repoUrl)
	if err != nil {
		println(err.Error())
		return false
	} else {
		return matched
	}
}

func getHostName(url string) string {
	if isGitSshUrl(url) {
		if strings.HasPrefix(url, "ssh://") {
			return strings.Split(strings.Split(url, "@")[1], "/")[0]
		} else {
			return strings.Split(strings.Split(url, "@")[1], ":")[0]
		}
	} else if isGitHttpUrlHasUsername(url) {
		return strings.Split(strings.Split(url, "@")[1], "/")[0]
	} else {
		return strings.Split(strings.Split(url, "://")[1], "/")[0]
	}
}

func getLevels(urlPath string) []string {
	levels := strings.Split(urlPath, "/")
	return levels
}

func getRepoName(inputUrlType GitOriginType, scmProvider ScmProvider, levels []string) string {
	if scmProvider == BitBucketDatacenter {
		if inputUrlType == Http {
			return levels[2]
		} else if inputUrlType == Browser {
			return levels[3]
		} else {
			return levels[1]
		}
	} else {
		return levels[1]
	}
}

func (o *GitOrigin) parseGitRemoteUrl(repoUrl string) {
	o = &GitOrigin{
		url: repoUrl,
	}
	if isGitSshUrl(repoUrl) {
		o.urlType = Ssh
		o.scheme = "https"
	} else {
		o.urlType = Http
		o.scheme = strings.Split(repoUrl, "://")[0]
	}
	o.host = getHostName(repoUrl)

	o.repoPath = repoUrl[strings.Index(repoUrl, o.host)+1+len(o.host) : strings.Index(repoUrl, ".git")]

	c := &GitrConfig{}
	o.provider, err = c.GetScmProvider(o.host)

	if err != nil {
		log.Fatal(err)
	}

	o.levels = getLevels(o.repoPath)
	o.repoName = getRepoName(o.urlType, o.provider, o.levels)
}
