package url

import (
	"github.com/skratchdot/open-golang/open"
	"regexp"
	"strings"
)

func IsGitUrl(repoUrl string) bool {
	return strings.HasSuffix(repoUrl, ".git")
}

func IsGitSshUrl(repoUrl string) bool {
	return strings.HasPrefix(repoUrl, "ssh://") || strings.HasPrefix(repoUrl, "git@")
}

func IsGitHttpUrlHasUsername(repoUrl string) bool {
	matched, err := regexp.MatchString("https*:\\/\\/.*@+.*", repoUrl)
	if err != nil {
		println(err.Error())
		return false
	} else {
		return matched
	}
}

func GetRepoName(repoPath string) string {
	return strings.Split(repoPath, "/")[strings.Count(repoPath, "/")]
}

func GetHost(url string) string {
	if url != "" {
		if IsGitSshUrl(url) {
			if strings.HasPrefix(url, "ssh://") {
				return strings.Split(strings.Split(url, "@")[1], "/")[0]
			} else {
				return strings.Split(strings.Split(url, "@")[1], ":")[0]
			}
		} else if IsGitHttpUrlHasUsername(url) {
			return strings.Split(strings.Split(url, "@")[1], "/")[0]
		} else {
			return strings.Split(strings.Split(url, "://")[1], "/")[0]
		}
	} else {
		return ""
	}
}

func GetRepoPath(url string) string {
	return url[strings.Index(url, GetHost(url))+1+len(GetHost(url)) : strings.Index(url, ".git")]
}

func OpenInBrowser(url string) {
	if url != "" {
		_ = open.Run(url)
	}
}
