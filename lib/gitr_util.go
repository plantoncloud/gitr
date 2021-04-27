package lib

import (
	"log"
	"regexp"
	"strings"
)

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
	if repoPath != "" {
		levels := strings.Split(repoPath, "/")
		if len(levels) < 2 {
			log.Fatal("failed to parse repo name")
		}
		return levels[1]
	} else {
		return ""
	}
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
