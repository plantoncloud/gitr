package lib

import (
	"regexp"
	"strings"
)

type GitrUtil struct{}

func (gru *GitrUtil) IsGitUrl(repoUrl string) bool {
	return strings.HasSuffix(repoUrl, ".git")
}

func (gru *GitrUtil) IsGitSshUrl(repoUrl string) bool {
	return strings.HasPrefix(repoUrl, "ssh://") || strings.HasPrefix(repoUrl, "git@")
}

func (gru *GitrUtil) IsGitHttpUrlHasUsername(repoUrl string) bool {
	matched, err := regexp.MatchString("https*:\\/\\/.*@+.*", repoUrl)
	if err != nil {
		println(err.Error())
		return false
	} else {
		return matched
	}
}

func (gru *GitrUtil) GetRepoName(repoPath string) string {
	return strings.Split(repoPath, "/")[strings.Count(repoPath, "/")]
}

func (gru *GitrUtil) GetHost(url string) string {
	if url != "" {
		if gru.IsGitSshUrl(url) {
			if strings.HasPrefix(url, "ssh://") {
				return strings.Split(strings.Split(url, "@")[1], "/")[0]
			} else {
				return strings.Split(strings.Split(url, "@")[1], ":")[0]
			}
		} else if gru.IsGitHttpUrlHasUsername(url) {
			return strings.Split(strings.Split(url, "@")[1], "/")[0]
		} else {
			return strings.Split(strings.Split(url, "://")[1], "/")[0]
		}
	} else {
		return ""
	}
}

func (gru *GitrUtil) GetRepoPath(url string) string {
	return url[strings.Index(url, gru.GetHost(url))+1+len(gru.GetHost(url)) : strings.Index(url, ".git")]
}
